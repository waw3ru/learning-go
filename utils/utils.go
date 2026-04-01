package utils

import (
	"fmt"
	"math"
)

/* *
 * Environments [practising on maps]
 * */
type EnvConfigVal map[string]interface{}
type EnvConfig map[string](EnvConfigVal)

var Env = EnvConfig{
	"prod": {
		"port": 2000,
	},
	"dev": {
		"port": 2000,
	},
	"test": {
		"port": 2000,
	},
}

func RoundTo(f float64, p uint) float64 {
	Env2 := make(EnvConfig)

	Env2["UAT"] = EnvConfigVal{
		"port": 8080,
	}

	fmt.Printf("%v", Env2["UAT"])

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
