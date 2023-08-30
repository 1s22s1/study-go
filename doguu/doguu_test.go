package doguu

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestArithmeticMean(t *testing.T) {
	type args struct {
		x []int
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "normal1",
			args: args{x: []int{1, 2, 3}},
			want: 2,
		},
		{
			name: "normal2",
			args: args{x: []int{1, 2, 2}},
			want: 1.66,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := []cmp.Option{cmpopts.EquateApprox(10, 1)}

			if diff := cmp.Diff(tt.want, arithmeticMean(tt.args.x), opts...); diff != "" {
				t.Error(diff)
			}
		})
	}
}
