package pixoo64

import "image"

type Picture struct {
	client *Client
	id     int
	frames []*Frame
}

func (p *Picture) Reset() error {
	var jsonData = []byte(`{"Command":"Draw/ResetHttpGifId"}`)
	_, err := p.client.Post(jsonData)
	p.frames = []*Frame{}
	return err
}

func (p *Picture) CreateFrame() *Frame {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{Width, Height}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})
	frame := &Frame{
		client: p.client,
		id:     len(p.frames),
		rgba:   img,
		speed:  100,
	}
	p.frames = append(p.frames, frame)
	return frame
}

func (p *Picture) Update() error {
	for _, frame := range p.frames {
		err := frame.Update(p.id, len(p.frames))
		if err != nil {
			return err
		}
	}
	return nil
}
