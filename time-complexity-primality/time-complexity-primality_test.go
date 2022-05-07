package main

import (
	"testing"
	"time"
)

type args struct {
	n int32
}
type TestCase struct {
	name string
	args args
	want string
}

func Test_primality(t *testing.T) {
	tests := []TestCase{
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
			want: "Not prime",
		},
		{
			name: "Case 6",
			args: args{
				n: 9891,
			},
			want: "Not prime",
		},
		{
			name: "Case 7",
			args: args{
				n: 2,
			},
			want: "Prime",
		},
		{
			name: "Case 8",
			args: args{
				n: 3,
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

func TestWithTimeOut(t *testing.T) {
	timeout := time.After(10 * time.Second)
	done := make(chan bool)
	go func() {
		tests := []TestCase{
			{
				name: "Big O Case 1 — should pass in less of 4 seconds",
				args: args{
					n: 1000000007,
				},
				want: "Prime",
			},
			{
				name: "Big O Case 2 — should pass in less of 4 seconds",
				args: args{
					n: 1000000009,
				},
				want: "Prime",
			},
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := primality(tt.args.n); got != tt.want {
					t.Errorf("primality() = %v, want %v", got, tt.want)
				}
			})
		}
		time.Sleep(5 * time.Second)
		done <- true
	}()

	select {
	case <-timeout:
		t.Fatal("Test didn't finish in time")
	case <-done:
	}
}
