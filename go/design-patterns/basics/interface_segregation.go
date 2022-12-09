package basics

// interface segregation principle
// - you should never put too much into a single interface

type Document struct{}

type Printer interface {
	Print(d Document)
}

type Scanner interface {
	Scan(d Document)
}

type Phone interface {
	Fax(d Document)
}

type MultiFunctionInterface struct{}

func (m MultiFunctionInterface) Print(d Document) {
	// ...
}

func (m MultiFunctionInterface) Fax(d Document) {
	// ...
}

func (m MultiFunctionInterface) Scan(d Document) {
	// ...
}
