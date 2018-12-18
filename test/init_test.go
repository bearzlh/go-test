package test

import (
	"io/ioutil"
	"mq/service"
	"os"
	"path/filepath"
	"testing"
)

const envFile = "../.env"

func TestMain(m *testing.M) {
	env := "dev"

	file, _ := filepath.Abs(envFile)

	fileState, _ := os.Stat(envFile)

	if !fileState.IsDir() {
		content, _ := ioutil.ReadFile(envFile)
		env = string(content)
	}

	dir := filepath.Dir(file)

	service.GetConfig(env, dir)
	m.Run()
}