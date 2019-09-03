package font

import (
	"fmt"
	"io/ioutil"
	"os"

	"ebby/common/logger"
	"ebby/errdef"

	"github.com/faiface/pixel/text"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
)

func loadTTF(path string, size float64) (font.Face, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	font, err := truetype.Parse(bytes)
	if err != nil {
		return nil, err
	}

	return truetype.NewFace(font, &truetype.Options{
		Size:              size,
		GlyphCacheEntries: 1,
	}), nil
}

func errLog(err error) {
	if err != nil {
		log := logger.New()
		log.Error(fmt.Sprintf("[common/font]<%s> %v", errdef.FontLoadFile.Str, err))
		panic(err)
	}
}

// GetAtlas : 获取字体 Atlas
// TODO: 字体部分还需要详细调整
func GetAtlas(path string, size float64) *text.Atlas {
	face, err := loadTTF(path, size)
	errLog(err)

	return text.NewAtlas(face, text.ASCII)
}

// GetDebugAtlas : 获取debug用字体 Atlas: basicfont.Face7x13
func GetDebugAtlas() *text.Atlas {
	return text.NewAtlas(basicfont.Face7x13, text.ASCII)
}
