package size

import (
	"fmt"
	"strconv"
	"time"

	"github.com/c2h5oh/datasize"
)

func ReadableSize(input int64) error {
	readableSize := strconv.FormatInt(input, 10)
	fmt.Printf("input: <%v>\n", readableSize)

	byteSize := datasize.ByteSize(input).HR()
	fmt.Printf("human readable: <%v>\n", byteSize)

	fmt.Println(time.Now().Unix())
	return nil
}
