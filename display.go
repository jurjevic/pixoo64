package pixoo64

import (
	"encoding/json"
	"strconv"
)

const (
	Width  = 64
	Height = 64
)

// Display manages the graphics on the pixoo64 device.
type Display struct {
	client   *Client
	pictures []*Picture
}

// NewDisplay create a new Display.
func NewDisplay(client *Client) (*Display, error) {
	d := &Display{
		client: client,
	}
	err := d.Reset()
	if err != nil {
		return nil, err
	}
	return d, nil
}

type GetHttpGifId struct {
	ErrorCode int `json:"error_code"`
	PicId     int `json:"PicId"`
}

func (d *Display) GetHttpGifId() (int, error) {
	var jsonData = []byte(`      {
    "Command":"Draw/GetHttpGifId"
    }`)
	body, err := d.client.Post(jsonData)
	if err != nil {
		return 0, err
	}
	result := GetHttpGifId{}
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		return 0, err
	}
	return result.PicId, nil
}

func (d *Display) CreatePicture() (*Picture, error) {
	nextId, err := d.GetHttpGifId()
	if err != nil {
		return nil, err
	}
	picture := &Picture{
		client: d.client,
		id:     nextId,
		frames: nil,
	}
	d.pictures = append(d.pictures, picture)
	return picture, nil
}

func (d *Display) Reset() error {
	err := d.ResetText()
	if err != nil {
		return err
	}
	for _, p := range d.pictures {
		err = p.Reset()
		if err != nil {
			return err
		}
	}
	return err
}

func (d *Display) Update() error {
	for _, p := range d.pictures {
		err := p.Update()
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Display) ResetText() error {
	var jsonData = []byte(`      {
    "Command":"Draw/ClearHttpText"
    }`)
	_, err := d.client.Post(jsonData)
	return err
}

func (d *Display) SetBrightness(brightness int) error {
	var jsonData = []byte(`      {
    "Command":"Channel/SetBrightness",
    "Brightness": ` + strconv.FormatInt(int64(brightness), 10) + `
    }`)
	_, err := d.client.Post(jsonData)
	if err != nil {
		return err
	}
	return nil
}

type Switch int

const (
	Off Switch = 0
	On  Switch = 1
)

func (d *Display) Switch(onoff Switch) error {
	var jsonData = []byte(`      {
    "Command":"Channel/OnOffScreen",
    "OnOff": ` + strconv.FormatInt(int64(onoff), 10) + `
    }`)
	_, err := d.client.Post(jsonData)
	if err != nil {
		return err
	}
	return nil
}
