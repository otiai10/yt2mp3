package yt2mp3

import "github.com/otiai10/yt2mp3/factory"
import . "github.com/r7kamura/gospel"
import "testing"

func TestUrl2Vid(t *testing.T) {
	Describe(t, "Url2vid", func() {
		Context("with valid `www.youtube.com` URL", func() {
			It("should give vid from query `?v=xxxxxxxxxxx`", func() {

				url := "https://www.youtube.com/watch?v=NCdDvXg6olE"
				vid, err := factory.Url2vid(url)
				Expect(vid).To(Equal, "NCdDvXg6olE")
				Expect(err).To(NotExist)

				url = "http://www.youtube.com/watch?v=fNDrLfEfRiE"
				vid, err = factory.Url2vid(url)
				Expect(vid).To(Equal, "fNDrLfEfRiE")
				Expect(err).To(NotExist)

			})
		})
		Context("with valid `youtu.be` URL", func() {
			It("should give vid from query `/xxxxxxxxxxx`", func() {

				url := "http://youtu.be/NCdDvXg6olE"
				vid, err := factory.Url2vid(url)
				Expect(vid).To(Equal, "NCdDvXg6olE")
				Expect(err).To(NotExist)

				url = "https://youtu.be/NCdDvXg6olE"
				vid, err = factory.Url2vid(url)
				Expect(vid).To(Equal, "NCdDvXg6olE")
				Expect(err).To(NotExist)

			})
		})
	})
}
