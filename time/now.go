package time

import (
	"fmt"
	"time"
)

func Now() {
	fmt.Println(time.Now().Unix())
}
