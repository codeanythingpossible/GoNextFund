package services

import (
	"errors"
	"os"
)

type Settings struct {
	StorageFolder string
	LogsFolder    string
}

func (s *Settings) StoreInFile(path string) error {
	stat, err := os.Stat(path)
	if err != nil {
		return err
	}
	if stat.IsDir() {
		return errors.New("cannot store in a directory")
	}

	return errors.New("cannot store in a file")
}
