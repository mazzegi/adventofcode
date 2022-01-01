package slice

import (
	"fmt"

	"github.com/pkg/errors"
)

type Claim struct {
	ID     int
	Left   int
	Top    int
	Width  int
	Height int
}

func ParseClaim(s string) (Claim, error) {
	var c Claim
	_, err := fmt.Sscanf(s, "#%d @ %d,%d: %dx%d", &c.ID, &c.Left, &c.Top, &c.Width, &c.Height)
	if err != nil {
		return Claim{}, errors.Wrapf(err, "scan claim %q", s)
	}
	return c, nil
}

func (c Claim) Contains(x, y int) bool {
	return x >= c.Left && x < c.Left+c.Width &&
		y >= c.Top && y < c.Top+c.Height

}
