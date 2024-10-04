package main

import (
	"fmt"
	"image"
	"image/color"
	"os"

	"github.com/VincentBrodin/imgmod"
)

func main() {
	// Load image from disk
	img, err := imgmod.LoadImage("FILE_PATH")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Apply a filter to the image

	// Creating a new laplacian kernel
	laplacian := imgmod.Kernel{
		// This is for a 3x3 kernel with these values
		// 0   1  0
		// 1 -4  1
		// 0  1  0
		KernelValues: [][]float64{{0, 1, 0}, {1, -4, 1}, {0, 1, 0}},
		// The kernel function gets applied to all pixels in the images and expects an color for that pixel back
		KernelFunction: func(x, y int, k imgmod.Kernel, img *image.RGBA) color.RGBA {
			var sum float64

			// Here we loop through the kernel so that the pixel (x,y) is in the center of the kernel
			for i := -1; i <= 1; i++ {
				for j := -1; j <= 1; j++ {
					nx, ny := x+i, y+j
					// Bounds check
					if nx < 0 || nx >= img.Bounds().Dx() || ny < 0 || ny >= img.Bounds().Dy() {
						continue
					}

					r, g, b, _ := img.At(nx, ny).RGBA()
					// Converting to black and white
					bw := (r + g + b) / 3
					// Multiply the pixel by the kernel value
					kernelValue := k.KernelValues[j+1][i+1]
					// Add it to the running sum for the pixel
					sum += float64(bw>>8) * kernelValue
				}
			}

			// Return the color for pixel (x,y)
			return color.RGBA{
				R: uint8(sum),
				G: uint8(sum),
				B: uint8(sum),
				A: uint8(255),
			}
		},
	}

	// Apply our new kernel to the image
	output := imgmod.ApplyKernel(laplacian, img)

	// Save the output to the disk
	err = imgmod.SaveImage(output, "NEW_FILE_PATH")
}
