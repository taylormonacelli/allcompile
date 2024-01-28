package sailboat

import (
	"testing"
)

func TestSailboatProperties(t *testing.T) {
	testCases := []struct {
		desc        string
		constructor func() *Sailboat
		expected    Sailboat
	}{
		{
			desc: "Sailboat 1: single property specified using builder pattern",
			constructor: func() *Sailboat {
				return NewSailboatBuilder().
					Brand("Hobie Cat").
					Build()
			},
			expected: Sailboat{
				Brand:  "Hobie Cat",
				Color:  "DefaultColor",
				Length: 999,
			},
		},
		{
			desc: "Sailboat 2: no properties specified using builder pattern",
			constructor: func() *Sailboat {
				return NewSailboatBuilder().
					Build()
			},
			expected: Sailboat{
				Brand:  "DefaultBrand",
				Color:  "DefaultColor",
				Length: 999,
			},
		},
		{
			desc: "Sailboat 3: all properties specified using builder pattern",
			constructor: func() *Sailboat {
				return NewSailboatBuilder().
					Color("Blue").
					Brand("Catalina").
					Length(16).
					Build()
			},
			expected: Sailboat{
				Brand:  "Catalina",
				Color:  "Blue",
				Length: 16,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			boat := tc.constructor()
			if *boat != tc.expected {
				t.Errorf("Expected %v, Got %v", tc.expected, *boat)
			}
		})
	}
}
