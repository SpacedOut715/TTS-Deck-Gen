package imageprocessing

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Deck(t *testing.T) {
	rootPath := "..\\test-images"
	exportPath := "..\\tmp-test"

	t.Run("Test Load Stats", func(t *testing.T) {
		deckDirs, err := FindAllEndDirsectories(rootPath)
		require.NoError(t, err)
		decks, err := LoadAllDecksDir(deckDirs)
		require.NoError(t, err)
		deck := decks[0]

		require.Equal(t, deck.Stats.cardWidth, 768)
		require.Equal(t, deck.Stats.cardHeight, 1202)
		require.Equal(t, deck.Stats.cardsRowCount, 2)
		require.Equal(t, deck.Stats.cardsColCount, 2)
		require.Equal(t, deck.Stats.pagesCount, 1)
	})

	t.Run("Test CheckSizes", func(t *testing.T) {
		deckDirs, err := FindAllEndDirsectories(rootPath)
		require.NoError(t, err)
		decks, err := LoadAllDecksDir(deckDirs)
		require.NoError(t, err)
		deck := decks[0]

		err = deck.CheckCardSizes()
		require.NoError(t, err)
	})

	t.Run("Export Test", func(t *testing.T) {
		deckDirs, err := FindAllEndDirsectories(rootPath)
		require.NoError(t, err)
		decks, err := LoadAllDecksDir(deckDirs)
		require.NoError(t, err)
		deck := decks[0]

		deckName, err := deck.ExportDeck(exportPath)
		require.NoError(t, err)
		require.NotEmpty(t, deckName)
	})

	t.Run("Export Test Big", func(t *testing.T) {
		deckDirs, err := FindAllEndDirsectories(rootPath)
		require.NoError(t, err)
		decks, err := LoadAllDecksDir(deckDirs)
		require.NoError(t, err)
		deck := decks[1]

		deckName, err := deck.ExportDeck(exportPath)
		require.NoError(t, err)
		require.NotEmpty(t, deckName)
	})

	t.Run("Export Test Big+Big Cards (local)", func(t *testing.T) {
		Source := "..\\generated"
		ResultDir := "..\\generated"

		deckDirs, err := FindAllEndDirsectories(Source)
		require.NoError(t, err)
		decks, err := LoadAllDecksDir(deckDirs)
		require.NoError(t, err)
		deck := decks[0]

		deckName, err := deck.ExportDeck(ResultDir)
		require.NoError(t, err)
		require.NotEmpty(t, deckName)
	})

	t.Run("Test Load Config", func(t *testing.T) {
		t.Skip()

		ConfigPath := "C:\\Users\\Sasa\\Desktop\\config.json"

		config, err := ParseFromJson(ConfigPath)
		require.NoError(t, err)

		decks, err := LoadAllDecksConfig(config)
		require.NoError(t, err)

		ExportDecks(decks, config.ExportPath)

	})

	t.Run("Test Full Functionality", func(t *testing.T) {
		t.Skip()

		Source := "E:\\Project\\Y\\Final\\A-Decks"
		ResultDir := "..\\generated"

		deckDirs, err := FindAllEndDirsectories(Source)
		require.NoError(t, err)
		decks, err := LoadAllDecksDir(deckDirs)
		require.NoError(t, err)

		ExportDecks(decks, ResultDir)
	})

	t.Run("Test Bug Width*10 >>> MaxRowCount", func(t *testing.T) {
		t.Skip()

		Source := "E:\\Project\\Y\\Final\\A-Decks\\Passives"
		ResultDir := "..\\generated"

		decks, err := LoadAllDecksDir([]string{Source})
		require.NoError(t, err)

		ExportDecks(decks, ResultDir)
	})
}

func Test_Directories(t *testing.T) {
	rootPath := "..\\test-images"

	t.Run("Find End Directories", func(t *testing.T) {
		dirs, err := FindAllEndDirsectories(rootPath)
		require.NoError(t, err)
		require.Len(t, dirs, 3)
	})
}
