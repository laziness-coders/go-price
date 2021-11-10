package goprice

import (
	"reflect"
	"testing"
)

func TestNewPrice(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want *Price
	}{
		{
			name: "test",
			args: args{
				value: 12.00,
			},
			want: NewPrice(12.00),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPrice(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}