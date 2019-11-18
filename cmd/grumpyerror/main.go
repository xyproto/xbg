package main

import (
	"time"

	"github.com/xyproto/xbg"
)

// This just demonstrates that SetWallpaper is thread safe
func main() {
	go func() {
		for {
			xbg.SetWallpaper("grumpy-cat.png", "zoom", true)
		}
	}()
	go func() {
		for {
			xbg.SetWallpaper("grumpy-cat.png", "zoom", true)
		}
	}()
	time.Sleep(10 * time.Second)
}
