package yt2mp3

import "fmt"

type Converter struct {
}

func Init() *Converter {
    return &Converter{}
}

func (c *Converter)Greet() string {
    return fmt.Sprintf("Hi, this is %s.", "yt2mp3.Converter")
}
