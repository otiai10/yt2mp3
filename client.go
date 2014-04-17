package yt2mp3

import "os/exec"
import "fmt"
/**
 * Client to handle `youtube-dl` command
 * @see https://github.com/rg3/youtube-dl/blob/master/README.md
 * @example ./youtube-dl {YouTubeURL} [-x|--extract-audio] [--audio-format mp3]
 */
var (
    // TODO : とりあえず定数
    COMMAND_PATH string = "../youtube-dl/youtube-dl"
    DEFAULT_URL_PREFIX string = "http://www.youtube.com/watch?v="
)
type Client interface {
    Execute(vid string) (fpath string, err error)
}
type DownloadClient struct {
	Command string
	Options DownloadOptions
}
type DownloadOptions struct {
}

func NewDownloadClient() DownloadClient {
	return DownloadClient{
		COMMAND_PATH,
		DownloadOptions{},
	}
}

func (client DownloadClient) Execute(vid string) (fpath string, err error) {
    url := DEFAULT_URL_PREFIX + vid
    cmd := exec.Command(client.Command, url, "-x", "--audio-format", "mp3")
    // err = cmd.Run()
    output, err := cmd.Output()
    fmt.Println(string(output))
    fpath = "hoge"
    return
}
