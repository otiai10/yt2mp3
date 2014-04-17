package yt2mp3

import (
	"github.com/otiai10/yt2mp3"
	. "github.com/r7kamura/gospel"
	"testing"

    "fmt"
)

func TestConverter(t *testing.T) {
	Describe(t, "Vid2mp3", func() {
        It("should convert video id to mp3.", func() {

            vid := "NCdDvXg6olE"
            converter, _ := yt2mp3.Init()
            mp3, err := converter.Vid2mp3(vid)
            fmt.Println(mp3, err)
        })
	})
}
