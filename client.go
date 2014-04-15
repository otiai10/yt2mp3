package yt2mp3

/**
 * Client to handle `youtube-dl` command
 * @see https://github.com/rg3/youtube-dl/blob/master/README.md
 * @example ./youtube-dl {YouTubeURL} [-x|--extract-audio] [--audio-format mp3]
 */
type DownloadClient struct {
    Command string
}
