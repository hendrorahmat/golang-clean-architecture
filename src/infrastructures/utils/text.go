package utils

import (
	"regexp"
	"strings"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
var space = regexp.MustCompile(`\s+`)
var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

func clearString(str string) string {
	return nonAlphanumericRegex.ReplaceAllString(str, "")
}

func ToSnakeCase(str string) string {
	snake := clearString(str)
	snake = matchFirstCap.ReplaceAllString(snake, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	snake = space.ReplaceAllString(snake, "")
	return strings.ToLower(snake)
}

func ToKebabCase(str string) string {
	snake := clearString(str)
	snake = space.ReplaceAllString(snake, "-")

	return strings.ToLower(snake)
}
