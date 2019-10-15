package generators

import "strconv"

// Alphabetical

type AlphabeticalGenerator struct {
	BaseAlphaGenerator
}

func NewAlphabeticalGenerator() *AlphabeticalGenerator {
	g := new(AlphabeticalGenerator)
	g.alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	g.idx = 0
	return g
}

// Numerical

type NumericalGenerator struct {
	IDGenerator
	idx int
}

func NewNumericalGenerator(start int) *NumericalGenerator {
	ng := new(NumericalGenerator)
	ng.idx = start
	return ng
}

func (ng *NumericalGenerator) Next() string {
	var next = ng.idx
	ng.idx += 1
	return strconv.Itoa(next)
}

func (ng *NumericalGenerator) Reset() {
	ng.idx = 0
}
