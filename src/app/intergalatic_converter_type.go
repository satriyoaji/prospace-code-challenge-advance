package app

type IntergalacticConverter struct {
	NumeralMappings map[string]string
	MetalValues     map[string]float64
}

type IIntergalacticConverter interface {
	ProcessInput(input string) error
}

func NewIntergalacticConverter() *IntergalacticConverter {
	return &IntergalacticConverter{
		NumeralMappings: make(map[string]string),
		MetalValues:     make(map[string]float64),
	}
}

// Todo: Define conversion methods and other logic here.
