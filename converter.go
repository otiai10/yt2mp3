package yt2mp3

type Converter struct {
	Client Client
}
type MyError struct {
	message string
}

func (e MyError) Error() string {
	return e.message
}
func NewError(message string) error {
	return MyError{
		message,
	}
}

func Init(args ...interface{}) (converter *Converter, err error) {
	if len(args) > 0 {
		if _, ok := interface{}(args[0]).(DownloadClient); ok {
			converter = &Converter{
				args[0].(DownloadClient),
			}
		} else {
			err = NewError(
				"Invalid argument for Init: type `DownloadClient` required",
			)
		}
		return
	}
	err = CheckEnv()
	if err == nil {
		converter = &Converter{
			NewDownloadClient(),
		}
	}
	return
}

func CheckEnv() (err error) {
	return err
}

func (c Converter) Vid2mp3(vid string) (fpath string, err error) {
	// TODO: Invalid Vid Format Error (for example)
	fpath, err = c.Client.Execute(vid)
	return
}
