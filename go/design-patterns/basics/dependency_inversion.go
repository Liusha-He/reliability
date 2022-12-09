package basics

import (
	"fmt"
	"strings"
)

type Relationship int

const (
	Parent Relationship = iota
	Child
	Sibling
)

type Person struct {
	Name string
}

type Info struct {
	From     *Person
	Relation Relationship
	To       *Person
}

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

// low-level module
type Relationships struct {
	Relations []Info
}

func (r *Relationships) AddChildAndParent(parent, child *Person) {
	r.Relations = append(
		r.Relations,
		Info{
			From:     parent,
			Relation: Parent,
			To:       child,
		},
	)
}

func (r Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)

	for i, v := range r.Relations {
		if v.Relation == Parent && v.From.Name == name {
			result = append(result, r.Relations[i].To)
		}
	}

	return result
}

// high-level module
type Research struct {
	// Relations Relationships
	Browser RelationshipBrowser
}

func (r *Research) Investigate(name string) string {
	names := make([]string, 0)
	count := 0

	for _, p := range r.Browser.FindAllChildrenOf(name) {
		count++
		names = append(names, p.Name)

	}

	if count > 0 {
		fmt.Printf("%s has %d children\n", name, count)
		return strings.Join(names, ",")
	} else {
		fmt.Printf("%s has no child\n", name)
		return ""
	}
}
