package model

import (
	"reflect"
	"testing"
)

func TestIotBeerCommandFromIot_UnmarshalAssociatedData(t *testing.T) {
	type fields struct {
		Format  string
		Payload string
		Qos     string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []BeerPrediction
		wantErr bool
	}{
		{"5 data",fields{Payload:"[(0.15291722, 'n04254120 soap dispenser'), (0.13250275, 'n02783161 ballpoint, ballpoint pen, ballpen, Biro'), (0.068085946, 'n02948072 candle, taper, wax light'), (0.06672461, 'n04131690 saltshaker, salt shaker'), (0.03989944, 'n04376876 syringe')]"},
			[]BeerPrediction{{0.15291722,"n04254120 soap dispenser"},{0.13250275,"n02783161 ballpoint, ballpoint pen, ballpen, Biro"},{0.068085946,"n02948072 candle, taper, wax light"},{0.06672461,"n04131690 saltshaker, salt shaker"},{0.03989944,"n04376876 syringe"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := IotBeerCommandFromIot{
				Format:  tt.fields.Format,
				Payload: tt.fields.Payload,
				Qos:     tt.fields.Qos,
			}
			got, err := cmd.UnmarshalAssociatedData()
			if (err != nil) != tt.wantErr {
				t.Errorf("IotBeerCommandFromIot.UnmarshalAssociatedData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IotBeerCommandFromIot.UnmarshalAssociatedData() = %v, want %v", got, tt.want)
			}
		})
	}
}
