package font

import (
	"io/ioutil"
	"os"

	"ebby/errdef"

	"github.com/faiface/pixel"
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

// Atlas : 获取字体 Atlas
// TODO: 字体部分还需要详细调整
func Atlas(path string, size float64) *text.Atlas {
	face, err := loadTTF(path, size)
	errdef.CheckErr(err, "common/font", errdef.FontLoadFile)

	return text.NewAtlas(face, text.ASCII)
}

// DebugAtlas : 获取debug用字体 Atlas: basicfont.Face7x13
func DebugAtlas() *text.Atlas {
	return text.NewAtlas(basicfont.Face7x13, text.ASCII)
}

// GetDebugLoggerDisplayCallBack : DebugLogger触发窗口变化后的回调函数
func GetDebugLoggerDisplayCallBack(logger *text.Text) func(pixel.Vec) {
	dl := logger
	return func(winVec pixel.Vec) {
		dl.Orig = pixel.V(4, winVec.Y-dl.LineHeight)
	}
}
