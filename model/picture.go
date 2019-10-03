package model

import (
	"github.com/conest/ebby/system"
	"image"
	// image/png : 解码png图片
	_ "image/png"
	"os"

	"github.com/faiface/pixel"
)

func loadPicture(path string) (*pixel.PictureData, error) {
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

// PictureDataFromFile : 从图片文件读取 PictureData
func PictureDataFromFile(picPath string) *pixel.PictureData {
	pictureData, err := loadPicture(picPath)
	system.CheckErr(err, "model/picture", system.ErrorTable["PictureLoadFile"])
	return pictureData
}

// PictureFromFile : 从图片文件读取 Picture
func PictureFromFile(picPath string) *pixel.Picture {
	pictureData := PictureDataFromFile(picPath)
	spritesheet := pixel.Picture(pictureData)
	return &spritesheet
}

// SeparateFrames : 根据Rect均匀分割frames
func SeparateFrames(stepX int, stepY int, ss pixel.Rect) []pixel.Rect {

	sX := float64(stepX)
	sY := float64(stepY)
	ssMinX := ss.Min.X
	ssMaxX := ss.Max.X
	ssMinY := ss.Min.Y
	ssMaxY := ss.Max.Y

	var frames []pixel.Rect
	for y := ssMaxY - sY; y >= ssMinY; y -= sY {
		for x := ssMinX; x < ssMaxX; x += sX {
			frames = append(frames, pixel.R(x, y, x+sX, y+sY))
		}
	}
	return frames
}

// CombinePictureData : 将多个PictureData拼合为一个，用于合并Sprite sheet以便在draw to Batch
func CombinePictureData(datas []*pixel.PictureData) *pixel.Picture {
	if len(datas) == 0 {
		p := pixel.Picture(&pixel.PictureData{})
		return &p
	}
	// 获取最宽图片边长以及总高度
	var width float64
	var height float64
	for _, pd := range datas {
		if pd.Bounds().W() > width {
			width = pd.Bounds().W()
		}
		height += pd.Bounds().H()
	}

	rect := pixel.R(0, 0, width, height)
	np := pixel.MakePictureData(rect)

	stride := int(width)
	var offset int
	for _, pd := range datas {
		// space := make([]color.RGBA, int(width-pd.Bounds().W()))
		for y := 0; y < int(pd.Bounds().H()); y++ {
			ps := pd.Pix[y*pd.Stride : y*pd.Stride+pd.Stride]
			np.Pix = append(np.Pix[:offset], ps...)
			offset += stride
		}

	}
	picture := pixel.Picture(np)
	return &picture
}

// ShiftFrames : 根据多个将要或以及合并的PictureData，返回新的Frames以便draw to Batch
func ShiftFrames(datas []*pixel.PictureData, framesList [][]pixel.Rect) [][]pixel.Rect {
	if len(datas) != len(framesList) || len(datas) < 2 {
		return framesList
	}

	fl := make([][]pixel.Rect, len(framesList))
	fl[0] = make([]pixel.Rect, len(framesList[0]))
	copy(fl[0], framesList[0])

	for i := 1; i < len(datas); i++ {
		fl[i] = make([]pixel.Rect, len(framesList[i]))
		copy(fl[i], framesList[i])

		shift := datas[i-1].Bounds().H()
		for u := 0; u < len(fl[i]); u++ {
			fl[i][u].Min.Y += shift
			fl[i][u].Max.Y += shift
		}
	}

	return fl
}
