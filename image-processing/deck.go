package imageprocessing

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

const (
	tts_maxWidth           = 10000
	tts_maxHeight          = 10000
	tts_minDeckHorizontalC = 2
	tts_maxDeckHorizontalC = 10
	tts_minDeckVerticalC   = 2
	tts_maxDeckVerticalC   = 7
	tts_maxDeckSize        = 70
)

type DeckStats struct {
	cardWidth  int
	cardHeight int

	pagesCount    int
	cardsRowCount int
	cardsColCount int
}

type Deck struct {
	Name  string
	Cards []image.Image
	Stats *DeckStats
}

func LoadAllDecksConfig(config *DecksConfig) ([]*Deck, error) {
	decks := make([]*Deck, len(config.Decks))

	for deckIdx, deckConfig := range config.Decks {

		imageFiles, err := GetImageFiles(deckConfig.DeckPath)
		if err != nil {
			return nil, err
		}

		images, err := LoadImages(imageFiles)
		if err != nil {
			return nil, err
		}

		deck, err := NewDeck(images, deckConfig.DeckFileName)
		if err != nil {
			return nil, err
		}

		decks[deckIdx] = deck

		fmt.Printf("Loaded deck %v with %v cards\n", deck.Name, len(deck.Cards))
	}

	fmt.Printf("Loaded %v decks\n", len(decks))

	return decks, nil
}

func LoadAllDecksDir(deckDirs []string) ([]*Deck, error) {
	decks := make([]*Deck, len(deckDirs))

	for deckIdx, deckDir := range deckDirs {

		imageFiles, err := GetImageFiles(deckDir)
		if err != nil {
			return nil, err
		}

		images, err := LoadImages(imageFiles)
		if err != nil {
			return nil, err
		}

		deckName := strings.Replace(deckDir[3:], "\\", "-", -1)
		deck, err := NewDeck(images, deckName)
		if err != nil {
			return nil, err
		}

		decks[deckIdx] = deck

		fmt.Printf("Loaded deck %v\n", deck.Name)
	}

	fmt.Printf("Created %v decks\n", len(decks))

	return decks, nil
}

func ExportDecks(decks []*Deck, resultDir string) error {
	if len(decks) == 0 {
		return errors.New("ExportDecks: empty decks")
	}

	for _, deck := range decks {
		deckName, err := deck.ExportDeck(resultDir)
		if err != nil {
			return err
		}

		fmt.Printf("Exported deck %v\n", deckName)
	}

	fmt.Printf("Exported %v decks to %v\n", len(decks), resultDir)

	return nil
}

func NewDeck(cards []image.Image, name string) (*Deck, error) {
	deck := &Deck{
		Name:  name,
		Cards: cards,
	}

	err := deck.CheckCardSizes()
	if err != nil {
		return nil, err
	}

	deck.Stats = deck.GetCount()

	return deck, nil
}

func (d *Deck) ExportDeck(resultDir string) (string, error) {
	var deckName string
	var err error

	for pageIdx := 0; pageIdx < d.Stats.pagesCount; pageIdx++ {
		// config for max cards per page
		rowCount := &d.Stats.cardsRowCount
		columnCount := &d.Stats.cardsColCount
		pageCardCount := *rowCount * *columnCount

		startIdx := pageIdx * pageCardCount
		endIdx := startIdx + min(pageCardCount, len(d.Cards)-startIdx)
		deckSlice := d.Cards[startIdx:endIdx]

		// If leftover slice is smaller than pageCardCount
		if len(deckSlice) < pageCardCount {
			*columnCount = len(deckSlice) / *rowCount + min(len(deckSlice)%*rowCount, 1)
			if len(deckSlice) < *rowCount {
				*rowCount = len(deckSlice)
			}
		}

		image := d.FillImage(deckSlice, *rowCount, *columnCount)

		deckName, err = d.SaveImage(image, resultDir, pageIdx, *rowCount, *columnCount, len(deckSlice))
		if err != nil {
			return deckName, err
		}
	}

	return deckName, nil
}

func (d *Deck) FillImage(deckSlice []image.Image, rowCount, columnCount int) image.Image {
	image := image.NewRGBA(image.Rect(0, 0, d.Stats.cardWidth*rowCount, d.Stats.cardHeight*columnCount))

	for columnIdx := 0; columnIdx < columnCount; columnIdx++ {
		for rowIdx := 0; rowIdx < rowCount; rowIdx++ {
			// If row is not full, break out
			if (columnIdx*rowCount + rowIdx) >= len(deckSlice) {
				break
			}

			d.FillCard(image, deckSlice[columnIdx*rowCount+rowIdx], rowIdx, columnIdx)
		}
	}

	return image
}

func (d *Deck) FillCard(image *image.RGBA, card image.Image, row, col int) {
	var color color.Color

	for cardHeight := 0; cardHeight < d.Stats.cardHeight; cardHeight++ {
		for cardWidth := 0; cardWidth < d.Stats.cardWidth; cardWidth++ {
			color = card.At(cardWidth, cardHeight)
			image.Set(row*d.Stats.cardWidth+cardWidth, col*d.Stats.cardHeight+cardHeight, color)
		}
	}
}

func (d *Deck) SaveImage(image image.Image, resultDir string, pageIdx, rowCount, columnCount, cardCount int) (string, error) {
	err := os.MkdirAll(resultDir, 0755)
	if err != nil {
		return "", fmt.Errorf("SaveImage: os.MkdirAll %v", err)
	}

	deckName := d.Name + fmt.Sprintf("_%vx%v_%v_%v.png", rowCount, columnCount, cardCount, pageIdx)
	file, err := os.Create(filepath.Join(resultDir, deckName))
	if err != nil {
		return "", fmt.Errorf("SaveImage: os.Create %v", err)
	}
	defer file.Close()

	err = png.Encode(file, image)
	if err != nil {
		return "", fmt.Errorf("SaveImage: png.Encode %v", err)
	}

	return deckName, nil
}

func (d *Deck) CheckCardSizes() error {
	firstCardDx := d.Cards[0].Bounds().Dx()
	firstCardDy := d.Cards[0].Bounds().Dy()

	for i := 1; i < len(d.Cards); i++ {
		if d.Cards[i].Bounds().Dx() != firstCardDx ||
			d.Cards[i].Bounds().Dy() != firstCardDy {
			return fmt.Errorf("invalid deck, cards are not same size %v", d.Name)
		}
	}

	return nil
}

func (d *Deck) GetCount() (stats *DeckStats) {
	stats = &DeckStats{
		cardWidth:  d.Cards[0].Bounds().Dx(),
		cardHeight: d.Cards[0].Bounds().Dy(),
	}

	if len(d.Cards) <= tts_maxDeckHorizontalC {
		stats.cardsRowCount = max(len(d.Cards)/tts_minDeckVerticalC, tts_minDeckHorizontalC)
		stats.cardsColCount = tts_minDeckVerticalC
		stats.pagesCount = 1

		return stats
	}

	stats.cardsRowCount = min(tts_maxWidth/stats.cardWidth, tts_maxDeckHorizontalC)
	// stats.cardsColCountPP = len(d.Cards)/stats.cardsRowCount + 1
	stats.cardsColCount = min(tts_maxHeight/stats.cardHeight, tts_maxDeckVerticalC)
	stats.pagesCount = (len(d.Cards)/stats.cardsRowCount)/tts_maxDeckVerticalC + 1

	return
}
