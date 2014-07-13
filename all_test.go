package yt2mp3_test

import "github.com/otiai10/yt2mp3"
import "testing"
import . "github.com/otiai10/mint"

func TestInit(t *testing.T) {
    converter, err := yt2mp3.Init()
    Expect(t, err).ToBe(nil)
    Expect(t, converter).TypeOf("*yt2mp3.Converter")
}
