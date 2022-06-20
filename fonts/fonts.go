// Package fonts contains the embedded fonts and utility functions.
package fonts

import (
	"fmt"

	rice "github.com/GeertJohan/go.rice"
)

// FontFamily encapsulates a font family for easy adding to the document
type FontFamily struct {
	Name   string
	Bold   []byte
	Italic []byte
	Normal []byte
}

func loadFromRice(filename, description string) ([]byte, error) {
	box, err := riceBox()
	if err != nil {
		return nil, err
	}
	data, err := box.Bytes(filename)
	if err != nil {
		return nil, fmt.Errorf("could not load %s: %s", description, err)
	}
	return data, nil

}

func riceBox() (*rice.Box, error) {
	box, err := rice.FindBox("")
	if err != nil {
		return nil, fmt.Errorf("rice find box failed: %s", err)
	}
	return box, nil
}

// NewLatoFamily returns a new FontFamily of the Lato font.
func NewLatoFamily() (*FontFamily, error) {
	heavy, err := LatoHeavy()
	if err != nil {
		return nil, err
	}
	italic, err := LatoItalic()
	if err != nil {
		return nil, err
	}
	normal, err := LatoRegular()
	if err != nil {
		return nil, err
	}
	return &FontFamily{
		Name:   "lato",
		Bold:   heavy,
		Italic: italic,
		Normal: normal,
	}, nil
}

// LatoHeavy returns the heavy style of the Lato font.
func LatoHeavy() ([]byte, error) {
	return loadFromRice("Lato-Heavy.ttf", "lato heavy")
}

// LatoItalic returns the italic style of the Lato font.
func LatoItalic() ([]byte, error) {
	return loadFromRice("Lato-Italic.ttf", "lato italic")
}

// LatoRegular returns the regular style of the Lato font.
func LatoRegular() ([]byte, error) {
	return loadFromRice("Lato-Regular.ttf", "lato normal")
}

// NewLiberationSansFamily returns a new FontFamily of the Liberation Sans font.
func NewLiberationSansFamily() (*FontFamily, error) {
	bold, err := LiberationSansBold()
	if err != nil {
		return nil, err
	}
	italic, err := LiberationSansItalic()
	if err != nil {
		return nil, err
	}
	normal, err := LiberationSansRegular()
	if err != nil {
		return nil, err
	}
	return &FontFamily{
		Name:   "liberation-sans",
		Bold:   bold,
		Italic: italic,
		Normal: normal,
	}, nil
}

// LiberationSansBold returns the bold style of the Liberation Sans font.
func LiberationSansBold() ([]byte, error) {
	return loadFromRice("LiberationSans-Bold.ttf", "liberations sans bold")
}

// LiberationSansItalic returns the italic style of the Liberation Sans font.
func LiberationSansItalic() ([]byte, error) {
	return loadFromRice("LiberationSans-Italic.ttf", "liberations sans italic")
}

// LiberationSansRegular returns the normal style of the Liberation Sans font.
func LiberationSansRegular() ([]byte, error) {
	return loadFromRice("LiberationSans-Regular.ttf", "liberations sans regular")
}

func NewKlavikaFamily() (*FontFamily, error) {
	bold, err := KlavikaBold()
	if err != nil {
		return nil, err
	}
	italic, err := KlavikaItalic()
	if err != nil {
		return nil, err
	}
	normal, err := KlavikaRegular()
	if err != nil {
		return nil, err
	}
	return &FontFamily{
		Name:   "klavika",
		Bold:   bold,
		Italic: italic,
		Normal: normal,
	}, nil
}

func KlavikaBold() ([]byte, error) {
	return loadFromRice("klavika-bold.otf", "klavika bold")
}

func KlavikaItalic() ([]byte, error) {
	return loadFromRice("klavika-regular-italic.otf", "klavika italic")
}

func KlavikaRegular() ([]byte, error) {
	return loadFromRice("klavika-medium.otf", "klavika regular")
}
