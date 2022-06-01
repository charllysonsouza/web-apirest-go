package utils

import (
	"regexp"
	"strings"
)

func DocumentIsValid(document string) bool {
	len := len(document)
	if len == 11 {
		return ValidateCPF(document)
	} else if len == 14 {
		return ValidateCNPJ(document)
	} else {
		return false
	}
}

func ValidateCPF(doc string) bool {
	doc = strings.TrimSpace(doc)
	cpfRegex := `\d{11}`
	matched, _ := regexp.MatchString(cpfRegex, doc)
	return matched
}

func ValidateCNPJ(doc string) bool {
	doc = strings.TrimSpace(doc)
	cnpjRegex := `\d{14}`
	matched, _ := regexp.MatchString(cnpjRegex, doc)
	return matched
}
