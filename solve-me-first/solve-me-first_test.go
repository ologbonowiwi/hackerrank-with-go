package main

import "testing"

func Test_solveMeFirst(t *testing.T) {
	type args struct {
		a uint32
		b uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "Getting 5 from sum of 3 and 2",
			args: args{
				a: 3,
				b: 2,
			},
			want: 5,
		},
		{
			name: "Getting 10 from sum of 6 and 4",
			args: args{
				a: 6,
				b: 4,
			},
			want: 10,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveMeFirst(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("solveMeFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}
