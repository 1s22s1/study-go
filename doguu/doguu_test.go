package doguu

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func TestArithmeticMean(t *testing.T) {
	type args struct {
		x []float64
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "normal1",
			args: args{x: []float64{1, 2, 3}},
			want: 2,
		},
		{
			name: "normal2",
			args: args{x: []float64{1, 2, 2}},
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

func TestGeometricMean(t *testing.T) {
	type args struct {
		x []float64
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "normal",
			args: args{x: []float64{1.1, 1.05, 1.03}},
			want: 1.06,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := []cmp.Option{cmpopts.EquateApprox(10, 1)}

			if diff := cmp.Diff(tt.want, geometricMean(tt.args.x), opts...); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestHamonicMean(t *testing.T) {
	type args struct {
		x []float64
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "normal",
			args: args{x: []float64{20, 30}},
			want: 24,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := []cmp.Option{cmpopts.EquateApprox(10, 1)}

			if diff := cmp.Diff(tt.want, harmonicMean(tt.args.x), opts...); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestVariance(t *testing.T) {
	type args struct {
		x []float64
	}

	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "normal",
			args: args{x: []float64{5, 6, 7, 7, 10}},
			want: 2.8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := []cmp.Option{cmpopts.EquateApprox(10, 1)}

			if diff := cmp.Diff(tt.want, variance(tt.args.x), opts...); diff != "" {
				t.Error(diff)
			}
		})
	}
}
