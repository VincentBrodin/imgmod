package main

import (
	"fmt"
	"github.com/VincentBrodin/imgmod"
	"github.com/VincentBrodin/imgmod/kernels"
)

func main() {
	// Load image from disk
	img, err := imgmod.LoadImage("FILE_PATH")
	if err != nil {
		fmt.Println(err.Error())
	}

	// Apply a filter to the image

	// Getting a premade box blur kernel from the imgmod/kernels package
	boxBlurKernel := kernels.BoxBlurKernel(3)
	output := imgmod.ApplyKernel(boxBlurKernel, img)

	// Save the output to the disk
	err = imgmod.SaveImage(output, "NEW_FILE_PATH")
}
