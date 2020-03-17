package controller

import (
	"github.com/iamgabrielwu/clock/pkg/controller/clock"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, clock.Add)
}
