package glsl

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Load : 读取glsl
func Load(path string) string {
	glsl, err := loadFile(path)
	if err != nil {
		panic(fmt.Errorf("[GLSL] Load error: %s", err))
	}
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
