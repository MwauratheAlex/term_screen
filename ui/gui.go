package ui

import (
	"fmt"
	"strings"
)

type Config struct {
	Width     byte
	Height    byte
	ColorMode byte //0x00 for monochrome, 0x01 for 16 colors, 0x02 for 256
}

type Ui struct {
	config    *Config
	cursorPos *Point
}

func New() *Ui {
	return &Ui{}
}

func (u *Ui) Setup(config *Config) error {
	if config.Width <= 0 || config.Height <= 0 {
		return fmt.Errorf("Width and Height must be creater than 0")
	}

	if config.ColorMode != 0x00 && config.ColorMode != 0x01 && config.ColorMode != 0x02 {
		var sb strings.Builder
		sb.WriteString("Invalid color mode.")
		sb.WriteString("Wanted: 0x00 -> monochrome OR 0x01 - 16 colors OR 0x02 - 256 colors")
		return fmt.Errorf(sb.String())
	}

	u.config = config

	fmt.Println("Setup Complete")
	return nil
}

func (u *Ui) DrawCharacter(ch *Character, pos *Point) error {
	fmt.Printf("Drawing character: %s\n", string(ch.DisplayChar))
	return nil
}

func (u *Ui) DrawLine(line *Line) error {
	fmt.Println("Drawing Line")
	return nil
}

func (u *Ui) RenderText(text *Text) error {
	fmt.Println("Rendering text")
	return nil
}

func (u *Ui) MoveCursor(point *Point) error {
	fmt.Println("Moving cursor")
	return nil
}

func (u *Ui) DrawAtCursor(ch *Character) error {
	fmt.Println("Drawing at character cursor")
	u.DrawCharacter(ch, u.cursorPos)
	return nil
}

func (u *Ui) ClearScreen() error {
	fmt.Println("Clearing the screen")
	return nil
}

func (u *Ui) IsSetup() bool {
	return u.config != nil
}

type Point struct {
	X byte
	Y byte
}

type Character struct {
	ColorIndex  byte
	DisplayChar byte
}

type Line struct {
	StartPosition *Point
	EndPosition   *Point
	ColorIndex    byte
	Character     byte
}

type Text struct {
	Position   *Point
	ColorIndex byte
	Chars      []byte
}
