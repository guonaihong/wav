package main

import (
	"flag"
	"fmt"
	"github.com/guonaihong/wav"
	"os"
)

func main() {

	sampleRate := flag.Int("ar", "", "Audio sample rate")
	file := flag.String("f", "", "Pcm audio file name")
	wavFile := flag.String("w", "", "The name of the wav file to be created")

	pcmFile, err := os.Open(*file)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	defer pcmFile.Close()

	pcmFi, err := pcmFile.Stat()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	wav := wavHead.New(1, *sampleRate, 16, pcmFi.Size())

	wavFd, err := os.Create(*wavFile)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	defer wavFd.Close()

	head, err := wav.Marshal()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	wavFd.Write(head)
	io.Copy(wavFd)
}
