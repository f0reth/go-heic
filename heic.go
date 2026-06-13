// Package heic implements an HEIC image decoder based on the libheif/libde265
// shared library used via purego (CGo-free).
package heic

import (
	"errors"
	"image"
	"io"
)

// Errors .
var (
	ErrMemRead  = errors.New("heic: mem read failed")
	ErrMemWrite = errors.New("heic: mem write failed")
	ErrDecode   = errors.New("heic: decode failed")
)

// Decode reads a HEIC image from r and returns it as an image.Image.
func Decode(r io.Reader) (image.Image, error) {
	if !dynamic {
		return nil, dynamicErr
	}

	img, _, err := decode(r, false)
	if err != nil {
		return nil, err
	}

	return img, nil
}

// DecodeConfig returns the color model and dimensions of a HEIC image without decoding the entire image.
func DecodeConfig(r io.Reader) (image.Config, error) {
	if !dynamic {
		return image.Config{}, dynamicErr
	}

	_, cfg, err := decode(r, true)
	if err != nil {
		return image.Config{}, err
	}

	return cfg, nil
}

// Dynamic returns error (if there was any) during opening dynamic/shared library.
func Dynamic() error {
	return dynamicErr
}

// Init is kept for backward compatibility and does nothing.
// The dynamic/shared library is loaded during package initialization;
// use Dynamic to check whether loading succeeded.
func Init() {}

const (
	heifMaxHeaderSize = 262144

	heifColorspaceUndefined  = 99
	heifColorspaceYCbCr      = 0
	heifColorspaceRGB        = 1
	heifColorspaceMonochrome = 2

	heifChannelY           = 0
	heifChannelCb          = 1
	heifChannelCr          = 2
	heifChannelR           = 3
	heifChannelG           = 4
	heifChannelB           = 5
	heifChannelAlpha       = 6
	heifChannelInterleaved = 10

	heifChromaUndefined       = 99
	heifChromaMonochrome      = 0
	heifChroma420             = 1
	heifChroma422             = 2
	heifChroma444             = 3
	heifChromaInterleavedRGBA = 11

	heifFiletypeYesSupported = 1
)

func yCbCrSize(r image.Rectangle, subsampleRatio image.YCbCrSubsampleRatio) (w, h, cw, ch int) {
	w, h = r.Dx(), r.Dy()

	switch subsampleRatio {
	case image.YCbCrSubsampleRatio422:
		cw = (r.Max.X+1)/2 - r.Min.X/2
		ch = h
	case image.YCbCrSubsampleRatio420:
		cw = (r.Max.X+1)/2 - r.Min.X/2
		ch = (r.Max.Y+1)/2 - r.Min.Y/2
	default:
		cw = w
		ch = h
	}

	return
}

func init() {
	image.RegisterFormat("heic", "????ftypheic", Decode, DecodeConfig)
}
