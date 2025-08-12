package adapterpattern_test

import (
	"strings"
	"testing"

	adapterpattern "github.com/Zubayear/design-pattern-golang/adapter_pattern"
)

func TestCheckout(t *testing.T) {
	tests := []struct {
		name       string
		processor  adapterpattern.PaymentProcessor
		amount     float32
		currency   string
		wantStatus string
	}{
		{
			name:       "custom processor",
			processor:  adapterpattern.NewCustomPaymentProcessor(),
			amount:     10.0,
			currency:   "USD",
			wantStatus: "SUCCESS",
		},
		{
			name:       "legacy processor adapter",
			processor:  &adapterpattern.LegacyAdapter{},
			amount:     10.0,
			currency:   "USD",
			wantStatus: "FAILED",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			adapterpattern.Checkout(tt.processor, tt.amount, tt.currency)
			got := adapterpattern.CheckStatus(tt.processor)
			if got != tt.wantStatus {
				t.Errorf("Expected %v, got %v", tt.wantStatus, got)
			}
		})
	}

	legacyAdapter := adapterpattern.LegacyAdapter{}
	adapterpattern.Checkout(&legacyAdapter, 78.99, "USD")
	txnID := adapterpattern.GetTransactionID(&legacyAdapter)
	if !strings.HasPrefix(txnID, "ltxn") {
		t.Errorf("Expected to start with ltxn, Got %v", txnID)
	}
}
