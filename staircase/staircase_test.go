package main

import (
	"reflect"
	"testing"
)

func Test_buildStarcaise(t *testing.T) {
	type args struct {
		n int32
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "example 1",
			args: args{
				n: 4,
			},
			want: []string{"   #", "  ##", " ###", "####"},
		},
		{
			name: "example 2",
			args: args{
				n: 10,
			},
			want: []string{"         #", "        ##", "       ###", "      ####", "     #####", "    ######", "   #######", "  ########", " #########", "##########"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildStarcaise(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildStarcaise() = %v, want %v", got, tt.want)
			}
		})
	}
}
