package utils

import (
	"fmt"
	"math"
)

/* *
 * Environments [practising on maps]
 * */
type EnvConfigVal struct {
	Value string
}

type EnvConfig map[string](map[string]EnvConfigVal)

var Env = EnvConfig{
	"prod": {
		"port": {Value: "9000"},
	},
	"dev": {
		"port": {Value: "8000"},
	},
	"test": {
		"port": {Value: "7000"},
	},
}

func RoundTo(f float64, p uint) float64 {
	Env2 := make(EnvConfig)

	Env2["UAT"] = map[string]EnvConfigVal{
		"port": {Value: "80"},
	}

	fmt.Printf("%+v", Env2)

	r := math.Pow(10, float64(p))
	return math.Round(f*r) / r
}

/* *
 * Implementation for a custom error
 * */
type CustomError struct {
	Msg  string
	Name string
	E    interface{}
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("[%s]: %s", e.Name, e.Msg)
}

func IsCustomErr(err error) bool {
	_, ok := err.(*CustomError)
	return ok
}

// Reduce boilerplate for throwing an error
func ThrowErr(name, message string, cause interface{}) *CustomError {
	return &CustomError{
		Msg:  message,
		Name: name,
		E:    cause,
	}
}
