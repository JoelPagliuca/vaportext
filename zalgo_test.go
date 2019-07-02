package main

import "testing"

func Test_zalgo(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"single word", args{"Suffer"}, "S͑u̧̇͌ͮ̐͟ͅf͔̝ͧ͞f͋̾̔̚e̬̘͌̊͠ŗ̶̛͌̇̈́͜"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := zalgo(tt.args.s); got != tt.want {
				t.Errorf("zalgo() = %v, want %v", got, tt.want)
			}
		})
	}
}
