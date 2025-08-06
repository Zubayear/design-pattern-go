package factorypattern

import (
	"bytes"
	"testing"
)

func TestGetCryptographyAlgorithmShouldReturnSHA256Algorithm(t *testing.T) {
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

func TestGetCryptographyAlgorithmShouldReturnSHA512Algorithm(t *testing.T) {
	cryptographyAlgo := GetCryptographyAlgorithm(SHA512)
	var buf bytes.Buffer
	msg := []byte("SHA512")
	cryptographyAlgo.GenerateHash(&buf, msg)
	result := buf.String()
	expected := "08af9a09ea069576ff02d11efe6822f9b09c105bcda44566e5b6b5c7a703dd778cbba35160190e79f3837c3275ba202231d19e764236c2bead580388e2fd0e7b\n"
	if result != expected {
		t.Errorf("Expected %q, but got %q", expected, result)
	}
}

func TestShouldPrintErrorWhenGeneratingHashWithNilForSHA512(t *testing.T) {
	cryptographyAlgo := GetCryptographyAlgorithm(SHA512)
	var buf bytes.Buffer
	// msg := []byte()
	cryptographyAlgo.GenerateHash(&buf, nil)
	result := buf.String()
	expected := "error\n"
	if result != expected {
		t.Errorf("Expected %q, but got %q", expected, result)
	}
}

func TestShouldPrintErrorWhenGeneratingHashWithNilForSHA256(t *testing.T) {
	cryptographyAlgo := GetCryptographyAlgorithm(SHA256)
	var buf bytes.Buffer
	// msg := []byte()
	cryptographyAlgo.GenerateHash(&buf, nil)
	result := buf.String()
	expected := "error\n"
	if result != expected {
		t.Errorf("Expected %q, but got %q", expected, result)
	}
}

func TestShouldReturnNilForDefaultCase(t *testing.T) {
	result := GetCryptographyAlgorithm(4)
	if result != nil {
		t.Errorf("Expected %v, but got %q", nil, result)
	}
}
