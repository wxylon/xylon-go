package mp

import (
	"fmt"
	"time"
)

type WAVPlayer struct {
	stat     int
	progress int
	signal   chan int
}

func (wav *WAVPlayer) Play(source string) {
	fmt.Println("Playing wav music", source)

	wav.progress = 0

	for wav.progress < 100 {
		time.Sleep(100 * time.Millisecond) // 假装正在播放
		fmt.Print(".")
		wav.progress += 10
	}
	fmt.Println()
}
