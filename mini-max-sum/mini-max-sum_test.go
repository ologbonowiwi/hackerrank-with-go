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
		want []int64
	}{
		{
			name: "Example 1",
			args: args{
				arr: []int32{int32(1), int32(3), int32(5), int32(7), int32(9)},
			},
			want: []int64{16, 24},
		},
		{
			name: "Example 2",
			args: args{
				arr: []int32{int32(1), int32(2), int32(3), int32(4), int32(5)},
			},
			want: []int64{10, 14},
		},
		{
			name: "Example 3",
			args: args{
				arr: []int32{int32(7), int32(69), int32(2), int32(221), int32(8974)},
			},
			want: []int64{299, 9271},
		},
		{
			name: "Example 4",
			args: args{
				arr: []int32{int32(501893267), int32(649027153), int32(379408215), int32(452968170), int32(487530619)},
			},
			want: []int64{1821800271, 2091419209},
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
