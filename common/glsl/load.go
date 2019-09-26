package glsl

import (
	"ebby/errdef"
	"io/ioutil"
	"os"
)

// Load : 读取glsl
func Load(path string) string {
	glsl, err := loadFile(path)
	errdef.CheckErr(err, "GLSL", errdef.GLSLLoad)

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
