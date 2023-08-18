package packer

import (
	"reflect"
	"testing"
)

func TestCalculatePacksWithPackSizes(t *testing.T) {
	packageSizes := []uint{300, 350, 400, 900, 1100, 1375, 2200, 5000}
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
			args: args{itemQuantity: 14001},
			want: map[uint]uint{
				5000: 2,
				2000: 2,
				250:  1,
			},
		},
		{
			name: "test5",
			args: args{itemQuantity: 751},
			want: map[uint]uint{
				1000: 1,
			},
		},
		{
			name: "test6",
			args: args{itemQuantity: 3758},
			want: map[uint]uint{
				2000: 2,
			},
		},
		{
			name: "test6",
			args: args{itemQuantity: 475},
			want: map[uint]uint{
				200: 3,
			},
		},
		{
			name: "test7",
			args: args{itemQuantity: 765},
			want: map[uint]uint{
				800: 1,
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
