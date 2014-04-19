package main

import "github.com/otiai10/yt2mp3"
import "fmt"
import "os"

import "flag"

var (
	YOUTUBE_VID_FORMAT = "^[A-Za-z0-9_-]{11}$"
	ERR_PREFIX         = "[yt2mp3]"
	ERR_SUFFIX         = "Please type `yt2mp3 help`."
	NOT_ENOUGH_ARGS    = "Not enough arguments!"
	INVALID_VID_FORMAT = "Invalid vid format."

	url = ""
)

func init() {
	_url := flag.String("url", "", "URL to convert, for example `http://www.youtube.com/watch?v=fNDrLfEfRiE`")
	flag.Parse()
	url = *_url
}

func main() {
	converter, err := yt2mp3.Init()
	if err != nil {
		fmt.Println(err.Error())
	}
	var fname string

	if url != "" {
		fname, err = converter.Url2mp3(url)
	} else {
		fname, err = converter.Vid2mp3(os.Args[1])
	}

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(fname)
}
