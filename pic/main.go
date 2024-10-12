package main

import (
    "golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
    // Create a slice to hold the grayscale image data
    image := make([][]uint8, dy)

    // Loop through each row of the image
    for y := range image {
        // Create a slice for each row to hold pixel values
        image[y] = make([]uint8, dx)

        // Loop through each pixel in the row
        for x := range image[y] {
            // Replace this with your desired pixel processing function
            // Choose one or experiment with different options:

            // Option 1: Grayscale gradient (x-axis)
            // newValue := uint8(x * 255 / (dx - 1))

            // Option 2: Grayscale gradient (y-axis)
            // newValue := uint8(y * 255 / (dy - 1))

            // Option 3: Checkerboard pattern
            // newValue := uint8((x % 2) ^ (y % 2)) * 255

            // Option 4: Simple spiral (adjust based on dx and dy)
            centerX, centerY := dx/2, dy/2
            distance := uint8((abs(x-centerX) + abs(y-centerY)) * 255 / (max(dx, dy) - 1))
            newValue := distance

            // Set the pixel value (0 for black, 255 for white)
            image[y][x] = newValue
        }
    }

    return image
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func main() {
    pic.Show(Pic)
}