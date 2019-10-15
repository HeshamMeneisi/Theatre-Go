package tests

import (
	"gotest.tools/assert"
	"testing"
	gn "theatre-mgr/generators"
)

func Test_BaseGenerator_CustomAlpha_CheckOutput(t *testing.T) {
	// Arrange
	var g = gn.NewBaseAlphaGenerator("xyz3456789")

	// Act
	n := 1042
	for i := 0; i < n; i++ {
		g.Next()
	}

	// Assert
	assert.Equal(t, "yx4z", g.Next())
}

func Test_NumericalGenerator_CheckOutput(t *testing.T) {
	// Arrange
	var g = gn.NewNumericalGenerator(0)

	// Act
	n := 1042
	for i := 0; i < n; i++ {
		g.Next()
	}

	// Assert
	assert.Equal(t, "1042", g.Next())
}

func Test_AlphabeticalGenerator_CheckOutput(t *testing.T) {
	// Arrange
	var g = gn.NewAlphabeticalGenerator()

	// Act
	n := 1042
	for i := 0; i < n; i++ {
		g.Next()
	}

	// Assert
	assert.Equal(t, "BOC", g.Next())
}
