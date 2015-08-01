# resourcelimit

#how to use

go get github.com/conc/resourcelimit

#example

```golang
package main

import (
	"github.com/conc/resourcelimit"
	"log"
	"time"
)

func main() {
	limit := resourcelimit.NewLimit(4)
	for i := 0; i < 10; i++ {
		limit.UseResource()
		go gosleep(i, limit)
	}

	time.Sleep(4 * time.Second)
}

func gosleep(serial int, limit *resourcelimit.SimpleLimit) {
	defer func() {
		limit.FreeResource()
	}()

	log.Println(serial, ": begin")
	time.Sleep(1 * time.Second)
	log.Println(serial, ": end")
	return
}
```