package hello_world

import "fmt"

func Hello(name string, lang ...string) string {
	if name == "" {
		name = "World"
	}

	prefix := "Hello"

	if len(lang) == 0 {
		return fmt.Sprintf("%v, %v", prefix, name)
	}

	switch lang[0] {
	case "Spanish":
		prefix = "Hola"
	case "French":
		prefix = "Bonjour"
	case "German":
		prefix = "Hallo"
	}

	return fmt.Sprintf("%v, %v", prefix, name)
}
