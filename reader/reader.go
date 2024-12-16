package reader

import (
	"io"
)

func Read(input io.Reader) <-chan byte {
	byteQueue := make(chan byte, 1024)
	go func() {
		defer close(byteQueue)
		buf := make([]byte, 1024)
		for {
			n, err := input.Read(buf)
			if n > 0 {
				for _, b := range buf[:n] {
					byteQueue <- b
				}
			}
			if err != nil {
				if err == io.EOF {
					break
				}
				panic(err)
			}
		}
	}()
	return byteQueue
}
