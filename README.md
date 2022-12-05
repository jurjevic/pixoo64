# pixoo-64

Go library to control the Divoom Pixoo-64 

__This library is highly work-in-progress...__

### Example:
![Result of example](https://github.com/jurjevic/pixoo64/raw/main/example.png)
````go
package main

import (
	"github.com/jurjevic/pixoo64"
	"golang.org/x/image/draw"
	"image/png"
	"log"
	"net/http"
	"net/netip"
)

const address = "192.168.188.62"

func main() {
	addr, err := netip.ParseAddr(address)
	if err != nil {
		log.Fatal("unsupported address", address)
	}
	device, err := pixoo64.NewDevice(addr)
	if err != nil {
		log.Fatal("failed to create device for", address)
	}
	// test images
	picture, err := device.Display().CreatePicture()
	if err != nil {
		log.Fatal("failed to create picture")
	}
	addFrame(picture, "https://github.com/MariaLetta/free-gophers-pack/raw/master/characters/png/1.png")
	addFrame(picture, "https://github.com/MariaLetta/free-gophers-pack/raw/master/characters/png/2.png")
	addFrame(picture, "https://github.com/MariaLetta/free-gophers-pack/raw/master/characters/png/3.png")
	addFrame(picture, "https://github.com/MariaLetta/free-gophers-pack/raw/master/characters/png/4.png")
	addFrame(picture, "https://github.com/MariaLetta/free-gophers-pack/raw/master/characters/png/5.png")
	addFrame(picture, "https://github.com/MariaLetta/free-gophers-pack/raw/master/characters/png/6.png")
	err = picture.Update()
	if err != nil {
		log.Fatal("failed to update picture")
	}
}

func addFrame(picture *pixoo64.Picture, imageSrc string) {
	frame := picture.CreateFrame()
	frame.SetSpeed(1000)

	// Just a simple GET request to the image URL
	// We get back a *Response, and an error
	res, err := http.Get(imageSrc)
	if err != nil {
		log.Fatal("failed to load png", err)
	}
	defer res.Body.Close()

	img, err := png.Decode(res.Body)
	if err != nil {
		log.Fatal("failed to decode png", err)
	}

	draw.NearestNeighbor.Scale(frame.Rgba(), frame.Rgba().Bounds(), img, img.Bounds(), draw.Over, nil)
}

````

__Notice:__ Images provided by https://github.com/MariaLetta/free-gophers-pack
