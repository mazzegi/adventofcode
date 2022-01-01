package passport

import (
	"encoding/hex"
	"strconv"
	"strings"
)

type ValueValidator interface {
	IsValid(s string) bool
}

//
type DigitCountValidator struct {
	cnt int
}

func NewDigitCountValidator(cnt int) DigitCountValidator {
	return DigitCountValidator{cnt: cnt}
}

func (v DigitCountValidator) IsValid(s string) bool {
	return len(s) == v.cnt
}

//
type RangeValidator struct {
	min, max int64
}

func NewRangeValidator(min, max int64) RangeValidator {
	return RangeValidator{
		min: min,
		max: max,
	}
}

func (v RangeValidator) IsValid(s string) bool {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return false
	}
	return n >= v.min && n <= v.max
}

//
type HeightValidator struct {
}

func NewHeightValidator() HeightValidator {
	return HeightValidator{}
}

func (v HeightValidator) IsValid(s string) bool {
	switch {
	case strings.HasSuffix(s, "cm"):
		return NewRangeValidator(150, 193).IsValid(strings.TrimSuffix(s, "cm"))
	case strings.HasSuffix(s, "in"):
		return NewRangeValidator(59, 76).IsValid(strings.TrimSuffix(s, "in"))
	default:
		return false
	}
}

//
type HexColorValidator struct {
}

func NewHexColorValidator() HexColorValidator {
	return HexColorValidator{}
}

func (v HexColorValidator) IsValid(s string) bool {
	if len(s) != 7 {
		return false
	}
	if !strings.HasPrefix(s, "#") {
		return false
	}
	_, err := hex.DecodeString(strings.TrimPrefix(s, "#"))
	return err == nil
}

//
type EyeColorValidator struct {
}

func NewEyeColorValidator() EyeColorValidator {
	return EyeColorValidator{}
}

func (v EyeColorValidator) IsValid(s string) bool {
	switch s {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	default:
		return false
	}
}

//
type NumberValidator struct {
}

func NewNumberValidator() NumberValidator {
	return NumberValidator{}
}

func (v NumberValidator) IsValid(s string) bool {
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}
