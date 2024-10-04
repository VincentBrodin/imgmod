package kernels

import (
	"image"
	"image/color"
	"math"

	"github.com/VincentBrodin/imgmod"
)

func BoxBlurKernel(kernelSize int) imgmod.Kernel {
	kernelValues := make([][]float64, kernelSize)
	for i := range kernelSize {
		kernelValues[i] = make([]float64, kernelSize)
		for j := range kernelSize {
			kernelValues[i][j] = 1.0
		}
	}

	return imgmod.Kernel{
		KernelValues: kernelValues,
		KernelFunction: func(x, y int, k imgmod.Kernel, img *image.RGBA) color.RGBA {
			var sumR, sumG, sumB, sumA float64
			pixels := 0
			halfSize := int(math.Floor(float64(kernelSize) / 2))

			for i := -halfSize; i <= halfSize; i++ {
				for j := -halfSize; j <= halfSize; j++ {
					nx, ny := x+i, y+j
					if nx < 0 || nx >= img.Bounds().Dx() || ny < 0 || ny >= img.Bounds().Dy() {
						continue
					}

					r, g, b, a := img.At(nx, ny).RGBA()
					kernelValue := k.KernelValues[j+halfSize][i+halfSize]
					sumR += float64(r>>8) * kernelValue
					sumG += float64(g>>8) * kernelValue
					sumB += float64(b>>8) * kernelValue
					sumA += float64(a>>8) * kernelValue
					pixels++
				}
			}

			return color.RGBA{
				R: uint8(uint32(sumR / float64(pixels))),
				G: uint8(uint32(sumG / float64(pixels))),
				B: uint8(uint32(sumB / float64(pixels))),
				A: uint8(uint32(sumA / float64(pixels))),
			}
		},
	}
}

func GaussianBlurKernel(kernelSize int, sigma float64) imgmod.Kernel {
	if kernelSize%2 == 0 {
		panic("kernelSize must be odd")
	}

	kernelValues := make([][]float64, kernelSize)
	sum := 0.0
	halfSize := kernelSize / 2

	// Calculate Gaussian kernel values
	for i := 0; i < kernelSize; i++ {
		kernelValues[i] = make([]float64, kernelSize)
		for j := 0; j < kernelSize; j++ {
			x := float64(i - halfSize)
			y := float64(j - halfSize)
			kernelValues[i][j] = (1 / (2 * math.Pi * sigma * sigma)) * math.Exp(-(x*x+y*y)/(2*sigma*sigma))
			sum += kernelValues[i][j]
		}
	}

	// Normalize the kernel values
	if sum == 0 {
		panic("sum of kernel values is zero")
	}
	for i := 0; i < kernelSize; i++ {
		for j := 0; j < kernelSize; j++ {
			kernelValues[i][j] /= sum
		}
	}

	return imgmod.Kernel{
		KernelValues: kernelValues,
		KernelFunction: func(x, y int, k imgmod.Kernel, img *image.RGBA) color.RGBA {
			var sumR, sumG, sumB, sumA float64
			halfSize := kernelSize / 2

			for i := -halfSize; i <= halfSize; i++ {
				for j := -halfSize; j <= halfSize; j++ {
					nx, ny := x+i, y+j
					if nx < 0 || nx >= img.Bounds().Dx() || ny < 0 || ny >= img.Bounds().Dy() {
						continue
					}

					r, g, b, a := img.At(nx, ny).RGBA()
					kernelValue := k.KernelValues[i+halfSize][j+halfSize]
					sumR += float64(r>>8) * kernelValue
					sumG += float64(g>>8) * kernelValue
					sumB += float64(b>>8) * kernelValue
					sumA += float64(a>>8) * kernelValue
				}
			}

			return color.RGBA{
				R: uint8(sumR),
				G: uint8(sumG),
				B: uint8(sumB),
				A: uint8(sumA),
			}
		},
	}
}

func LaplacianKernel() imgmod.Kernel {
	return imgmod.Kernel{
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

}
