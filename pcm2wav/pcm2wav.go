package main

import (
	"flag"
	"fmt"
	"github.com/guonaihong/wav"
	"io"
	"os"
)

func main() {

	sampleRate := flag.Int("ar", 0, "Audio sample rate")
	file := flag.String("f", "", "Pcm audio file name")
	wavFile := flag.String("w", "", "The name of the wav file to be created")

	flag.Parse()

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

	wavHead := wav.New(uint16(1), uint32(*sampleRate), uint16(16), uint32(pcmFi.Size()))

	wavFd, err := os.Create(*wavFile)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	defer wavFd.Close()

	head, err := wavHead.Marshal()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	wavFd.Write(head)
	io.Copy(wavFd, pcmFile)
}
