package main

import "testing"

func TestVapor(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"single word", args{"Suffer"}, "Ｓｕｆｆｅｒ"},
		{"lowercase chars", args{"abcdefghijklmnopqrstuvwxyz"}, "ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚ"},
		{"uppercase chars", args{"ABCDEFGHIJKLMNOPQRSTUVWXYZ"}, "ＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺ"},
		{"special chars", args{"!@#$%^&*()"}, "！＠＃＄％＾＆＊（）"},
		{"spaces", args{"My Dude"}, "Ｍｙ Ｄｕｄｅ"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := vapor(tt.args.s); got != tt.want {
				t.Errorf("vapor() = %v, want %v", got, tt.want)
			}
		})
	}
}
