package controller

import (
	"github.com/binju2197/demo-operator/pkg/controller/democluster"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, democluster.Add)
}
