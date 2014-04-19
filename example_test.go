package yt2mp3

import (
	"fmt"
	"github.com/otiai10/yt2mp3"
)

func ExampleVid2mp3() {
	converter, _ := yt2mp3.Init()
	fname, _ := converter.Vid2mp3("fNDrLfEfRiE")
	fmt.Printf("Downloaded %s", fname)
}
