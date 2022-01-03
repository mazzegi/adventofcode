package testutil

import "fmt"

func TestName(i int) string {
	return fmt.Sprintf("test #%02d", i)
}
