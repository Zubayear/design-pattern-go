package facadepattern

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// Provides a simple, easy to understand/user interface
// over a large and sophisticated body of code

type ChargeRequest struct {
	Amount                     float64
	Currency                   string
	CardNumber                 string
	CardExpiryMM, CardExpiryYY int
	CVV                        string
	CustomerID                 string
	Metadata                   map[string]string
}

type ChargeResponse struct {
	TransactionID string
	Status        string
	ProviderRef   string
	Amount        float64
	Currency      string
	Timestamp     time.Time
}

// ----------- Interfaces (subsystems) ------------

type CardTokenizer interface {
	Tokenize(ctx context.Context, cardNumber string, expiryMM, expiryYY int, cvv string) (token string, err error)
}

type FraudChecker interface {
	Check(ctx context.Context, req ChargeRequest) (ok bool, reason string, err error)
}

type CurrencyConverter interface {
	Convert(ctx context.Context, amount float64, from, to string) (converted float64, err error)
}

type GatewayClient interface {
	Charge(ctx context.Context, token string, amount float64, currency string, metadata map[string]string) (providedRef string, err error)
}

type TransactionStore interface {
	Save(ctx context.Context, resp ChargeResponse) error
}

// ----------- Simple Implementation ---------------

type SimpleTokenizer struct {
}

func (st *SimpleTokenizer) Tokenize(ctx context.Context, cardNumber string, expiryMM, expiryYY int, cvv string) (string, error) {
	if len(cardNumber) < 12 {
		return "", errors.New("invalid card number")
	}
	return fmt.Sprintf("tok_%x", time.Now().UnixNano()), nil
}

type SimpleFraud struct {
}

func (st *SimpleFraud) Check(ctx context.Context, req ChargeRequest) (bool, string, error) {
	if req.Amount > 10000 {
		return false, "amount too large", nil
	}
	if req.CustomerID == "" {
		return false, "missing customer id", nil
	}
	return true, "", nil
}

type FixedConverter struct {
	TargetCurrency string
	Rate           float64
}

func (f *FixedConverter) Convert(ctx context.Context, amount float64, from, to string) (float64, error) {
	// For demo: if same currency, no conversion. otherwise use rate.
	if from == to {
		return amount, nil
	}
	if to != f.TargetCurrency {
		// Not a real converter — just simple error for unsupported
		return 0, fmt.Errorf("unsupported conversion to %s", to)
	}
	return amount * f.Rate, nil
}

type HttpGateway struct {
}

func (hg *HttpGateway) Charge(ctx context.Context, token string, amount float64, currency string, metadata map[string]string) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case <-time.After(50 * time.Millisecond):
	}
	return fmt.Sprintf("prov_%x", time.Now().UnixNano()), nil
}

type MemoryStore struct {
	records []ChargeResponse
}

func (m *MemoryStore) Save(ctx context.Context, resp ChargeResponse) error {
	m.records = append(m.records, resp)
	return nil
}

// ----------- Facade ------------

type PaymentFacade struct {
	tokenizer CardTokenizer
	fraud     FraudChecker
	conv      CurrencyConverter
	client    GatewayClient
	store     TransactionStore

	providerCurrency string
	maxRetries       int
	retryInterval    time.Duration
}

func NewPaymentFacade(tokenizer CardTokenizer, fraud FraudChecker, conv CurrencyConverter, client GatewayClient, store TransactionStore) *PaymentFacade {
	return &PaymentFacade{
		tokenizer:        tokenizer,
		fraud:            fraud,
		conv:             conv,
		client:           client,
		store:            store,
		providerCurrency: "USD",
		maxRetries:       3,
		retryInterval:    100 * time.Millisecond,
	}
}

// ChargeCard orchestrates tokenization, fraud check, conversion, gateway call, and persistence.
func (p *PaymentFacade) ChargeCard(ctx context.Context, req ChargeRequest) (ChargeResponse, error) {
	var res ChargeResponse
	// 1. basic validation
	if req.Amount <= 0 {
		return res, errors.New("invalid amount")
	}

	// 2. fraud check
	ok, reason, err := p.fraud.Check(ctx, req)
	if err != nil {
		return res, fmt.Errorf("fraud check error: %w", err)
	}
	if !ok {
		res.Status = "declined"
		res.Timestamp = time.Now().UTC()
		_ = p.store.Save(ctx, res) // attempt store, ignore error for now
		return res, fmt.Errorf("fraud declined: %s", reason)
	}

	// 3. tokenize card
	token, err := p.tokenizer.Tokenize(ctx, req.CardNumber, req.CardExpiryMM, req.CardExpiryYY, req.CVV)
	if err != nil {
		return res, fmt.Errorf("tokenization failed: %w", err)
	}

	// 4. convert currency
	amountToCharge, err := p.conv.Convert(ctx, req.Amount, req.Currency, p.providerCurrency)
	if err != nil {
		return res, fmt.Errorf("currency conversion failed: %w", err)
	}

	// 5. attempt gateway call with retries
	var providerRef string
	var lastErr error
	for attempt := 0; attempt <= p.maxRetries; attempt++ {
		providerRef, lastErr = p.client.Charge(ctx, token, amountToCharge, p.providerCurrency, req.Metadata)
		if lastErr == nil {
			break
		}
		// simple backoff
		select {
		case <-ctx.Done():
			return res, ctx.Err()
		case <-time.After(p.retryInterval):
		}
	}
	if lastErr != nil {
		res.Status = "failed"
		res.Timestamp = time.Now().UTC()
		_ = p.store.Save(ctx, res)
		return res, fmt.Errorf("payment provider error: %w", lastErr)
	}

	// 6. success -> persist and return
	res = ChargeResponse{
		TransactionID: fmt.Sprintf("tx_%x", time.Now().UnixNano()),
		Status:        "success",
		ProviderRef:   providerRef,
		Amount:        amountToCharge,
		Currency:      p.providerCurrency,
		Timestamp:     time.Now().UTC(),
	}
	if err := p.store.Save(ctx, res); err != nil {
		// persist failure should not hide success but should be logged (here return error)
		return res, fmt.Errorf("payment succeeded but store failed: %w", err)
	}
	return res, nil
}
