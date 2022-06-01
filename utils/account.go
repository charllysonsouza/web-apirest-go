package utils

import (
	"math/rand"
	"strconv"
	"time"
)

func ConvertToCents(amount float64) int {
	return int(amount * 100)
}

func GenerateAccountNumber() string {
	rand.Seed(time.Now().UnixNano())
	min := 10000
	max := 99999
	accountNum := rand.Intn(max-min) + min
	return strconv.Itoa(accountNum)
}

func GenerateAgencyNumber() string {
	rand.Seed(time.Now().UnixNano())
	min := 1000
	max := 9999
	agencyNum := rand.Intn(max-min) + min
	return strconv.Itoa(agencyNum)
}
