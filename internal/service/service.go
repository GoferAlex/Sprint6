package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

var ErrEmptyString = errors.New("string is empty")

func Convert(s string) (string, error) {
	var res string
	if num := strings.IndexAny(s, ".-"); num == 0 {
		res = morse.ToText(s)
		if res == "" {
			return "", ErrEmptyString
		} else {
			return res, nil
		}
	}
	res = morse.ToMorse(s)
	if res == "" {
		return "", ErrEmptyString
	} else {
		return res, nil
	}
}
