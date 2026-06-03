package tools

import (
	"math/rand"
	"strconv"
)

func Gen4Number() string {
	intn := rand.Intn(9999)
	if intn < 1000 {
		intn += 1000
	}
	return strconv.Itoa(intn)
}
