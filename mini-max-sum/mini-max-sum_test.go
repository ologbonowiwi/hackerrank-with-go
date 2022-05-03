package main

import (
	"reflect"
	"testing"
)

func Test_calcSums(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "Example 1",
			args: args{
				arr: []int32{int32(1), int32(3), int32(5), int32(7), int32(9)},
			},
			want: []int32{16, 24},
		},
		{
			name: "Example 2",
			args: args{
				arr: []int32{int32(1), int32(2), int32(3), int32(4), int32(5)},
			},
			want: []int32{10, 14},
		},
		{
			name: "Example 3",
			args: args{
				arr: []int32{int32(7), int32(69), int32(2), int32(221), int32(8974)},
			},
			want: []int32{10, 14},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcSums(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcSums() = %v, want %v", got, tt.want)
			}
		})
	}
}
