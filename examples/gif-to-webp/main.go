package main

import (
	"bytes"
	"image/gif"
	"io/ioutil"
	"log"
	"os"

	"github.com/snowmerak/webpanimation"
)

func main() {
	var buf bytes.Buffer
	gifFile, err := os.Open("animation.gif")
	if err != nil {
		log.Fatal(err)
	}
	gif, err := gif.DecodeAll(gifFile)
	if err != nil {
		log.Fatal(err)
	}

	webpanim := webpanimation.NewWebpAnimation(gif.Config.Width, gif.Config.Height, gif.LoopCount)
	webpanim.WebPAnimEncoderOptions.SetKmin(9)
	webpanim.WebPAnimEncoderOptions.SetKmax(17)
	defer webpanim.ReleaseMemory() // dont forget call this or you will have memory leaks
	webpConfig := webpanimation.NewWebpConfig()
	webpConfig.SetLossless(1)

	timeline := 0

	for i, img := range gif.Image {

		err = webpanim.AddFrame(img, timeline, webpConfig)
		if err != nil {
			log.Fatal(err)
		}
		timeline += gif.Delay[i] * 10
	}
	err = webpanim.AddFrame(nil, timeline, webpConfig)
	if err != nil {
		log.Fatal(err)
	}

	err = webpanim.Encode(&buf) // encode animation and write result bytes in buffer
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile("animation.webp", buf.Bytes(), 0777)
}
