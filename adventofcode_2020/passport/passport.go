package passport

import (
	"strings"

	"github.com/pkg/errors"
)

type Key string

const (
	BirthYear      Key = "byr"
	IssueYear      Key = "iyr"
	ExpirationYear Key = "eyr"
	Height         Key = "hgt"
	HairColor      Key = "hcl"
	EyeColor       Key = "ecl"
	PassportID     Key = "pid"
	CountryID      Key = "cid"
)

func (k Key) Validators() []ValueValidator {
	switch k {
	case BirthYear:
		return []ValueValidator{NewDigitCountValidator(4), NewRangeValidator(1920, 2002)}
	case IssueYear:
		return []ValueValidator{NewDigitCountValidator(4), NewRangeValidator(2010, 2020)}
	case ExpirationYear:
		return []ValueValidator{NewDigitCountValidator(4), NewRangeValidator(2020, 2030)}
	case Height:
		return []ValueValidator{NewHeightValidator()}
	case HairColor:
		return []ValueValidator{NewHexColorValidator()}
	case EyeColor:
		return []ValueValidator{NewEyeColorValidator()}
	case PassportID:
		return []ValueValidator{NewDigitCountValidator(9), NewNumberValidator()}
	default:
		return []ValueValidator{}
	}
}

func (k Key) IsValidValue(s string) bool {
	vs := k.Validators()
	for _, val := range vs {
		if !val.IsValid(s) {
			return false
		}
	}
	return true
}

func AllKeys() []Key {
	return []Key{BirthYear, IssueYear, ExpirationYear, Height, HairColor, EyeColor, PassportID, CountryID}
}

func MandatoryKeys() []Key {
	return []Key{BirthYear, IssueYear, ExpirationYear, Height, HairColor, EyeColor, PassportID}
}

func (k Key) IsValid() bool {
	for _, ak := range AllKeys() {
		if ak == k {
			return true
		}
	}
	return false
}

type Passport map[Key]string

func (p Passport) IsValid() bool {
	for _, mk := range MandatoryKeys() {
		if _, ok := p[mk]; !ok {
			return false
		}
	}
	for k, v := range p {
		if !k.IsValidValue(v) {
			return false
		}
	}
	return true
}

func ParseStrings(sl []string) (Passport, error) {
	p := Passport{}
	for _, l := range sl {
		sl := strings.Split(l, " ")
		for _, s := range sl {
			s = strings.Trim(s, " \r\n\t")
			fields := strings.Split(s, ":")
			if len(fields) != 2 {
				return p, errors.Errorf("invalid field %q", s)
			}
			key := Key(fields[0])
			val := fields[1]
			if !key.IsValid() {
				return p, errors.Errorf("invalid key %q", key)
			}
			p[key] = val
		}
	}
	return p, nil
}
