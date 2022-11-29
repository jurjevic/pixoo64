package pixoo64

import "strconv"

type Channel struct {
	client *Client
}

func (c *Channel) ActivateVisualizer(eqPosition int) error {
	var jsonData = []byte(`      {
    "Command":"Channel/SetEqPosition",
    "EqPosition": ` + strconv.FormatInt(int64(eqPosition), 10) + `
    }`)
	_, err := c.client.Post(jsonData)
	if err != nil {
		return err
	}
	return nil
}

type CloudIndex int

const (
	RecommendCloudIndex       CloudIndex = 0
	FavouriteCloudIndex       CloudIndex = 1
	SubscribeArtistCloudIndex CloudIndex = 2
	AlbumCloudIndex           CloudIndex = 3
)

func (c *Channel) ActivateCloudIndex(index CloudIndex) error {
	var jsonData = []byte(`      {
    "Command":"Channel/CloudIndex",
    "Index": ` + strconv.FormatInt(int64(index), 10) + `
    }`)
	_, err := c.client.Post(jsonData)
	if err != nil {
		return err
	}
	return nil
}

func (c *Channel) ActivateCustomPage(index int) error {
	var jsonData = []byte(`      {
    "Command":"Channel/SetCustomPageIndex",
    "CustomPageIndex": ` + strconv.FormatInt(int64(index), 10) + `
    }`)
	_, err := c.client.Post(jsonData)
	if err != nil {
		return err
	}
	return nil
}

type Index int

const (
	FacesIndex        Index = 0
	CloudChannelIndex Index = 1
	VisualizerIndex   Index = 2
	CustomIndex       Index = 3
	BlackScreenIndex  Index = 4
)

func (c *Channel) ActivateSetIndex(index Index) error {
	var jsonData = []byte(`      {
    "Command":"Channel/SetIndex",
    "SelectIndex": ` + strconv.FormatInt(int64(index), 10) + `
    }`)
	_, err := c.client.Post(jsonData)
	if err != nil {
		return err
	}
	return nil
}

func (c *Channel) SetClock(id int) error {
	var jsonData = []byte(`      {
    "Command":"Channel/SetClockSelectId",
    "ClockId": ` + strconv.FormatInt(int64(id), 10) + `
    }`)
	_, err := c.client.Post(jsonData)
	if err != nil {
		return err
	}
	return nil
}
