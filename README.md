# imgmod - Image Processing for Golang

Welcome to **imgmod**, an easy-to-use image processing package for Golang! **imgmod** gives you the tools you need to manipulate and process images seamlessly.
And it's built entirely on top of Go's standard library `image` package, ensuring full compatibility with your existing code and packages.

## Features
- 📦 **Premade kernels for quick filtering** like Box Blur, Gaussian Blur, and Laplacian.
- 🔧 **Custom kernel support** to create your own filters.
- ⚙️ **Flexible and extensible**, making it easy to integrate into your projects.

## Installation
Getting started is super simple! Just run the following command in your Go project:

```bash
go get github.com/VincentBrodin/imgmod
```

And you’re all set to start processing images.

## Demo

Here’s a quick example to show how you can load an image, apply a filter, and save the modified image.

```go
package main

import (
    "fmt"
    "github.com/VincentBrodin/imgmod"
    "github.com/VincentBrodin/imgmod/kernels"
)

func main() {
    // Load an image from disk
    img, err := imgmod.LoadImage("FILE_PATH")
    if err != nil {
        fmt.Println(err.Error())
    }

    // Apply a premade box blur filter to the image
    // The kernels.BoxBlurKernel() function takes in an int for the kernel size
    // So in this case it is a 3x3 kernel
    boxBlurKernel := kernels.BoxBlurKernel(3)
    output := imgmod.ApplyKernel(boxBlurKernel, img)

    // Save the processed image to disk
    err = imgmod.SaveImage(output, "NEW_FILE_PATH")
    if err != nil {
        fmt.Println(err.Error())
    }

    fmt.Println("Image processing complete!")
}
```

💡 **Tip:** Replace `FILE_PATH` and `NEW_FILE_PATH` with your actual file paths to load and save images.

## Premade Kernels

imgmod provides several built-in kernels for common image processing tasks:

- 🔲 **Box Blur** - Smooth out image noise with a simple box blur.
- 🌫️ **Gaussian Blur** - A more advanced blur that preserves edges better than the box blur.
- 🔍 **Laplacian** - Highlight edges and detect features in your images.

## Upcoming Features

Here’s a sneak peek of what’s on the horizon for imgmod:

- 🔼 **Upscaling** - Improve image quality by increasing resolution (can be done with kernels).
- 🔽 **Downscaling** - Reduce image resolution without losing key details (can be done with kernels).
- ✂️ **Cropping** - Easily crop your images to focus on specific areas.

