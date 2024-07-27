package imageprocessing

import (
	"errors"
	"fmt"
	"image"
)

const (
	tts_maxWidth           = 10000
	tts_maxHeight          = 10000
	tts_maxDeckHorizontalC = 10
	tts_maxDeckVerticalC   = 7
	tts_maxDeckSize        = 70
)

type Deck struct {
	DirPath string
	Cards   []image.Image
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
		imageFiles, err := GetImageFiles(deckDir)
		if err != nil {
			return nil, err
		}

		images, err := LoadImages(imageFiles)
		if err != nil {
			return nil, err
		}

		deck := &Deck{
			DirPath: deckDir,
			Cards:   images,
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
		return errors.New("error: empty decks")
	}

	return nil
}

func (d *Deck) ExportDeck(resultDir string) error {

	// check if all cards are same size Dx and Dy
	// Calculate RowC ColumnC (10k 10k)
	// Check if there is more than 1 page (70 max)

	if err := d.CheckCardSizes(); err != nil {
		return err
	}

	// rowC, colC, pageC :=

	//get deck name, substring after /

	return nil
}

func (d *Deck) CheckCardSizes() error {
	firstCardDx := d.Cards[0].Bounds().Dx()
	firstCardDy := d.Cards[0].Bounds().Dy()

	for i := 1; i < len(d.Cards); i++ {
		if d.Cards[i].Bounds().Dx() != firstCardDx ||
			d.Cards[i].Bounds().Dy() != firstCardDy {
			return fmt.Errorf("invalid deck, cards are not same size %v", d.DirPath)
		}
	}

	return nil
}

func (d *Deck) GetCount() (rowC, colC, pageC int) {
	if len(d.Cards) <= tts_maxDeckHorizontalC {
		return len(d.Cards), 1, 1
	}

	cardDx := d.Cards[0].Bounds().Dx()
	cardDy := d.Cards[0].Bounds().Dy()

	rowC = tts_maxWidth / cardDx
	colC = tts_maxHeight / cardDy

	if rowC > tts_maxDeckHorizontalC {
		rowC = tts_maxDeckHorizontalC
	}
	if colC > tts_maxDeckVerticalC {
		colC = tts_maxDeckVerticalC
	}
	pageC = len(d.Cards)/tts_maxDeckSize + 1

	return
}
