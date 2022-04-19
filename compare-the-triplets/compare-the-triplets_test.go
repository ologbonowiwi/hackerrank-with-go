package main

import (
	"reflect"
	"testing"
)

func Test_compareTriplets(t *testing.T) {
	type args struct {
		a []int32
		b []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "should return [1,1]",
			args: args{
				a: []int32{5, 6, 7},
				b: []int32{3, 6, 10},
			},
			want: []int32{1, 1},
		},
		{
			name: "should return [2, 1]",
			args: args{
				a: []int32{17, 28, 30},
				b: []int32{99, 16, 8},
			},
			want: []int32{2, 1},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareTriplets(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("compareTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
