package main

import "github.com/otiai10/yt2mp3"
import "fmt"
import "os"
import "errors"
import "regexp"

var (
	YOUTUBE_VID_FORMAT = "^[A-Za-z0-9_-]{11}$"
	ERR_PREFIX         = "[yt2mp3]"
	ERR_SUFFIX         = "Please type `yt2mp3 help`."
	NOT_ENOUGH_ARGS    = "Not enough arguments!"
	INVALID_VID_FORMAT = "Invalid vid format."
)

func main() {
	converter, err := yt2mp3.Init()
	if err != nil {
		fmt.Println(err.Error())
	}
	err = validate()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var fname string
	fname, err = converter.Vid2mp3(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(fname)
}

func validate() (err error) {
	if len(os.Args) < 2 {
		return errors.New(errorMessage(NOT_ENOUGH_ARGS))
	}
	exp, _ := regexp.Compile(YOUTUBE_VID_FORMAT)
	if !exp.MatchString(os.Args[1]) {
		return errors.New(errorMessage(INVALID_VID_FORMAT))
	}
	return
}

func errorMessage(message string) string {
	return fmt.Sprintf("%s %s %s", ERR_PREFIX, message, ERR_SUFFIX)
}
