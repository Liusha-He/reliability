package basics

// import "fmt"

// OCP
// open for extension and closed for modification
// Specification

type Color int

const (
	Red Color = iota
	Green
	Blue
)

type Size int

const (
	Small Size = iota
	Medium
	Large
)

type Filter struct {
	//
}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	result := make([]*Product, 0)

	for i, p := range products {
		if p.Color == color {
			result = append(result, &products[i])
		}
	}

	return result
}

func (f *Filter) FilterBySize(products []Product, size Size) []*Product {
	result := make([]*Product, 0)

	for i, p := range products {
		if p.Size == size {
			result = append(result, &products[i])
		}
	}

	return result
}

func (f *Filter) FilterBySizeAndColor(
	products []Product, size Size, color Color) []*Product {
	result := make([]*Product, 0)

	for i, p := range products {
		if p.Size == size && p.Color == color {
			result = append(result, &products[i])
		}
	}

	return result
}

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSprcification struct {
	Color Color
}

func (c ColorSprcification) IsSatisfied(p *Product) bool {
	return p.Color == c.Color
}

type SizeSpecification struct {
	Size Size
}

func (s SizeSpecification) IsSatisfied(p *Product) bool {
	return p.Size == s.Size
}

type BetterFilter struct{}

func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)

	for i, p := range products {
		if spec.IsSatisfied(&p) {
			result = append(result, &products[i])
		}
	}

	return result
}

type Product struct {
	Name  string
	Color Color
	Size  Size
}

type AndSpecification struct {
	First, Second Specification
}

func (a AndSpecification) IsSatisfied(p *Product) bool {
	return a.First.IsSatisfied(p) && a.Second.IsSatisfied(p)
}
