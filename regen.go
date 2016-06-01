package regen

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Generate return a regexp for extract field from raw
func Generate(raw, expected string) (expr string, err error) {
	// TODO  multi occurs
	from := strings.Index(raw, expected)
	if from < 0 {
		return "", fmt.Errorf(`"%s" not found in "%s"`, expected, raw)
	}
	to := from + len(expected)

	prefix := raw[:from]
	body := raw[from:to]
	suffix := raw[to:]
	if len(suffix) == 0 {
		return genPfx(prefix) + "(.*)$", nil
	}

	r, _ := utf8.DecodeRuneInString(suffix)
	if !strings.ContainsRune(body, r) && (unicode.IsSymbol(r) || unicode.IsPunct(r) || unicode.IsSpace(r)) {
		return genPfx(prefix) + fmt.Sprintf(`([^%[1]s]+)%[1]s`, genSfx(r)), nil
	}
	return genPfx(prefix) + "(" + genBody(body) + ")" + genSfx(r), nil
}

// GenerateAndVerify generate a regexp and then verify does it match field correctly
func GenerateAndVerify(raw, expected string) (string, error) {
	pattern, err := Generate(raw, expected)
	if err != nil {
		return "", err
	}

	re, err := regexp.Compile(pattern)
	if err != nil {
		return pattern, fmt.Errorf("the pattern cannot be parsed")
	}

	matches := re.FindStringSubmatch(raw)
	if len(matches) < 2 || matches[1] != expected {
		return pattern, fmt.Errorf("the pattern does not match the field")
	}
	return pattern, nil
}

func genPfx(raw string) string {
	lastIndexOfPunct := strings.LastIndexFunc(raw, func(ch rune) bool {
		return unicode.IsSymbol(ch) || unicode.IsPunct(ch) || (unicode.IsSpace(ch) && ch != ' ')
	})

	var pattern bytes.Buffer
	if lastIndexOfPunct >= 0 {
		lastPunct, _ := utf8.DecodeRuneInString(raw[lastIndexOfPunct:])
		punctCount := strings.Count(raw, string(lastPunct))
		pattern.WriteString(fmt.Sprintf(`(?:.*?\%s){%d}`, string(lastPunct), punctCount))
	}

	pattern.WriteString(genBody(raw[lastIndexOfPunct+1:]))
	return pattern.String()
}

func genBody(raw string) string {
	var (
		pattern                   bytes.Buffer
		previous, current, suffix string
	)

	for _, r := range raw {
		switch {
		case unicode.IsDigit(r):
			current = `\d`
		case unicode.IsLetter(r):
			current = `\w`
		case unicode.IsSpace(r):
			current = `\s`
		case unicode.IsSymbol(r) || unicode.IsPunct(r):
			current = fmt.Sprintf(`\%s`, string(r))
		default:
			current = string(r)
		}

		if previous != current {
			pattern.WriteString(current)
			suffix = current
		} else if suffix != "+" {
			pattern.WriteString("+")
			suffix = "+"
		}

		previous = current
	}
	return pattern.String()
}

func genSfx(term rune) string {
	switch {
	case unicode.IsDigit(term) || unicode.IsLetter(term):
		return fmt.Sprintf(`%s`, string(term))
	case unicode.IsSpace(term):
		return `\s`
	default:
		return fmt.Sprintf(`\%s`, string(term))
	}
}
