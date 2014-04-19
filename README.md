# yt2mp3
A go package extracting mp3 file from YouTube URL.

# Usage
## in command line
```sh
% yt2mp3 NCdDvXg6olE
# The mp3 of http://www.youtube.com/watch?v=NCdDvXg6olE will be downloaded
% yt2mp3 -url http://www.youtube.com/watch?v=fNDrLfEfRiE
# The mp3 of http://www.youtube.com/watch?v=fNDrLfEfRiE will be downloaded
```
## in go code
```go
package main

import "yt2mp3"
import "fmt"

func main() {
    converter, _ := yt2mp3.Init()

    fname, _ := converter.Vid2mp3("5blm22DeeHY")
    // The mp3 of http://www.youtube.com/watch?v=5blm22DeeHY will be downloaded
    fmt.Println(fnam)

    fname, _ = converter.Url2mp3("http://www.youtube.com/watch?v=fNDrLfEfRiE")
    // The mp3 of http://www.youtube.com/watch?v=fNDrLfEfRiE will be downloaded
    fmt.Println(fnam)
}
```

# Setup
Install basic library
```sh
apt-get install ffmpeg
ffmpeg -version
# The output below for example
#
# ffmpeg version 0.8.10-6:0.8.10-1, Copyright (c) 2000-2013 the Libav developers
#   built on Feb  5 2014 03:52:19 with gcc 4.7.2
# ffmpeg 0.8.10-6:0.8.10-1
# libavutil    51. 22. 2 / 51. 22. 2
# libavcodec   53. 35. 0 / 53. 35. 0
# libavformat  53. 21. 1 / 53. 21. 1
# libavdevice  53.  2. 0 / 53.  2. 0
# libavfilter   2. 15. 0 /  2. 15. 0
# libswscale    2.  1. 0 /  2.  1. 0
# libpostproc  52.  0. 0 / 52.  0. 0
```
Clone repository
```sh
go get github.com/otiai10/yt2mp3
```
Enable youtube-dl command
```sh
cd $GOPATH/src/github.com/otiai10/yt2mp3
git submodule update --init
./youtube-dl/youtube-dl --version
```
Install go command bin (to use `yt2mp3` command)
```sh
go install github.com/otiai10/yt2mp3/yt2mp3
```
# Run the tests
```sh
go get github.com/r7kamura/gospel
```
```sh
go test ./tests/...
```
