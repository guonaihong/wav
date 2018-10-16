参考资料
```
https://stackoverflow.com/questions/21131595/convert-from-pcm-to-wav-is-it-possible
https://blog.csdn.net/u010011236/article/details/53026127
```

安装示例
```console
env GOPATH=`pwd` go get github.com/guonaihong/wav/pcm2wav

./pcm2wav -ar 16000 -f ./good.pcm -w good.wav
```

##### `wav.go`
```
func New(numChannels uint16, sampleRate uint32, bitsPerSample uint16, wavLen uint32) *WavHead

func (wh *WavHead) Marshal() ([]byte, error)

```

##### 示例内容
```

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
```
