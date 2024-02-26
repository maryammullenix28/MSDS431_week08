package main

import (
	"math"
	"testing"

	"gonum.org/v1/gonum/stat"
)

// TestBootstrapCorrelation verifies that the bootstrapCorrelation function
// computes the correlation coefficient of the input data correctly.
func TestBootstrapCorrelation(t *testing.T) {
	// Create test data
	XDrafted := []float64{10, 20, 30, 40, 50}
	FPTS := []float64{50, 40, 30, 20, 10}
	R := 100

	// Compute the correlation coefficient
	result := bootstrapCorrelation(XDrafted, FPTS, R)

	// Verify the result
	expected := stat.Correlation(XDrafted, FPTS, nil)
	for i, r := range result {
		if math.Abs(r-expected) > 1e-5 {
			t.Errorf("BootstrapCorrelation result[%d] = %v, expected %v", i, r, expected)
		}
	}
}
