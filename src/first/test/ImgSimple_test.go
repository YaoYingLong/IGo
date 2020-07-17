package test

import (
	"first/other"
	"math/rand"
	"os"
	"testing"
	"time"
)

func TestGif(t *testing.T) {
	rand.Seed(time.Now().UTC().UnixNano())
	other.Lissajous(os.Stdout)
}
