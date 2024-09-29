package services

import (
	"encoding/binary"
	"errors"
	"log"
	"os"
	"path/filepath"
)

type Settings struct {
	StorageFolder string
	LogsFolder    string
}

func DefaultSettings() *Settings {
	return &Settings{
		LogsFolder:    ".",
		StorageFolder: ".",
	}
}

func (s *Settings) StoreInFile(path string) error {
	stat, err := os.Stat(path)
	if err != nil {
		return err
	}
	if stat.IsDir() {
		return errors.New("cannot store in a directory")
	}

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)

	err = binary.Write(file, binary.LittleEndian, s.StorageFolder)
	if err != nil {
		return err
	}

	err = binary.Write(file, binary.LittleEndian, s.LogsFolder)
	if err != nil {
		return err
	}

	return nil
}

func LoadFromFile(path string) (*Settings, error) {
	stat, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if stat.IsDir() {
		return nil, errors.New("cannot store in a directory")
	}

	file, err := os.OpenFile(path, os.O_RDONLY, 0600)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)

	var storageFolder string
	var logsFolder string

	err = binary.Read(file, binary.LittleEndian, &storageFolder)
	if err != nil {
		return nil, err
	}
	err = binary.Read(file, binary.LittleEndian, &logsFolder)
	if err != nil {
		return nil, err
	}

	return &Settings{StorageFolder: storageFolder, LogsFolder: logsFolder}, nil
}

func CheckIfSettingsFileExists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func GetDefaultSettingsFilepath() string {
	exePath, err := os.Executable()
	if err != nil {
		panic(err)
	}

	exeDir := filepath.Dir(exePath)
	return filepath.Join(exeDir, "settings.bin")
}
