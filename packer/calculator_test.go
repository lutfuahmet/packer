package packer

import (
	"reflect"
	"testing"
)

func TestCalculatePacksWithPackSizes(t *testing.T) {
	packageSizes := []uint{250, 500, 1000, 2000, 5000}
	type args struct {
		itemQuantity uint
	}
	tests := []struct {
		name string
		args args
		want map[uint]uint
	}{
		{
			name: "test1",
			args: args{itemQuantity: 1},
			want: map[uint]uint{
				250: 1,
			},
		},
		{
			name: "test2",
			args: args{itemQuantity: 250},
			want: map[uint]uint{
				250: 1,
			},
		},
		{
			name: "test3",
			args: args{itemQuantity: 251},
			want: map[uint]uint{
				500: 1,
			},
		},
		{
			name: "test4",
			args: args{itemQuantity: 12001},
			want: map[uint]uint{
				5000: 2,
				2000: 1,
				250:  1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalculatePacksWithPackSizes(packageSizes, tt.args.itemQuantity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CalculatePacksWithPackSizes() = %v, want %v", got, tt.want)
			}
		})
	}
}
