package imageprocessing

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Deck(t *testing.T) {
	rootPath := "..\\test-images"
	exportPath := "..\\tmp-test"

	t.Run("Test Load Stats", func(t *testing.T) {
		decks, err := LoadAllDecks(rootPath)
		require.NoError(t, err)
		deck := decks.Decks[0]

		require.Equal(t, deck.Stats.cardWidth, 768)
		require.Equal(t, deck.Stats.cardHeight, 1202)
		require.Equal(t, deck.Stats.cardsRowCount, 3)
		require.Equal(t, deck.Stats.cardsColCount, 1)
		require.Equal(t, deck.Stats.pagesCount, 1)
	})

	t.Run("Test CheckSizes", func(t *testing.T) {
		decks, err := LoadAllDecks(rootPath)
		require.NoError(t, err)
		deck := decks.Decks[0]

		err = deck.CheckCardSizes()
		require.NoError(t, err)
	})

	t.Run("Export Test", func(t *testing.T) {
		decks, err := LoadAllDecks(rootPath)
		require.NoError(t, err)

		err = decks.Decks[0].ExportDeck(exportPath)
		require.NoError(t, err)
	})

	t.Run("Export Test Big", func(t *testing.T) {
		decks, err := LoadAllDecks(rootPath)
		require.NoError(t, err)

		err = decks.Decks[1].ExportDeck(exportPath)
		require.NoError(t, err)
	})

	t.Run("Test Functionality", func(t *testing.T) {
		Source := "E:\\Project\\Y\\Final\\A-Decks"
		ResultDir := "..\\Generated"

		decks, err := LoadAllDecks(Source)
		require.NoError(t, err)

		decks.ExportDecks(ResultDir)
	})
}

func Test_Directories(t *testing.T) {
	rootPath := "..\\test-images"

	t.Run("Find End Directories", func(t *testing.T) {
		dirs, err := findEndDirectories(rootPath)
		require.NoError(t, err)
		require.Len(t, dirs, 3)
	})
}
