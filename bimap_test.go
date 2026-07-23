package gotl_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/treaster/gotl"
)

func TestBimap(t *testing.T) {
	bimap := gotl.NewBimap[int, string](false)

	// Test missing records
	{
		_, ok := bimap.GetByT1(1)
		require.False(t, ok)

		_, ok = bimap.GetByT2("one")
		require.False(t, ok)
	}

	// Add some records
	{
		err := bimap.Add(1, "one")
		require.NoError(t, err)

		err = bimap.Add(2, "two")
		require.NoError(t, err)

		err = bimap.Add(3, "three")
		require.NoError(t, err)
	}

	// Test existing records
	{
		{
			result, ok := bimap.GetByT1(1)
			require.True(t, ok)
			require.Equal(t, "one", result)
		}

		{
			result, ok := bimap.GetByT2("one")
			require.True(t, ok)
			require.Equal(t, 1, result)
		}

		{
			result, ok := bimap.GetByT1(3)
			require.True(t, ok)
			require.Equal(t, "three", result)
		}

		{
			result, ok := bimap.GetByT2("three")
			require.True(t, ok)
			require.Equal(t, 3, result)
		}

		{
			_, ok := bimap.GetByT2("four")
			require.False(t, ok)
		}
	}

	// Test double-adding a record
	{
		err := bimap.Add(3, "four")
		require.Error(t, err)

		// The record is unchanged
		{
			result, ok := bimap.GetByT2("three")
			require.True(t, ok)
			require.Equal(t, 3, result)
		}

		{
			result, ok := bimap.GetByT1(3)
			require.True(t, ok)
			require.Equal(t, "three", result)
		}
	}
}

func TestBimapWithOverwrite(t *testing.T) {
	{
		bimap := gotl.NewBimap[int, string](true)

		// Add a record
		{
			err := bimap.Add(1, "one")
			require.NoError(t, err)

			result, ok := bimap.GetByT2("one")
			require.True(t, ok)
			require.Equal(t, 1, result)
		}

		// Overwrite the record
		{
			err := bimap.Add(1, "five")
			require.NoError(t, err)

			// The record is unchanged
			{
				result, ok := bimap.GetByT1(1)
				require.True(t, ok)
				require.Equal(t, "five", result)
			}

			{
				result, ok := bimap.GetByT2("five")
				require.True(t, ok)
				require.Equal(t, 1, result)
			}

			// The old record no longer works
			{
				_, ok := bimap.GetByT2("one")
				require.False(t, ok)
			}
		}
	}

	// Try overwriting in the opposite direction.
	{
		bimap := gotl.NewBimap[int, string](true)

		// Add a record
		{
			err := bimap.Add(1, "one")
			require.NoError(t, err)

			result, ok := bimap.GetByT2("one")
			require.True(t, ok)
			require.Equal(t, 1, result)
		}

		// Overwrite the record
		{
			err := bimap.Add(5, "one")
			require.NoError(t, err)

			// The record is unchanged
			{
				result, ok := bimap.GetByT1(5)
				require.True(t, ok)
				require.Equal(t, "one", result)
			}

			{
				result, ok := bimap.GetByT2("one")
				require.True(t, ok)
				require.Equal(t, 5, result)
			}

			// The old record no longer works
			{
				_, ok := bimap.GetByT1(1)
				require.False(t, ok)
			}
		}
	}
}
