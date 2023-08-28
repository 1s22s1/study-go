package sample

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDouble(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal1",
			args: args{n: 1},
			want: 2,
		},
		{
			name: "normal2",
			args: args{n: 2},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(tt.want, double(tt.args.n)); diff != "" {
				t.Error(diff)
			}
		})
	}
}
