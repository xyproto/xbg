package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/xyproto/xbg"
)

// exists checks if the given path exists
func exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// SetWallpaper will set the given image filename as the wallpaper,
// regardless of which display server, window manager or desktop environment is in use.
func SetWallpaper(imageFilename, mode string, verbose bool) error {
	if !exists(imageFilename) {
		return fmt.Errorf("no such file: %s", imageFilename)
	}
	var lastErr error
	wm := xbg.X11{}
	if wm.Running() && wm.ExecutablesExists() {
		if verbose {
			fmt.Printf("Using the %s backend.\n", wm.Name())
		}
		wm.SetVerbose(verbose)
		wm.SetMode(mode)
		if err := wm.SetWallpaper(imageFilename); err != nil {
			lastErr = err
			switch wm.Name() {
			case "Weston":
				// If the current windowmanager is Weston, no method is currently available
				return err
			default:
				if verbose {
					fmt.Fprintf(os.Stderr, "failed: %v\n", err)
				}
				// If the wallpaper mode is wrong, don't try the next backend, but return the error
				if strings.Contains(err.Error(), "invalid desktop wallpaper mode") {
					return err
				}
				// This did not work out
				return errors.New("could not use the X11 backend")
			}
		} else {
			return nil
		}
	}
	if lastErr != nil {
		return fmt.Errorf("found no working method for setting the desktop wallpaper:\n%v", lastErr)
	}
	return errors.New("found no working method for setting the desktop wallpaper")

}

func main() {
	SetWallpaper("grumpy-cat.png", "zoom", true)
}
