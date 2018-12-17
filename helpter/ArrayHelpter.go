package helpter

import (
	"math/rand"
	"time"
)

type ArrayHelper struct {

}

func (A *ArrayHelper)RandInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - min) + min
}