package facadepattern_test

import (
	"context"
	"testing"

	facadepattern "github.com/Zubayear/design-pattern-golang/facade_pattern"
)

func TestCharge(t *testing.T) {
	ctx := context.Background()
	facade := facadepattern.NewPaymentFacade(
		&facadepattern.SimpleTokenizer{},
		&facadepattern.SimpleFraud{},
		&facadepattern.FixedConverter{TargetCurrency: "USD", Rate: 100.0},
		&facadepattern.HttpGateway{},
		&facadepattern.MemoryStore{},
	)
	req := facadepattern.ChargeRequest{
		Amount:       12.34,
		Currency:     "USD",
		CardNumber:   "42424242424242424242",
		CardExpiryMM: 12,
		CardExpiryYY: 29,
		CVV:          "123",
		CustomerID:   "cust_1",
	}

	_, err := facade.ChargeCard(ctx, req)
	if err != nil {
		t.Error("Error occurred")
	}
}
