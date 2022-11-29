package pixoo64

import (
	"encoding/base64"
	"image"
	"image/color"
	"strconv"
)

type Frame struct {
	client *Client
	id     int
	rgba   *image.RGBA
	speed  int
}

func (f *Frame) Id() int {
	return f.id
}

func (f *Frame) Rgba() *image.RGBA {
	return f.rgba
}

func (f *Frame) Speed() int {
	return f.speed
}

func (f *Frame) SetSpeed(speed int) {
	f.speed = speed
}

func (f *Frame) Fill(color color.Color) {
	for y := 0; y < f.rgba.Rect.Size().Y; y++ {
		for x := 0; x < f.rgba.Rect.Size().X; x++ {
			f.rgba.Set(x, y, color)
		}
	}
}

func (f *Frame) Update(pictureId int, frameCount int) error {
	count := 0
	raw := make([]byte, f.rgba.Rect.Size().X*f.rgba.Rect.Size().Y*3)
	for y := 0; y < f.rgba.Rect.Size().Y; y++ {
		for x := 0; x < f.rgba.Rect.Size().X; x++ {
			raw[count*3] = f.rgba.RGBAAt(x, y).R
			raw[count*3+1] = f.rgba.RGBAAt(x, y).G
			raw[count*3+2] = f.rgba.RGBAAt(x, y).B
			count++
		}
	}

	base64 := base64.StdEncoding.EncodeToString(raw)

	var jsonData = []byte(`{
	    "Command": "Draw/SendHttpGif",
	    "PicNum": ` + strconv.FormatInt(int64(frameCount), 10) + `,
	    "PicWidth": ` + strconv.FormatInt(int64(Width), 10) + `,
	    "PicOffset": ` + strconv.FormatInt(int64(f.id), 10) + `,
	    "PicID": ` + strconv.FormatInt(int64(pictureId), 10) + `,
	    "PicSpeed": ` + strconv.FormatInt(int64(f.speed), 10) + `,
	    "PicData": "` + base64 + `"
	}`)
	_, err := f.client.Post(jsonData)
	return err
}
