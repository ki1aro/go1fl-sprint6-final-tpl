package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

const (
	DOT  = "."
	DASH = "-"
)

func ConvertText(text string) string {
	isMorse := strings.ContainsAny(text, DOT+DASH)
	if isMorse {
		return morse.ToText(text)
	}
	return morse.ToMorse(text)
}
