package operations

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewCreateParkingLot(t *testing.T) {
	tests := []struct {
		name string
		want *CreateParkingLot
	}{
		{
			name: "new parking log",
			want: &CreateParkingLot{
				opName: "create_parking_lot",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCreateParkingLot(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCreateParkingLot() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateParkingLot_GetName(t *testing.T) {
	type fields struct {
		opName   string
		capacity int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get Name",
			fields: fields{
				opName: "create_parking_lot",
			},
			want: "create_parking_lot",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cpl := &CreateParkingLot{
				opName:   tt.fields.opName,
				capacity: tt.fields.capacity,
			}
			if got := cpl.GetName(); got != tt.want {
				t.Errorf("CreateParkingLot.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateParkingLot_Parse(t *testing.T) {
	type fields struct {
		opName   string
		capacity int
	}
	type args struct {
		argVal string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "parse[valid]",
			fields: fields{},
			args: args{
				argVal: "6",
			},
			wantErr: false,
		},
		{
			name:   "parse[invalid]",
			fields: fields{},
			args: args{
				argVal: "6a",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cpl := &CreateParkingLot{
				opName:   tt.fields.opName,
				capacity: tt.fields.capacity,
			}
			if err := cpl.Parse(tt.args.argVal); (err != nil) != tt.wantErr {
				t.Errorf("CreateParkingLot.Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCreateParkingLot_Execute(t *testing.T) {
	type fields struct {
		opName   string
		capacity int
	}
	type args struct {
		argVal string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Execute[valid]",
			fields: fields{
				capacity: 1,
				opName:   "create_parking_lot",
			},
			args: args{
				argVal: "1",
			},
			want: fmt.Sprintf("Created a parking lot with %v slots", 1),
		},
		{
			name: "Execute[invalid]",
			fields: fields{
				capacity: 1,
				opName:   "create_parking_lot",
			},
			args: args{
				argVal: "1a",
			},
			want: "Invalid arguments",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cpl := &CreateParkingLot{
				opName:   tt.fields.opName,
				capacity: tt.fields.capacity,
			}
			if got := cpl.Execute(tt.args.argVal); got != tt.want {
				t.Errorf("CreateParkingLot.Execute() = %v, want %v", got, tt.want)
			}
		})
	}
}
