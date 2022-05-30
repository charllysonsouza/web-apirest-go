package utils

import (
	"regexp"
	"strings"
)

func ConvertToCents(amount float64) int {
	return int(amount * 100)
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
