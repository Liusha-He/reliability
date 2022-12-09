package builders

type Person struct {
	// Address
	StreetName, Postcode, City string

	// Occupation
	CompanyName, JobTitle string
	AnnualIncome          float32
}

type PersonBuilder struct {
	Person *Person
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}
