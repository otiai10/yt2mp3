package yt2mp3

import (
	"github.com/otiai10/yt2mp3"
	. "github.com/r7kamura/gospel"
	"testing"

	"fmt"
)

func TestClient(t *testing.T) {
	Describe(t, "SetOpt", func() {
		It("should change option value.", func() {
			client := yt2mp3.NewDownloadClient()
			client.SetOpt("outputdir", "~/unko/{ext}/")
			fname, e := client.Execute("NCdDvXg6olE")
			// TODO: assertion
			fmt.Println(fname, e)
		})
	})
}
