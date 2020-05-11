package main

import (
	"testing"
)

func Test_vapor(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty string", args{""}, ""},
		{"single word", args{"Suffer"}, "Ｓｕｆｆｅｒ"},
		{"lowercase chars", args{allLower}, "ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚ"},
		{"uppercase chars", args{allUpper}, "ＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺ"},
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

func Test_fraktur(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty string", args{""}, ""},
		{"single word", args{"Suffer"}, "𝕾𝔲𝔣𝔣𝔢𝔯"},
		{"lowercase chars", args{allLower}, "𝔞𝔟𝔠𝔡𝔢𝔣𝔤𝔥𝔦𝔧𝔨𝔩𝔪𝔫𝔬𝔭𝔮𝔯𝔰𝔱𝔲𝔳𝔴𝔵𝔶𝔷"},
		{"uppercase chars", args{allUpper}, "𝕬𝕭𝕮𝕯𝕰𝕱𝕲𝕳𝕴𝕵𝕶𝕷𝕸𝕹𝕺𝕻𝕼𝕽𝕾𝕿𝖀𝖁𝖂𝖃𝖄𝖅"},
		{"special chars", args{"!@#$%^&*()"}, "!@#$%^&*()"},
		{"spaces", args{"My Dude"}, "𝕸𝔶 𝕯𝔲𝔡𝔢"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fraktur(tt.args.s); got != tt.want {
				t.Errorf("fraktur() = %v, want %v", got, tt.want)
			}
		})
	}
}
