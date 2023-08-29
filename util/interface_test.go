package util_test

import (
	"testing"

	"github.com/oasdiff/go-common/util"
	"github.com/stretchr/testify/require"
)

func TestToInterfaceSlice(t *testing.T) {

	type Employee struct {
		Name string
		Age  int
	}

	require.Len(t, util.ToInterfaceSlice([]*Employee{{
		Name: "Bill",
		Age:  72,
	}, {
		Name: "Tuna",
		Age:  67,
	}}), 2)
}
