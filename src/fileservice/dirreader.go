package fileservice

import (
	"errors"
	"io/ioutil"
	"os"
)

func RenderTop(dir string) (<-chan FileEntity, error) {

	if s, e := os.Stat(dir); e == nil {
		if !s.IsDir() || !validate(dir) {
			return nil, errors.New("invalid path")
		}
	} else {
		return nil, e
	}

	ch := make(chan FileEntity)
	go render(dir, ch)

	return ch, nil
}

func render(path string, c chan FileEntity) {
	defer close(c)
	if found, e := ioutil.ReadDir(path); e == nil {
		for _, f := range found {
			c <- FileEntity{
				Name:  f.Name(),
				IsDir: f.IsDir(),
			}
		}
	} else {
		c <- FileEntity{
			IsError: true,
		}
	}
}

func validate(path string) bool {
	// TODO: check for travessal based on root path
	return true
}
