package doguu

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestArithmeticMean(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "black",
			args: args{n: 3},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(tt.want, arithmeticMean(tt.args.n)); diff != "" {
				t.Error(diff)
			}
		})
	}
}
