package irotsukegakari

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDouble(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "red",
			args: args{s: "This is usagi-san."},
			want: "\\e[30This is usagi-san.\\e[0m",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(tt.want, irotsuke(tt.args.s)); diff != "" {
				t.Error(diff)
			}
		})
	}
}
