package errutil

import (
	"fmt"
	"os"
)

func ExitOnErr(err error) {
	if err == nil {
		return
	}
	fmt.Printf("ERROR: %v\n", err)
	os.Exit(1)
}
