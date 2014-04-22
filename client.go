package yt2mp3

import "os/exec"
import "fmt"
import "regexp"
import "path"
import "runtime"

var (
	// TODO : optionalize
	command_path         string = "/youtube-dl/youtube-dl"
	default_url_prefix   string = "http://www.youtube.com/watch?v=%s"
	default_output_dir   string = "downloads/{id}/"
	default_output_fname string = "{title}.{ext}"
)

// Client to handle `youtube-dl` command
// @see https://github.com/rg3/youtube-dl/blob/master/README.md
// @example ./youtube-dl {YouTubeURL} [-x|--extract-audio] [--audio-format mp3]
// @example ./youtube-dl/youtube-dl -x http://www.youtube.com/watch?v=-Xo7qAs1sjM --audio-format mp3 -o './downloads/%(id)s/%(title)s.%(ext)s'

type Client interface {
	Execute(vid string) (fpath string, err error)
	SetOpt(key string, value interface{}) (ok bool, err error)
}
type DownloadClient struct {
	CommandBase string
	Options     DownloadOptions
	command     *exec.Cmd
}
type DownloadOptions struct {
	OutputDir string
}

func NewDownloadClient() *DownloadClient {
	_, file, _, _ := runtime.Caller(0)
	commandBase := path.Dir(file) + command_path
	return &DownloadClient{
		commandBase,
		DownloadOptions{
			OutputDir: default_output_dir,
		},
		nil,
	}
}

func (client DownloadClient) Execute(vid string) (fname string, err error) {
	url := fmt.Sprintf(default_url_prefix, vid)
	client.command = exec.Command(
		client.CommandBase,
		url,
		"-x",
		"--audio-format",
		"mp3",
		"-o",
		client.getDestinationFullPath(),
	)
	output := client.executeCommand()
	fname = client.extraceFileName(output)
	return
}
func (client *DownloadClient) SetOpt(key string, value interface{}) (ok bool, err error) {
	if key == "outputdir" {
		dirname, _ok := value.(string)
		ok = _ok
		client.Options.OutputDir = dirname
	}
	return
}
func (client DownloadClient) getDestinationFullPath() (path string) {
	return client.compilePath(client.Options.OutputDir + default_output_fname)
}
func (client DownloadClient) compilePath(rawFilePath string) (path string) {
	exp := regexp.MustCompile("\\{id\\}")
	path = exp.ReplaceAllString(rawFilePath, "%(id)s")
	exp = regexp.MustCompile("\\{title\\}")
	path = exp.ReplaceAllString(path, "%(title)s")
	exp = regexp.MustCompile("\\{ext\\}")
	path = exp.ReplaceAllString(path, "%(ext)s")
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
	if len(matches) < 2 {
		// TODO: error
		return
	}
	fname = matches[1]
	return
}
