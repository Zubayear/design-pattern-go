// Package adapterpattern implements the Adapter Design Pattern that allows
// incompatible interfaces to work together by converting the interface of
// one type into another that client expects.
package adapterpattern

import (
	"fmt"
	"time"
)

type PaymentProcessor interface {
	ProcessPayment(amount float32, currency string) bool
	PaymentStatus() string
	TransactionID() string
}

type CustomPaymentProcessor struct {
	paymentStatus string
	transactionID string
}

func NewCustomPaymentProcessor() *CustomPaymentProcessor {
	return &CustomPaymentProcessor{}
}

func (p *CustomPaymentProcessor) ProcessPayment(amount float32, currency string) bool {
	p.transactionID = fmt.Sprintf("txn_%d", time.Now().UnixMilli())
	p.paymentStatus = "SUCCESS"
	return true
}

func (p *CustomPaymentProcessor) PaymentStatus() string {
	return p.paymentStatus
}

func (p *CustomPaymentProcessor) TransactionID() string {
	return p.transactionID
}

type LegacyProcessor struct {
	referenceNumber uint64
	status          int16
}

func (l *LegacyProcessor) ExecuteTransaction(amount float64, currency string) uint64 {
	l.referenceNumber = uint64(time.Now().UnixMilli())
	l.status = int16(amount) % 2
	return l.referenceNumber
}

func (l *LegacyProcessor) CheckStatus(refID uint64) int16 {
	return l.status
}

func (l *LegacyProcessor) GetReferenceNumber() uint64 {
	return l.referenceNumber
}

// Checkout tester
func Checkout(paymentProcessor PaymentProcessor, amount float32, currency string) {
	paymentProcessor.ProcessPayment(amount, currency)
}

func CheckStatus(paymentProcessor PaymentProcessor) string {
	return paymentProcessor.PaymentStatus()
}

func GetTransactionID(paymentProcessor PaymentProcessor) string {
	return paymentProcessor.TransactionID()
}

// Adaptee - LegacyProcessor
// Target type - PaymentProcessor
// We need some type(Adapter) to convert/translate between Adaptee & Target type
// LegacyAdapter compose LegacyProcessor and then translate between type
type LegacyAdapter struct {
	legacyProcessor LegacyProcessor
	currentRefNum   uint64
}

func NewLegacyAdapter() *LegacyAdapter {
	return &LegacyAdapter{}
}

func (l *LegacyAdapter) ProcessPayment(amount float32, currency string) bool {
	res := l.legacyProcessor.ExecuteTransaction(float64(amount), currency)
	return res > 1
}

func (l *LegacyAdapter) PaymentStatus() string {
	status := l.legacyProcessor.CheckStatus(l.currentRefNum)
	if status == 1 {
		return "SUCCESS"
	}
	return "FAILED"
}

func (l *LegacyAdapter) TransactionID() string {
	ref := l.legacyProcessor.GetReferenceNumber()
	return fmt.Sprintf("ltxn_%d", ref)
}
