package main

import "testing"

func Test_primality(t *testing.T) {
	type args struct {
		n int32
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Case 1",
			args: args{
				n: 12,
			},
			want: "Not prime",
		},
		{
			name: "Case 2",
			args: args{
				n: 5,
			},
			want: "Prime",
		},
		{
			name: "Case 3",
			args: args{
				n: 7,
			},
			want: "Prime",
		},
		{
			name: "Case 4",
			args: args{
				n: 5,
			},
			want: "Prime",
		},
		{
			name: "Case 5",
			args: args{
				n: 1,
			},
			want: "Prime",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := primality(tt.args.n); got != tt.want {
				t.Errorf("primality() = %v, want %v", got, tt.want)
			}
		})
	}
}
