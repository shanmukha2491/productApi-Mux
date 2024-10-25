package data

import "testing"

func TestValidation(t *testing.T) {
	p := &ProductDetails{Name: "SHanks", Price: 34}
	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
