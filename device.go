package pixoo64

import (
	"net/http"
	"net/netip"
	"strconv"
)

type Device struct {
	client  *Client
	channel *Channel
	display *Display
}

func (d *Device) Display() *Display {
	return d.display
}

func (d *Device) Channel() *Channel {
	return d.channel
}

// NewDevice creates a new device with the given address.
func NewDevice(addr netip.Addr) (*Device, error) {
	client := &Client{
		addr:       addr,
		httpClient: http.DefaultClient,
	}
	display, err := NewDisplay(client)
	if err != nil {
		return nil, err
	}
	return &Device{
		client:  client,
		channel: &Channel{client: client},
		display: display,
	}, nil
}

func (d *Device) Alert() error {
	return d.PlayBuzzer(500, 500, 3000)
}

func (d *Device) PlayBuzzer(activeTimeInCycle int, offTimeInCycle int, playTotalTime int) error {
	var jsonData = []byte(`{
			"Command":"Device/PlayBuzzer",
			"ActiveTimeInCycle":` + strconv.FormatInt(int64(activeTimeInCycle), 10) + `,
			"OffTimeInCycle":` + strconv.FormatInt(int64(offTimeInCycle), 10) + `,
			"PlayTotalTime":` + strconv.FormatInt(int64(playTotalTime), 10) + `
		}`)
	_, err := d.client.Post(jsonData)
	if err != nil {
		return err
	}
	return nil
}

type Status int

const (
	Stop  Status = 0
	Start Status = 1
)

func (d *Device) Countdown(minute int, second int, status Status) error {
	var jsonData = []byte(`      {
    "Command":"Tools/SetTimer",
    "Minute": ` + strconv.FormatInt(int64(minute), 10) + `,
    "Second": ` + strconv.FormatInt(int64(second), 10) + `,
    "Status": ` + strconv.FormatInt(int64(status), 10) + `
    }`)
	_, err := d.client.Post(jsonData)
	if err != nil {
		return err
	}
	return nil
}

func (d *Device) Stopwatch(status Status) error {
	var jsonData = []byte(`      {
    "Command":"Tools/SetStopWatch",
    "Status": ` + strconv.FormatInt(int64(status), 10) + `
    }`)
	_, err := d.client.Post(jsonData)
	if err != nil {
		return err
	}
	return nil
}

func (d *Device) Scoreboard(blueScore int, redScore int) error {
	var jsonData = []byte(`      {
    "Command":"Tools/SetScoreBoard",
    "BlueScore": ` + strconv.FormatInt(int64(blueScore), 10) + `,
    "RedScore": ` + strconv.FormatInt(int64(redScore), 10) + `
    }`)
	_, err := d.client.Post(jsonData)
	if err != nil {
		return err
	}
	return nil
}

func (d *Device) Noise(status Status) error {
	var jsonData = []byte(`      {
    "Command":"Tools/SetNoiseStatus",
    "NoiseStatus": ` + strconv.FormatInt(int64(status), 10) + `
    }`)
	_, err := d.client.Post(jsonData)
	if err != nil {
		return err
	}
	return nil
}
