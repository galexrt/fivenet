package images

import (
	"bytes"
	"image/png"
	"io"
)

func ResizePNG(input io.Reader, height int, width int) ([]byte, error) {
	// Decode the image (from PNG to image.Image)
	src, err := png.Decode(input)
	if err != nil {
		return nil, err
	}

	dst := resizeImageIfNecessary(src, height, width)
	if dst == nil {
		return nil, nil
	}

	// Encode to output
	output := new(bytes.Buffer)
	if err := png.Encode(output, dst); err != nil {
		return nil, err
	}

	return output.Bytes(), nil
}