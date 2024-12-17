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

type Screen struct {
	config             *Config
	startingCursorPos  *Point
	currCursorPosition *Point
}

func NewScreen() *Screen {
	return &Screen{}
}

func (s *Screen) Setup(config *Config) error {
	if config.Width <= 0 || config.Height <= 0 {
		return fmt.Errorf("Width and Height must be creater than 0")
	}

	if config.ColorMode != 0x00 && config.ColorMode != 0x01 && config.ColorMode != 0x02 {
		var sb strings.Builder
		sb.WriteString("Invalid color mode.")
		sb.WriteString("Wanted: 0x00 -> monochrome OR 0x01 - 16 colors OR 0x02 - 256 colors")
		return fmt.Errorf(sb.String())
	}

	s.config = config
	s.startingCursorPos = &Point{5, 2}
	s.ClearScreen()
	s.drawBorder()

	return nil
}

// DrawCharacter draws a character <ch> at position <pos>
func (s *Screen) DrawCharacter(ch *Character, pos *Point) error {
	if pos == nil {
		return fmt.Errorf("Cannot draw character at <nil> position")
	}
	if err := s.MoveCursor(pos); err != nil {
		return err
	}
	fmt.Printf("%s", string(ch.DisplayChar))
	s.resetStyling()
	fmt.Println() // prevents the showing of the percentage sign at cursor pos
	return nil
}

// DrawLine draws a line on the screen
func (s *Screen) DrawLine(line *Line) error {
	start := line.StartPosition
	end := line.EndPosition

	for x := start.X; x <= end.X; x++ {
		for y := start.Y; y <= end.Y; y++ {
			s.DrawCharacter(line.Ch, &Point{x, y})
		}
	}

	return nil
}

func (s *Screen) RenderText(text *Text) error {
	return nil
}

// MoveCursor moves the cursor to a specific position (pos.X, pos.Y)
func (s *Screen) MoveCursor(pos *Point) error {
	if pos == nil {
		return fmt.Errorf("Cannot move cursor to <nil> position")
	}
	fmt.Printf("\033[%d;%dH", pos.Y, pos.X)
	s.currCursorPosition = pos
	return nil
}

// DrawAtCursor draws a character at the current cursor position
func (s *Screen) DrawAtCursor(ch *Character) error {
	return s.DrawCharacter(ch, s.currCursorPosition)
}

// ClearScreen deletes everything on the screen
func (s *Screen) ClearScreen() error {
	fmt.Print("\033[2J")
	return nil
}

func (s *Screen) IsSetup() bool {
	return s.config != nil
}

// resetStyling resets any styling/color
func (s *Screen) resetStyling() {
	fmt.Printf("\033[0m")
}

// drawBorder draws a border of the working area based on the screen dimensions
func (s *Screen) drawBorder() {
	borderCharacter := &Character{
		ColorIndex:  0x02,
		DisplayChar: byte('*'),
	}

	// top
	s.DrawLine(&Line{
		StartPosition: s.startingCursorPos,
		EndPosition:   s.startingCursorPos.Add(&Point{s.config.Width, 0}),
		Ch:            borderCharacter,
	})
	// bottom
	s.DrawLine(&Line{
		StartPosition: s.startingCursorPos.Add(&Point{0, s.config.Height}),
		EndPosition: s.startingCursorPos.Add(
			&Point{s.config.Width, s.config.Height}),
		Ch: borderCharacter,
	})
	// left
	s.DrawLine(&Line{
		StartPosition: s.startingCursorPos,
		EndPosition:   s.startingCursorPos.Add(&Point{0, s.config.Height}),
		Ch:            borderCharacter,
	})
	// right
	s.DrawLine(&Line{
		StartPosition: s.startingCursorPos.Add(&Point{s.config.Width, 0}),
		EndPosition: s.startingCursorPos.Add(
			&Point{s.config.Width, s.config.Height}),
		Ch: borderCharacter,
	})
}

type Character struct {
	ColorIndex  byte
	DisplayChar byte
}

type Line struct {
	StartPosition *Point
	EndPosition   *Point
	Ch            *Character
}

type Text struct {
	Position   *Point
	ColorIndex byte
	Chars      []byte
}
