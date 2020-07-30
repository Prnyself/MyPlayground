package channel

import "fmt"

func ReadFromBufferChan() {
	bufferChan := make(chan error, 10)
	close(bufferChan)
	for value := range bufferChan {
		fmt.Println(value)
	}
}
