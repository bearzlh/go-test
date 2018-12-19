package helper

import (
	"math/rand"
	"time"
)

type ArrayHelper struct {

}

var HA = ArrayHelper{}

func (A *ArrayHelper)RandInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - min) + min
}