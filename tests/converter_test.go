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

			// TODO: mocking
			converter, _ := yt2mp3.Init()

			vid := "NCdDvXg6olE"
			mp3, err := converter.Vid2mp3(vid)
			fmt.Println(mp3, err)
			// TODO: more assertions
		})
	})
	Describe(t, "Url2mp3", func() {
		It("should convert URL to mp3.", func() {

			// TODO: mocking
			converter, _ := yt2mp3.Init()

			url := "http://www.youtube.com/watch?v=fNDrLfEfRiE"
			mp3, err := converter.Url2mp3(url)
			fmt.Println(mp3, err)
			// TODO: more assertions
		})
	})
}
