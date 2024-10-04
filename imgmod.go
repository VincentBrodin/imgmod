package imgmod

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

const (
	PNG  = ".png"
	JPG  = ".jpg"
	JPEG = ".jpeg"
)

// Kernel represents a convolution kernel used in image processing.
type Kernel struct {
	// KernelValues holds the matrix values of the kernel.
	KernelValues [][]float64

	// KernelFunction applies the kernel to a specific pixel in the image.
	KernelFunction func(x, y int, k Kernel, img *image.RGBA) color.RGBA
}

func ApplyKernel(kernel Kernel, img image.Image) image.Image {
	input := ImageToRGBA(img)
	output := ImageToRGBA(img)
	for x := range input.Bounds().Dx() {
		for y := range input.Bounds().Dy() {
			color := kernel.KernelFunction(x, y, kernel, input)
			output.SetRGBA(x, y, color)
		}
	}
	return output
}

// Loads any image from disk into ram as a image.Image
func LoadImage(imagePath string) (image.Image, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	return img, err
}

// Converts any image into a RGBA image
func ImageToRGBA(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	newImg := image.NewRGBA(image.Rect(0, 0, bounds.Dx(), bounds.Dy()))
	draw.Draw(newImg, newImg.Bounds(), img, bounds.Min, draw.Src)
	return newImg
}

// Encodes and saves the given image to the given path
func SaveImage(img image.Image, imagePath string) error {
	// Get the extension type and checks if it is a valid extension
	extension := getExtension(imagePath)
	if !validExtension(extension) {
		return fmt.Errorf("Can't save %s images\n", extension)
	}

	// Create or truncate file
	file, err := os.Create(imagePath)
	if err != nil {
		return err
	}

	// Encode image
	switch extension {
	case PNG:
		err = png.Encode(file, img)
	case JPEG, JPG:
		err = jpeg.Encode(file, img, nil)
	}

	return err
}

// Checks that the given extension can be encoded
func validExtension(extension string) bool {
	extensions := []string{PNG, JPG, JPEG}

	for _, ext := range extensions {
		if ext == extension {
			return true
		}
	}

	return false
}

// Loops from the end of the string to the front and returns everything past the last "." in the string
func getExtension(imagePath string) string {
	index := -1
	for i := len(imagePath) - 1; i >= 0; i-- {
		if imagePath[i] == '.' {
			index = i
			break
		}
	}

	return imagePath[index:]
}
