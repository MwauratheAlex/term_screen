package processor

import (
	"fmt"
	"term_screen/parser"
	"term_screen/ui"
)

type processor struct {
	cmdQueue <-chan *parser.Command
	screen   *ui.Screen
}

func New(
	commandQueue <-chan *parser.Command,
	screen *ui.Screen,
) *processor {
	return &processor{
		cmdQueue: commandQueue,
		screen:   screen,
	}
}

func (p *processor) ProcessCommands() error {

	for cmd := range p.cmdQueue {
		if cmd == nil {
			break
		}
		switch cmd.CommandByte {
		case 0x01:
			if err := p.screen.Setup(&ui.Config{
				Width:     cmd.Data[0],
				Height:    cmd.Data[1],
				ColorMode: cmd.Data[2],
			}); err != nil {
				return fmt.Errorf("Error setting up p.screen screen: %v", err)
			}
		case 0x02:
			if !p.screen.IsSetup() {
				continue
			}
			if err := p.screen.DrawCharacter(
				&ui.Character{
					ColorIndex:  cmd.Data[2],
					DisplayChar: cmd.Data[3],
				},
				&ui.Point{
					X: cmd.Data[0],
					Y: cmd.Data[1],
				},
			); err != nil {
				return fmt.Errorf("Error drawing character: %v", err)
			}
			// fmt.Println("Character")
			// fmt.Println(string(cmd.Data[3]))
			// fmt.Println("Pos")
			// fmt.Println(int(cmd.Data[0]))
			// fmt.Println(int(cmd.Data[1]))
		case 0x03:
			if !p.screen.IsSetup() {
				continue
			}
			if err := p.screen.DrawLine(&ui.Line{
				StartPosition: &ui.Point{
					X: cmd.Data[0],
					Y: cmd.Data[1],
				},
				EndPosition: &ui.Point{
					X: cmd.Data[2],
					Y: cmd.Data[3],
				},
				Ch: &ui.Character{
					ColorIndex:  cmd.Data[4],
					DisplayChar: cmd.Data[5],
				},
			}); err != nil {
				return fmt.Errorf("Error drawing line: %v", err)
			}
		case 0x04:
			if !p.screen.IsSetup() {
				continue
			}
			if err := p.screen.RenderText(&ui.Text{
				Position: &ui.Point{
					X: cmd.Data[0],
					Y: cmd.Data[1],
				},
				ColorIndex: cmd.Data[2],
				Chars:      cmd.Data[3:],
			}); err != nil {
				return fmt.Errorf("Error rendering text: %v", err)
			}
		case 0x05:
			if !p.screen.IsSetup() {
				continue
			}
			if err := p.screen.MoveCursor(&ui.Point{
				X: cmd.Data[0],
				Y: cmd.Data[1],
			}); err != nil {
				return fmt.Errorf("Error moving cursor: %v", err)
			}
		case 0x06:
			if !p.screen.IsSetup() {
				continue
			}
			if err := p.screen.DrawAtCursor(&ui.Character{
				DisplayChar: cmd.Data[0],
				ColorIndex:  cmd.Data[1],
			}); err != nil {
				return fmt.Errorf("Error drawing at cursor position: %v", err)
			}
		case 0x07:
			if !p.screen.IsSetup() {
				continue
			}
			if err := p.screen.ClearScreen(); err != nil {
				return fmt.Errorf("Error clearing screen: %v", err)
			}
		case 0x08:
			fmt.Println("EOF reached. Thankyou for using our program.")
			fmt.Println("Press 'X' or Ctrl+C to end the program.")
		}
	}
	return nil
}
