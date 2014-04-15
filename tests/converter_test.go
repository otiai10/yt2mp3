package yt2mp3

import (
	"github.com/otiai10/yt2mp3"
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestHelloYt2Mp3(t *testing.T) {
	Describe(t, "Converter.Greet", func() {
		It("should introduce itself.", func() {
            converter := yt2mp3.Init()
			Expect(converter.Greet()).To(Equal, "Hi, this is yt2mp3.Converter.")
		})
	})
}
