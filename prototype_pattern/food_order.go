package prototypepattern

import (
	"bytes"
	"encoding/gob"
	"log"
)

type prototype interface {
	deepCopy() (prototype, error)
}

type order struct {
	Id string
}

type foodOrder struct {
	CustomerName  string
	IsDelivery    bool
	OrderContents []string
	OrderId       *order
}

func (fo *foodOrder) ChangeOrderId(oid string) {
	fo.OrderId.Id = oid
}

func (fo *foodOrder) deepCopy() (prototype, error) {
	b := &bytes.Buffer{}
	e := gob.NewEncoder(b)
	err := e.Encode(fo)
	if err != nil {
		log.Printf("error: %v\n", err)
		return nil, nil
	}

	d := gob.NewDecoder(b)
	result := &foodOrder{}
	err = d.Decode(result)
	if err != nil {
		log.Printf("error: %v\n", err)
		return nil, err
	}
	return result, nil
}

func (f *foodOrder) DeepCopy(oi string) (id1, id2 string) {
	f2, err := f.deepCopy()
	if err != nil {
		log.Printf("Error: %v", err)
	}
	f3 := f2.(*foodOrder)
	f3.ChangeOrderId(oi)
	return f.OrderId.Id, f3.OrderId.Id
}

func (f *foodOrder) ShallowCopy(oi string) (id1, id2 string) {
	f2 := f
	f2.ChangeOrderId(oi)
	return f.OrderId.Id, f2.OrderId.Id
}
