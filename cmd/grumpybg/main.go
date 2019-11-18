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
	// This check is redundant, but it's nice to check it before checking if the WM is ready
	if !exists(imageFilename) {
		return fmt.Errorf("no such file: %s", imageFilename)
	}
	wm := xbg.X11{}
	if !(wm.Running() && wm.ExecutablesExists()) {
		return errors.New("found no working method for setting the desktop wallpaper, maybe X11 is not installed or DISPLAY not set")
	}
	if verbose {
		fmt.Printf("Using the %s backend.\n", wm.Name())
	}
	wm.SetVerbose(verbose)
	wm.SetMode(mode)
	if err := wm.SetWallpaper(imageFilename); err != nil {
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
	return nil
}

func main() {
	//SetWallpaper("grumpy-cat.png", "center", true)
	SetWallpaper("grumpy-cat.png", "zoom", true)
	//SetWallpaper("grumpy-cat.png", "scale", true)
}
