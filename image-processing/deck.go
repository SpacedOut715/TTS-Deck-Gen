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
	tts_maxDeckHorizontalC = 10
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

func NewDeck(cards []image.Image, deckDir string) (*Deck, error) {
	name := strings.Replace(deckDir[3:], "\\", "-", -1)

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

type Decks struct {
	Decks []*Deck
}

func LoadAllDecks(rootDir string) (*Decks, error) {
	var decks []*Deck

	deckDirs, err := findEndDirectories(rootDir)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Found %v deck directories\n", len(deckDirs))

	for _, deckDir := range deckDirs {
		deckDir, err = filepath.Abs(deckDir)
		if err != nil {
			return nil, err
		}

		imageFiles, err := GetImageFiles(deckDir)
		if err != nil {
			return nil, err
		}

		images, err := LoadImages(imageFiles)
		if err != nil {
			return nil, err
		}

		deck, err := NewDeck(images, deckDir)
		if err != nil {
			return nil, err
		}

		decks = append(decks, deck)
	}

	fmt.Printf("Created %v decks\n", len(decks))

	return &Decks{
		Decks: decks,
	}, nil
}

func (d *Decks) ExportDecks(resultDir string) error {
	if d.Decks == nil || len(d.Decks) == 0 {
		return errors.New("empty decks")
	}

	err := os.MkdirAll(resultDir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	for _, deck := range d.Decks {
		err := deck.ExportDeck(resultDir)
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *Deck) ExportDeck(resultDir string) error {
	var color color.Color

	for pageIdx := 0; pageIdx < d.Stats.pagesCount; pageIdx++ {
		fmt.Println("page", pageIdx)
		startIdx := pageIdx * tts_maxDeckSize
		endIdx := pageIdx*tts_maxDeckSize + min(tts_maxDeckSize, len(d.Cards)-pageIdx*tts_maxDeckSize)
		deckSlice := d.Cards[startIdx:endIdx]

		rowC := d.Stats.cardsRowCount
		colC := len(deckSlice)/tts_maxDeckHorizontalC + min(len(deckSlice)%tts_maxDeckHorizontalC, 1) // if rounded up add 0, if not add 1
		if len(deckSlice) < d.Stats.cardsRowCount {
			rowC = len(deckSlice)
			colC = 1
		}
		image := image.NewRGBA(image.Rect(0, 0, d.Stats.cardWidth*rowC, d.Stats.cardHeight*colC))

		for deckColumnIdx := 0; deckColumnIdx < colC; deckColumnIdx++ {
			for deckRowIdx := 0; deckRowIdx < rowC; deckRowIdx++ {
				if (deckColumnIdx*rowC + deckRowIdx) >= len(deckSlice) {
					break
				}
				card := deckSlice[deckColumnIdx*rowC+deckRowIdx]
				for cardHeight := 0; cardHeight < d.Stats.cardHeight; cardHeight++ {
					for cardWidth := 0; cardWidth < d.Stats.cardWidth; cardWidth++ {
						color = card.At(cardWidth, cardHeight)
						image.Set(deckRowIdx*d.Stats.cardWidth+cardWidth, deckColumnIdx*d.Stats.cardHeight+cardHeight, color)
					}
				}
			}
		}

		file, err := os.Create(filepath.Join(resultDir, d.Name+fmt.Sprintf("_%v.png", pageIdx)))
		if err != nil {
			return err
		}
		defer file.Close()

		err = png.Encode(file, image)
		if err != nil {
			return err
		}
	}

	return nil
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
		stats.cardsRowCount = len(d.Cards)
		stats.cardsColCount = 1
		stats.pagesCount = 1

		return stats
	}

	stats.cardsRowCount = min(tts_maxWidth/stats.cardWidth, tts_maxDeckHorizontalC)
	// stats.cardsColCount = tts_maxHeight / stats.cardHeight
	stats.cardsColCount = len(d.Cards)/stats.cardsRowCount + 1

	stats.pagesCount = len(d.Cards)/tts_maxDeckSize + 1 // ili ColCount/7+1

	return
}
