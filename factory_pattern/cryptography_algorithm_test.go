package factory_pattern

import (
	"bytes"
	"testing"
)

func TestGetCryptographyAlgorithm(t *testing.T) {
	cryptographyAlgo := GetCryptographyAlgorithm(SHA256)
	var buf bytes.Buffer
	msg := []byte("SHA256")
	cryptographyAlgo.GenerateHash(&buf, msg)
	result := buf.String()
	expected := "b3abe5d8c69b38733ad57ea75e83bcae42bbbbac75e3a5445862ed2f8a2cd677\n"
	if result != expected {
		t.Errorf("Expected %q, but got %q", expected, result)
	}
}
