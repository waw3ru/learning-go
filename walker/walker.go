package walker

import "reflect"

func walk(x interface{}, fn func(string)) {
	val := reflect.ValueOf(x)

	for field := range val.NumField() {
		field := val.Field(field)
		if field.Kind() == reflect.Ptr {
			field = field.Elem()
		}

		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			walk(field.Interface(), fn)
		}
	}
}
