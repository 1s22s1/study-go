package irotsukegakari

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDouble(t *testing.T) {
	type args struct {
		s string
		c Color
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "black",
			args: args{s: "This is usagi-san.", c: Black},
			want: "\\e[30This is usagi-san.\\e[0m",
		},
		{
			name: "red",
			args: args{s: "This is usagi-san.", c: Red},
			want: "\\e[31This is usagi-san.\\e[0m",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(tt.want, irotsuke(tt.args.s, tt.args.c)); diff != "" {
				t.Error(diff)
			}
		})
	}
}
