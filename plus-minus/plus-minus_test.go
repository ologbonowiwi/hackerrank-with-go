package main

import (
	"reflect"
	"testing"
)

func Test_calculateProportions(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "example 1",
			args: args{
				arr: []int32{1, 1, 0, -1, -1},
			},
			want: []string{"0.400000", "0.400000", "0.200000"},
		},
		{
			name: "example 2",
			args: args{
				arr: []int32{-4, 3, -9, 0, 4, 1},
			},
			want: []string{"0.500000", "0.333333", "0.166667"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateProportions(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calculateProportions() = %v, want %v", got, tt.want)
			}
		})
	}
}
