package gotl_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/treaster/gotl"
)

func TestBimap(t *testing.T) {
	type item struct {
		I int
		S string
	}

	dataSet := []item{
		{1, "one"},
		{2, "two"},
		{3, "three"},
	}

	getByInt := func(key int) (string, error) {
		for _, item := range dataSet {
			if key == item.I {
				return item.S, nil
			}
		}

		return "", fmt.Errorf("int query key %d not found", key)
	}

	getByStr := func(key string) (int, error) {
		for _, item := range dataSet {
			if key == item.S {
				return item.I, nil
			}
		}

		return 0, fmt.Errorf("str query key %q not found", key)
	}

	bimap := gotl.NewBimap[int, string](
		getByInt,
		getByStr,
	)

	{
		result, err := bimap.GetByT1(1)
		require.NoError(t, err)
		require.Equal(t, "one", result)
	}

	{
		_, err := bimap.GetByT1(9)
		require.Error(t, err)
	}

	{
		result, err := bimap.GetByT2("one")
		require.NoError(t, err)
		require.Equal(t, 1, result)
	}

	{
		_, err := bimap.GetByT2("nine")
		require.Error(t, err)
	}
}
