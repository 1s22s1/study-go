package doguu

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestArithmeticMean(t *testing.T) {
	type args struct {
		x []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "normal1",
			args: args{x: []int{1, 2, 3}},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(tt.want, arithmeticMean(tt.args.x)); diff != "" {
				t.Error(diff)
			}
		})
	}
}
