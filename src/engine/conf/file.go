package conf

import (
	"io/ioutil"
)

func Read(filename string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(ResDir(filename))
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
