package yt2mp3

import (
	"github.com/otiai10/yt2mp3"
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestYt2mp3(t *testing.T) {
	Describe(t, "Init", func() {
		Context("with correct environments", func() {
			It("should give converter.", func() {
				converter, err := yt2mp3.Init()
				Expect(err).To(NotExist)
				Expect(converter).To(Exist)
			})
		})
	})
}
