package sprite

import (
	"ebby/errdef"
	"image"

	// image/png : 解码png图片
	_ "image/png"
	"os"

	"github.com/faiface/pixel"
)

func loadPicture(path string) (pixel.Picture, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return pixel.PictureDataFromImage(img), nil
}

// SheetFromPicture : 从图片文件读取
func SheetFromPicture(picPath string) pixel.Picture {
	spritesheet, err := loadPicture(picPath)
	errdef.CheckErr(err, "sprite/picture", errdef.PictureLoadFile)
	return spritesheet
}

// SeparateFrames : 根据sprite sheet均匀分割frames
func SeparateFrames(stepX int, stepY int, ss *pixel.Picture) []pixel.Rect {

	sX := float64(stepX)
	sY := float64(stepY)

	ssMinX := (*ss).Bounds().Min.X
	ssMaxX := (*ss).Bounds().Max.X
	ssMinY := (*ss).Bounds().Min.Y
	ssMaxY := (*ss).Bounds().Max.Y

	var frames []pixel.Rect

	for y := ssMaxY - sY; y > ssMinY; y -= sY {
		for x := ssMinX; x < ssMaxX; x += sX {
			frames = append(frames, pixel.R(x, y, x+sX, y+sY))
		}
	}

	return frames

}
