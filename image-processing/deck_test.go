package imageprocessing

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Deck(t *testing.T) {
	rootPath := "..\\test-images"

	t.Run("Test CheckSizes", func(t *testing.T) {
		decks, err := LoadAllDecks(rootPath)
		require.NoError(t, err)

		err = decks.Decks[0].CheckCardSizes()
		require.NoError(t, err)
	})

	t.Run("Test Get Count", func(t *testing.T) {
		decks, err := LoadAllDecks(rootPath)
		require.NoError(t, err)

		err = decks.Decks[0].CheckCardSizes()
		require.NoError(t, err)

		rowC, colC, pageC := decks.Decks[0].GetCount()
		require.Equal(t, rowC, 3)
		require.Equal(t, colC, 1)
		require.Equal(t, pageC, 1)
	})

	t.Run("Test Get Count Big", func(t *testing.T) {
		decks, err := LoadAllDecks(rootPath)
		require.NoError(t, err)

		err = decks.Decks[1].CheckCardSizes()
		require.NoError(t, err)

		rowC, colC, pageC := decks.Decks[1].GetCount()
		require.Equal(t, rowC, 10)
		require.Equal(t, colC, 7)
		require.Equal(t, pageC, 2)
	})
}

func Test_Directories(t *testing.T) {
	rootPath := "..\\test-images"

	t.Run("Find End Directories", func(t *testing.T) {
		dirs, err := findEndDirectories(rootPath)
		require.NoError(t, err)
		require.Len(t, dirs, 2)
	})
}
