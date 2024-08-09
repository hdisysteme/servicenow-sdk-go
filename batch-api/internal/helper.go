package internal

import (
	"github.com/hdisysteme/servicenow-sdk-go/internal"
	"github.com/stretchr/testify/mock"
)

func ResetCalls(calls ...*mock.Call) {
	for _, call := range calls {
		if !internal.IsNil(call) {
			call.Unset()
		}
	}
}
