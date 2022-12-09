package tests

import (
	"fmt"
	"os"
	"testing"

	basics "design-patterns/basics"

	assert "github.com/stretchr/testify/assert"
)

func TestUskovSubstitution(t *testing.T) {
	rc := &basics.Rectangle{Width: 2, Height: 3}
	ex, ac := basics.UsedSized(rc)

	assert.Equal(t, ex, ac, fmt.Sprintf("The area should be %d.", ac))

	sq := basics.Square{Size: 5}.Rectangle()
	ex, ac = basics.UsedSized(&sq)

	assert.Equal(t, ex, ac, fmt.Sprintf("The area should be %d.", ac))
}

func TestDependencyInversion(t *testing.T) {
	parent := basics.Person{"John"}
	child1 := basics.Person{"Chris"}
	child2 := basics.Person{"Finlay"}

	relations := basics.Relationships{}

	relations.AddChildAndParent(&parent, &child1)
	relations.AddChildAndParent(&parent, &child2)

	r := basics.Research{&relations}

	foundNames := r.Investigate("John")
	assert.Equal(t, foundNames, "Chris,Finlay", "found names wrong.")
}

func TestOpenClosed(t *testing.T) {
	// old way
	apple := basics.Product{"apple", basics.Green, basics.Small}
	tree := basics.Product{"tree", basics.Green, basics.Large}
	house := basics.Product{"house", basics.Blue, basics.Large}

	products := []basics.Product{apple, tree, house}

	f := basics.Filter{}
	expectedNames := []string{"apple", "tree"}

	for i, p := range f.FilterByColor(products, basics.Green) {
		assert.Equal(t, expectedNames[i], p.Name, "result incorrect...")
	}

	// specification
	fmt.Println("use the new method")

	bf := basics.BetterFilter{}
	greenspec := basics.ColorSprcification{basics.Green}

	for i, p := range bf.Filter(products, greenspec) {
		assert.Equal(t, expectedNames[i], p.Name, "result incorrect...")
	}

	largespec := basics.SizeSpecification{basics.Large}

	expectedNames = []string{"tree", "house"}

	for i, p := range bf.Filter(products, largespec) {
		assert.Equal(t, expectedNames[i], p.Name, "result incorrect...")
	}

	largegreenspec := basics.AndSpecification{greenspec, largespec}
	expectedNames = []string{"tree"}

	for i, p := range bf.Filter(products, largegreenspec) {
		assert.Equal(t, expectedNames[i], p.Name, "result incorrect")
	}
}

func TestSimpleResponse(t *testing.T) {
	j := basics.Journal{}

	j.AddEntry("hello")
	j.AddEntry("world")
	j.AddEntry("how")
	j.AddEntry("are")
	j.AddEntry("you")

	j.Save("log.txt")

	j.Load("log.txt")

	assert.Equal(
		t,
		"1: hello\n2: world\n3: how\n4: are\n5: you\n1: hello\n2: world\n3: how\n4: are\n5: you",
		j.String(),
		"output incorrect",
	)

	j.RemoveEntry(3)

	assert.Equal(
		t,
		"1: hello\n2: world\n3: how\n5: you\n1: hello\n2: world\n3: how\n4: are\n5: you",
		j.String(),
		"output incorrect...",
	)

	os.Remove("log.txt")

}

func TestInterfaceSegregation(t *testing.T) {
	// no test here
}
