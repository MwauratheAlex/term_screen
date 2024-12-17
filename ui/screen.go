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
	config        *Config
	origin        *Point // this will be our (0, 0)
	currCursorPos *Point
}

func NewScreen() *Screen {
	return &Screen{}
}

// Setup initializes the Screen
func (s *Screen) Setup(config *Config) error {
	if config.Width <= 0 || config.Height <= 0 {
		return fmt.Errorf("Width and Height must be creater than 0")
	}

	if config.ColorMode != 0x00 &&
		config.ColorMode != 0x01 && config.ColorMode != 0x02 {
		var sb strings.Builder
		sb.WriteString("Invalid color mode.")
		sb.WriteString("Wanted: 0x00 -> monochrome OR 0x01 - 16 colors OR 0x02 - 256 colors")
		return fmt.Errorf(sb.String())
	}

	s.config = config
	s.origin = &Point{5, 2}
	s.ClearScreen()
	s.drawBorder()

	return nil
}

// DrawCharacter performs the drawing of the character <ch> at the position <pos>
func (s *Screen) DrawCharacter(ch *Character, pos *Point) error {
	if pos == nil {
		return fmt.Errorf("Cannot draw character at <nil> position")
	}
	if err := s.MoveCursor(pos); err != nil {
		return err
	}
	fmt.Printf("%s", string(ch.DisplayChar))
	s.reset()
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
	// always move in reference to origin
	newPos := s.origin.Add(pos)
	fmt.Printf("\033[%d;%dH", newPos.Y, newPos.X)
	s.currCursorPos = pos
	return nil
}

// DrawAtCursor draws a character at the current cursor position
func (s *Screen) DrawAtCursor(ch *Character) error {
	return s.DrawCharacter(ch, s.currCursorPos)
}

// ClearScreen deletes everything on the screen
func (s *Screen) ClearScreen() error {
	fmt.Print("\033[2J")
	return nil
}

// IsSetup returns true or false depencing on wheter screen has been setup
func (s *Screen) IsSetup() bool {
	return s.config != nil
}

// reset resets the terminal screen
func (s *Screen) reset() {
	// reset any styling/color
	fmt.Printf("\033[0m")
	// move cursor to final resting position
	// this ensures borders are not deleted.
	// e.g. if you draw a char above bottom border and leaving the cursor there,
	// the rest of the bottom borders are deleted
	// This ensures cursor is always at bottom border + 1 hence nothing in
	// our screen gets deleted.
	s.MoveCursor(&Point{0, s.config.Height})
	// prevent '%' from appearing at cursor position
	fmt.Println()
}

// drawBorder draws a border of the working area based on the screen dimensions
func (s *Screen) drawBorder() {
	borderCharacter := &Character{
		ColorIndex:  0x02,
		DisplayChar: byte('*'),
	}
	// top
	s.DrawLine(&Line{
		StartPosition: &Point{0, 0},
		EndPosition:   &Point{s.config.Width, 0},
		Ch:            borderCharacter,
	})
	// bottom
	s.DrawLine(&Line{
		StartPosition: &Point{0, s.config.Height},
		EndPosition:   &Point{s.config.Width, s.config.Height},
		Ch:            borderCharacter,
	})
	// left
	s.DrawLine(&Line{
		StartPosition: &Point{0, 0},
		EndPosition:   &Point{0, s.config.Height},
		Ch:            borderCharacter,
	})
	// right
	s.DrawLine(&Line{
		StartPosition: &Point{s.config.Width, 0},
		EndPosition:   &Point{s.config.Width, s.config.Height},
		Ch:            borderCharacter,
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
