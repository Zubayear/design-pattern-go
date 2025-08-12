package singletonpattern_test

import (
	"os"
	"sync"
	"testing"

	singletonpattern "github.com/Zubayear/design-pattern-golang/singleton_pattern"
)

func TestSingletonSameInstance(t *testing.T) {

	t.Cleanup(func() {
		os.Remove("app.log")
	})

	l1 := singletonpattern.GetInstance()
	l2 := singletonpattern.GetInstance()

	if l1 != l2 {
		t.Errorf("Expected both loggers to be the same instance")
	}

}

func TestSingletonConcurrentAccess(t *testing.T) {

	t.Cleanup(func() {
		os.Remove("app.log")
	})

	const routines = 100
	var wg sync.WaitGroup
	wg.Add(routines)

	instances := make(chan *singletonpattern.Logger, routines)

	for range routines {
		go func() {
			defer wg.Done()
			instances <- singletonpattern.GetInstance()
		}()
	}

	wg.Wait()
	close(instances)

	var first *singletonpattern.Logger
	for inst := range instances {
		if first == nil {
			first = inst
		} else if first != inst {
			t.Errorf("Expected all instances to be same in concurrent scenario")
			break
		}
	}
}
