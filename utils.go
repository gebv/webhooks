package main

import (
	"encoding/json"
	"errors"
	// "fmt"
	"io"
	// "io/ioutil"
	// "log"
	"os"
	"path/filepath"
)

var (
	ErrNotFound     = errors.New("not_found")
	ErrNotSupported = errors.New("not_supported")
	ErrNotAllowed   = errors.New("not_allowed")
	ErrNotValid     = errors.New("not_valid")
)

var Cfg = &Config{}

var CfgBuildStamp string = ""
var CfgVersion string = ""
var CfgGitHash string = ""

type Config struct {
	Hooks HooksConfig
}

type HooksConfig []HookConfig

func (h HooksConfig) ByToken(token string) *HookConfig {
	for _, config := range h {
		if config.Token == token {
			return &config
		}
	}

	return nil
}

type HookConfig struct {
	Token        string
	ProviderName string

	ShellCommand string
	IsEnabled    bool
	Description  string

	Check struct {
		UserName       string
		BranchName     string
		RepositoryName string
		MessageTagName string
	}
}

// FromJson
func FromJson(obj interface{}, data interface{}) error {
	switch data.(type) {
	case io.Reader:
		decoder := json.NewDecoder(data.(io.Reader))
		return decoder.Decode(obj)
	case []byte:
		return json.Unmarshal(data.([]byte), obj)
	}

	return ErrNotSupported
}

func FindConfigFile(fileName string) string {
	if len(fileName) == 0 {
		panic("Empty file name")
	}

	if _, err := os.Stat("./config/" + fileName); err == nil {
		fileName, _ = filepath.Abs("./config/" + fileName)
	} else if _, err := os.Stat("../config/" + fileName); err == nil {
		fileName, _ = filepath.Abs("../config/" + fileName)
	} else if _, err := os.Stat(fileName); err == nil {
		fileName, _ = filepath.Abs(fileName)
	} else {
		panic("Not found " + fileName)
	}

	return fileName
}

func LoadConfig(fileName string) {
	fileName = FindConfigFile(fileName)

	file, err := os.Open(fileName)

	if err != nil {
		panic("Error opening config file=" + fileName + ", err=" + err.Error())
	}

	decoder := json.NewDecoder(file)

	err = decoder.Decode(Cfg)

	if err != nil {
		panic("Error decoding config file=" + fileName + ", err=" + err.Error())
	}
}
