package yt2mp3

import "os/exec"
import "fmt"
import "regexp"
import "path"
import "runtime"

/**
 * Client to handle `youtube-dl` command
 * @see https://github.com/rg3/youtube-dl/blob/master/README.md
 * @example ./youtube-dl {YouTubeURL} [-x|--extract-audio] [--audio-format mp3]
 */
var (
	// TODO : とりあえず定数
	COMMAND_PATH       string = "/youtube-dl/youtube-dl"
	DEFAULT_URL_PREFIX string = "http://www.youtube.com/watch?v=%s"
)

type Client interface {
	Execute(vid string) (fpath string, err error)
}
type DownloadClient struct {
	CommandBase string
	Options     DownloadOptions
	command     *exec.Cmd
}
type DownloadOptions struct {
}

func NewDownloadClient() DownloadClient {
	_, file, _, _ := runtime.Caller(0)
	commandBase := path.Dir(file) + COMMAND_PATH
	return DownloadClient{
		commandBase,
		DownloadOptions{},
		nil,
	}
}

func (client DownloadClient) Execute(vid string) (fname string, err error) {
	url := fmt.Sprintf(DEFAULT_URL_PREFIX, vid)
	client.command = exec.Command(client.CommandBase, url, "-x", "--audio-format", "mp3")
	output := client.executeCommand()
	fname = client.extraceFileName(output)
	return
}

func (client DownloadClient) executeCommand() (output string) {
	output_b, _ := client.command.Output()
	output = string(output_b)
	return
}
func (client DownloadClient) extraceFileName(output string) (fname string) {
	exp := regexp.MustCompile(`\[ffmpeg\] Destination: (.+\.mp3)\n`)
	matches := exp.FindStringSubmatch(output)
	fname = matches[1]
	return
}
