package prototype_pattern

import (
	"testing"
)

func Test_foodOrder_DeepCopy(t *testing.T) {
	type fields struct {
		CustomerName  string
		IsDelivery    bool
		OrderContents []string
		OrderId       *order
	}
	type args struct {
		oi string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantId1 string
		wantId2 string
	}{
		// TODO: Add test cases.
		{
			name: "DeepCopyTest",
			fields: fields{
				CustomerName:  "Elliot",
				IsDelivery:    false,
				OrderContents: []string{"Chicken", "Meat", "Peas"},
				OrderId: &order{
					Id: "ORD-12345",
				},
			},
			args: args{
				oi: "ORD-09876",
			},
			wantId1: "ORD-12345",
			wantId2: "ORD-09876",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &foodOrder{
				CustomerName:  tt.fields.CustomerName,
				IsDelivery:    tt.fields.IsDelivery,
				OrderContents: tt.fields.OrderContents,
				OrderId:       tt.fields.OrderId,
			}
			gotId1, gotId2 := f.DeepCopy(tt.args.oi)
			if gotId1 != tt.wantId1 {
				t.Errorf("foodOrder.DeepCopy() gotId1 = %v, want %v", gotId1, tt.wantId1)
			}
			if gotId2 != tt.wantId2 {
				t.Errorf("foodOrder.DeepCopy() gotId2 = %v, want %v", gotId2, tt.wantId2)
			}
		})
	}
}

func Test_foodOrder_ShallowCopy(t *testing.T) {
	type fields struct {
		CustomerName  string
		IsDelivery    bool
		OrderContents []string
		OrderId       *order
	}
	type args struct {
		oi string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantId1 string
		wantId2 string
	}{
		// TODO: Add test cases.
		{
			name: "ShallowCopyTest",
			fields: fields{
				CustomerName:  "Elliot",
				IsDelivery:    false,
				OrderContents: []string{"Chicken", "Meat", "Peas"},
				OrderId: &order{
					Id: "ORD-12345",
				},
			},
			args: args{
				oi: "ORD-09876",
			},
			wantId1: "ORD-09876",
			wantId2: "ORD-09876",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &foodOrder{
				CustomerName:  tt.fields.CustomerName,
				IsDelivery:    tt.fields.IsDelivery,
				OrderContents: tt.fields.OrderContents,
				OrderId:       tt.fields.OrderId,
			}
			gotId1, gotId2 := f.ShallowCopy(tt.args.oi)
			if gotId1 != tt.wantId1 {
				t.Errorf("foodOrder.ShallowCopy() gotId1 = %v, want %v", gotId1, tt.wantId1)
			}
			if gotId2 != tt.wantId2 {
				t.Errorf("foodOrder.ShallowCopy() gotId2 = %v, want %v", gotId2, tt.wantId2)
			}
		})
	}
}
