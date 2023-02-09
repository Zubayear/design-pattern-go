package factory_pattern

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"io"
)

type CryptographyAlgorithmType int

const (
	SHA256 CryptographyAlgorithmType = iota + 1
	SHA512
)

// The Factory Design Pattern is a creational design pattern that provides an interface for
// creating objects in a superclass, but allows subclasses to alter the type of objects that will be created.
// The Factory Design Pattern is used when a class cannot anticipate the type of objects it needs to create,
// or when a class wants its subclasses to be able to specify the objects it creates.
type CryptographyAlgorithm interface {
	GenerateHash(w io.Writer, msg []byte)
}

type SHA256Algorithm struct {
}

type SHA512Algorithm struct {
}

func (s *SHA256Algorithm) GenerateHash(w io.Writer, msg []byte) {
	h := sha256.New()
	_, err := h.Write(msg)
	if err != nil {
		fmt.Printf("error %q", err)
	}
	bs := h.Sum(nil)
	hashStr := fmt.Sprintf("%x", bs)
	fmt.Fprintln(w, hashStr)
}

func (s *SHA512Algorithm) GenerateHash(w io.Writer, msg []byte) {
	h := sha512.New()
	_, err := h.Write(msg)
	if err != nil {
		fmt.Printf("error %q", err)
	}
	bs := h.Sum(nil)
	hashStr := fmt.Sprintf("%x", bs)
	fmt.Fprintln(w, hashStr)
}

func GetCryptographyAlgorithm(c CryptographyAlgorithmType) CryptographyAlgorithm {
	switch c {
	case SHA256:
		return &SHA256Algorithm{}
	case SHA512:
		return &SHA512Algorithm{}
	default:
		return nil
	}
}
