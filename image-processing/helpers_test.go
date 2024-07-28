package imageprocessing

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Something(t *testing.T) {
	deckDirs, err := FindAllEndDirsectories("..\\test-images")
	require.NoError(t, err)
	decks, err := LoadAllDecksDir(deckDirs)
	require.NoError(t, err)
	require.NotNil(t, decks)
}

func Test_Helpers(t *testing.T) {
	imagesPath := "..\\test-images\\A-green"
	imagePath := "..\\test-images\\A-green\\Elixir of Cleansing.png"

	t.Run("Get Image Files", func(t *testing.T) {
		imageFiles, err := GetImageFiles(imagesPath)
		require.NoError(t, err)
		require.Len(t, imageFiles, 3)
	})

	t.Run("Load Images", func(t *testing.T) {
		imageFiles, err := GetImageFiles(imagesPath)
		require.NoError(t, err)

		images, err := LoadImages(imageFiles)
		require.NoError(t, err)
		require.Len(t, images, 3)
	})

	t.Run("Load Image", func(t *testing.T) {
		image, err := LoadImage(imagePath)
		require.NoError(t, err)
		require.NotNil(t, image)
		require.Equal(t, image.Bounds().Dx(), 768)
		require.Equal(t, image.Bounds().Dy(), 1202)
	})
}
