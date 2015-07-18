package mp

import (
	"fmt"
	"time"
)

type MP3Player struct {
	stat     int
	progress int
}

func (mp3 *MP3Player) Play(source string) {
	fmt.Println("Playing MP3 music", source)

	mp3.progress = 0

	for mp3.progress < 100 {
		time.Sleep(100 * time.Millisecond) // 假装正在播放
		fmt.Print(".")
		mp3.progress += 10
	}

	fmt.Println("\nFinished playing", source)
}
