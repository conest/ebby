package model

import (
	"ebby/system"
	"io/ioutil"
	"os"
)

// LoadGLSL : 读取glsl
func LoadGLSL(path string) string {
	glsl, err := loadFile(path)
	system.CheckErr(err, "GLSL", system.ErrorTable["GLSLLoad"])

	return glsl
}

func loadFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
