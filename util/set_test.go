package util_test

import (
	"testing"

	"github.com/oasdiff/go-common/util"
	"github.com/stretchr/testify/require"
)

func TestStringSet_Add(t *testing.T) {

	const item = "hello world"
	set := util.StringSet{}
	require.True(t, set.Add(item).Has(item))
}

func TestStringSet_Size(t *testing.T) {

	set := util.StringSet{}
	require.Equal(t, set.Add("1").Add("2").Size(), 2)
}

func TestStringSet_Clear(t *testing.T) {

	set := util.StringSet{}
	require.Equal(t, 2, set.Add("aaa").Add("bbbb").Size())
	set.Clear()
	require.Equal(t, 0, set.Size())
}

func TestIntSet_Add(t *testing.T) {

	const item = 7
	set := util.IntSet{}
	require.True(t, set.Add(item).Has(item))
}

func TestIntSet_Size(t *testing.T) {

	set := util.IntSet{}
	require.Equal(t, 2, set.Add(7).Add(55).Size())
}

func TestIntSet_Clear(t *testing.T) {

	set := util.IntSet{}
	require.Equal(t, set.Add(47).Add(1778).Size(), 2)
	set.Clear()
	require.Equal(t, 0, set.Size())
}
