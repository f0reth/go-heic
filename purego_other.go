//go:build !unix && !darwin && !windows

package heic

import (
	"fmt"
	"image"
	"io"
)

var (
	dynamic    = false
	dynamicErr = fmt.Errorf("heic: platform not supported")
)

func decode(r io.Reader, configOnly bool) (image.Image, image.Config, error) {
	return nil, image.Config{}, dynamicErr
}

func loadLibrary() (uintptr, error) {
	return 0, dynamicErr
}
