package generators

type IDGenerator interface {
	Next() string
	Reset()
}

type BaseAlphaGenerator struct {
	idx int
	IDGenerator
	alphabet string
}

func NewBaseAlphaGenerator(alphabet string) *BaseAlphaGenerator {
	bg := new(BaseAlphaGenerator)
	bg.idx = 0
	bg.alphabet = alphabet
	return bg
}

func (bg *BaseAlphaGenerator) Next() string {
	id := ""
	var i = bg.idx
	var l = len(bg.alphabet)
	for i >= l {
		id = string(bg.alphabet[i%l]) + id
		i /= l
	}
	id = string(bg.alphabet[i]) + id
	bg.idx += 1
	return id
}

func (bg *BaseAlphaGenerator) Reset() {
	bg.idx = 0
}
