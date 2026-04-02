package walker

import "testing"

type ProfilePhoto struct {
	URL string
}

func TestWalker(t *testing.T) {
	testTable := []struct {
		Name  string
		Input interface{}
		calls []string
	}{
		{
			Name: "struct with a string field",
			Input: struct {
				Name string
			}{"John"},
			calls: []string{"John"},
		},
		{
			Name: "struct with a user defined struct field",
			Input: struct {
				Name  string
				Age   int
				Photo ProfilePhoto
			}{
				Name:  "John",
				Age:   10,
				Photo: ProfilePhoto{URL: "https://example.com/photo.jpg"},
			},
			calls: []string{"John", "https://example.com/photo.jpg"},
		},
		{
			Name: "struct with a user defined struct field with a pointer",
			Input: struct {
				Name  string
				Age   int
				Photo *ProfilePhoto
			}{
				Name:  "John",
				Age:   10,
				Photo: &ProfilePhoto{URL: "https://example.com/photo.jpg"},
			},
			calls: []string{"John", "https://example.com/photo.jpg"},
		},
	}

	for _, tt := range testTable {
		t.Run(tt.Name, func(t *testing.T) {
			var got []string
			walk(tt.Input, func(in string) {
				got = append(got, in)
			})

			if len(got) != len(tt.calls) {
				t.Errorf("wrong number of function calls, got %d want %d", len(got), len(tt.calls))
			}
		})
	}
}
