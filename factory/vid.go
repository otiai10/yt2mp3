package factory

import "regexp"
import "errors"

var (
	url_youtube_com = "www.youtube.com"
	url_youtu_be    = "youtu.be"
	e_url_too_short = "Too short URL format."
)

// TODO: separate
type regexx struct {
	*regexp.Regexp
}

func (re *regexx) SubmatchMap(str string) (mapped map[string]string) {
	mapped = make(map[string]string)
	names := re.SubexpNames()
	names[0] = "_origin"
	for i, match := range re.FindStringSubmatch(str) {
		mapped[names[i]] = match
	}
	return
}

func Url2vid(url string) (vid string, err error) {
	exp := regexx{
		regexp.MustCompile(
			// TODO: consturct dynamically
			`^(?P<protocol>https?)://(?P<domain>www.youtube.com|youtu.be)/(?P<query>.+)`,
		),
	}
	matches := exp.SubmatchMap(url)
	if len(matches) < 4 {
		err = errors.New(e_url_too_short)
		return
	}
	switch matches["domain"] {
	case url_youtube_com:
		vid, err = vidFromYoutubeCom(matches["query"])
		return
	case url_youtu_be:
		vid, err = vidFromYoutuBe(matches["query"])
		return
	}
	return
}

func vidFromYoutubeCom(query string) (vid string, err error) {
	exp := regexx{
		regexp.MustCompile(
			`^watch?.*v=(?P<vid>[a-zA-Z0-9_-]{11}).*`,
		),
	}
	// TODO: DRY
	matches := exp.SubmatchMap(query)
	if matches["vid"] != "" {
		vid = matches["vid"]
		return
	}
	// TODO: error
	return
}

func vidFromYoutuBe(query string) (vid string, err error) {
	exp := regexx{
		regexp.MustCompile(
			`^(?P<vid>[a-zA-Z0-9_-]{11}).*`,
		),
	}
	// TODO: DRY
	matches := exp.SubmatchMap(query)
	if matches["vid"] != "" {
		vid = matches["vid"]
		return
	}
	// TODO: error
	return
}
