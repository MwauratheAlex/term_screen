package parser

import (
	"fmt"
)

type Command struct {
	CommandByte byte
	Length      byte
	Data        []byte
}

type parser struct {
	byteQueue <-chan byte
}

func New(byteQueue <-chan byte) *parser {
	if byteQueue == nil {
		panic("Parser Requires a byteQueue")
	}
	return &parser{
		byteQueue: byteQueue,
	}
}

func (p *parser) ParseQueue() <-chan *Command {
	commandQueue := make(chan *Command, 1024)
	go func() {
		defer close(commandQueue)
		for cmdByte := range p.byteQueue {
			cmd, err := p.parseCommand(cmdByte)
			if err != nil {
				fmt.Printf("Error parsing command: %v\n", err)
				continue
			}
			commandQueue <- cmd
		}
	}()

	return commandQueue

}

func (p *parser) parseCommand(cmdByte byte) (*Command, error) {
	lengthByte, ok := <-p.byteQueue
	if !ok {
		return nil, fmt.Errorf("failed to read length byte")
	}

	data := make([]byte, lengthByte)
	for i := range int(lengthByte) {
		val, ok := <-p.byteQueue
		if !ok {
			return nil, fmt.Errorf("failed to read byte at index %d", i)
		}
		data[i] = val
	}
	return &Command{
		CommandByte: cmdByte,
		Length:      lengthByte,
		Data:        data,
	}, nil
}
