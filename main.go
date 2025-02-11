package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var (
	mutex   sync.Mutex
)

func incrementCounter() int {
	mutex.Lock()
	defer mutex.Unlock()

	// Read current counter value from file
	data, err := os.ReadFile("counter.txt")
	if err != nil {
		// If file doesn't exist, start from 0
		data = []byte("0")
	}

	counter, err := strconv.Atoi(string(data))
	if err != nil {
		counter = 0
	}

	counter++

	// Write new value back to file
	err = os.WriteFile("counter.txt", []byte(fmt.Sprintf("%d", counter)), 0644)
	if err != nil {
		// Handle error, for now just print
		fmt.Println("Error writing counter:", err)
	}

	return counter
}

func getDigitPattern(d rune) ([14][10]int, bool) {
	patterns := map[rune][14][10]int{
		'0': {
			{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 0, 0, 0, 0, 1, 1, 1},
			{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
			{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
			{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
			{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
			{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
			{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
			{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
			{1, 1, 1, 0, 0, 0, 0, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
			{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
		},
		'1': {
			{0, 0, 0, 1, 1, 1, 0, 0, 0, 0},
			{0, 0, 1, 1, 1, 1, 0, 0, 0, 0},
			{0, 1, 1, 1, 1, 1, 0, 0, 0, 0},
			{1, 1, 1, 1, 1, 1, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 1, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 1, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 1, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 1, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 1, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 1, 0, 0, 0, 0},
			{0, 0, 0, 1, 1, 1, 0, 0, 0, 0},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
			{0, 1, 1, 1, 1, 1, 1, 1, 0, 0},
		},
		'2': {
			{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 0, 0, 0, 0, 0, 1, 1, 1},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
			{0, 0, 0, 0, 0, 0, 1, 1, 1, 0},
			{0, 0, 0, 0, 1, 1, 1, 1, 0, 0},
			{0, 0, 1, 1, 1, 1, 1, 0, 0, 0},
			{0, 1, 1, 1, 1, 0, 0, 0, 0, 0},
			{1, 1, 1, 0, 0, 0, 0, 0, 0, 0},
			{1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
		},
		'3': {
			{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 0, 0, 0, 0, 0, 1, 1, 1},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
			{0, 0, 0, 0, 0, 0, 1, 1, 1, 0},
			{0, 0, 0, 1, 1, 1, 1, 1, 0, 0},
			{0, 0, 0, 1, 1, 1, 1, 1, 1, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
			{0, 0, 0, 0, 0, 0, 0, 0, 1, 1},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
			{1, 1, 0, 0, 0, 0, 0, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
			{0, 1, 1, 1, 1, 1, 1, 1, 0, 0},
			{0, 0, 1, 1, 1, 1, 1, 0, 0, 0},
		},
		'4': {
			{0, 0, 0, 0, 0, 1, 1, 1, 0, 0},
			{0, 0, 0, 0, 1, 1, 1, 1, 0, 0},
			{0, 0, 0, 1, 1, 1, 1, 1, 0, 0},
			{0, 0, 1, 1, 0, 1, 1, 1, 0, 0},
			{0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
			{1, 1, 0, 0, 0, 1, 1, 1, 0, 0},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{0, 0, 0, 0, 0, 1, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 1, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 1, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 1, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 1, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 1, 1, 1, 0, 0},
		},
		'5': {
			{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 1, 1, 0, 0, 0, 0, 0, 0, 0},
			{1, 1, 1, 1, 1, 1, 1, 1, 0, 0},
			{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
			{0, 0, 0, 0, 0, 0, 0, 0, 1, 1},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
			{1, 1, 0, 0, 0, 0, 0, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
			{0, 1, 1, 1, 1, 1, 1, 1, 0, 0},
			{0, 0, 1, 1, 1, 1, 1, 0, 0, 0},
		},
		'6': {
			{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
			{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
			{1, 1, 1, 0, 0, 0, 0, 1, 1, 1},
			{1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 1, 1, 1, 1, 1, 1, 1, 0, 0},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
			{1, 1, 0, 0, 0, 0, 0, 1, 1, 1},
			{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
			{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
			{1, 1, 1, 0, 0, 0, 0, 1, 1, 1},
			{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
			{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
			{0, 0, 0, 1, 1, 1, 1, 0, 0, 0},
		},
		'7': {
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1},
			{0, 0, 0, 0, 0, 0, 1, 1, 1, 0},
			{0, 0, 0, 0, 0, 1, 1, 1, 0, 0},
			{0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 0, 0, 1, 1, 1, 0, 0, 0, 0},
			{0, 0, 1, 1, 1, 0, 0, 0, 0, 0},
			{0, 1, 1, 1, 0, 0, 0, 0, 0, 0},
			{0, 1, 1, 1, 0, 0, 0, 0, 0, 0},
			{0, 1, 1, 1, 0, 0, 0, 0, 0, 0},
			{0, 1, 1, 1, 0, 0, 0, 0, 0, 0},
			{0, 1, 1, 1, 0, 0, 0, 0, 0, 0},
			{0, 1, 1, 1, 0, 0, 0, 0, 0, 0},
		},
		'8': {
			{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 0, 0, 0, 0, 1, 1, 1},
			{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
			{1, 1, 1, 0, 0, 0, 0, 1, 1, 1},
			{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
			{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
			{1, 1, 1, 0, 0, 0, 0, 1, 1, 1},
			{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
			{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
			{1, 1, 1, 0, 0, 0, 0, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
			{0, 0, 1, 1, 1, 1, 1, 1, 0, 0},
		},
		'9': {
			{0, 1, 1, 1, 1, 1, 1, 1, 1, 0},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 0, 0, 0, 0, 1, 1, 1},
			{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
			{1, 1, 0, 0, 0, 0, 0, 0, 1, 1},
			{1, 1, 1, 0, 0, 0, 0, 1, 1, 1},
			{0, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{0, 0, 1, 1, 1, 1, 1, 1, 1, 1},
			{0, 0, 0, 0, 0, 0, 0, 0, 1, 1},
			{0, 0, 0, 0, 0, 0, 0, 0, 1, 1},
			{1, 1, 0, 0, 0, 0, 0, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
			{0, 1, 1, 1, 1, 1, 1, 1, 0, 0},
			{0, 0, 1, 1, 1, 1, 0, 0, 0, 0},
		},
	}

	return patterns[d], true
}

func drawNumber(num int) *image.RGBA {
	// Convert number to string
	numStr := fmt.Sprintf("%d", num)

	// Calculate width based on number of digits
	digitWidth := 10 // Each digit takes 10 pixels (10 * scale(1))
	spacing := 3     // Space between digits
	padding := 3     // Padding on all sides
	totalWidth := len(numStr)*digitWidth + (len(numStr)-1)*spacing + 2*padding
	height := 14 + 2*padding // 14 rows + padding top/bottom

	// Create an image with dynamic width
	img := image.NewRGBA(image.Rect(0, 0, totalWidth, height))

	// Fill background with white
	for x := 0; x < totalWidth; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, color.White)
		}
	}

	// Draw each digit
	x := padding // Starting x position with padding
	scale := 1   // Size of each pixel in the digit pattern

	for _, d := range numStr {
		if pattern, ok := getDigitPattern(d); ok {
			for i := 0; i < 14; i++ {
				for j := 0; j < 10; j++ {
					if pattern[i][j] == 1 {
						// Draw a filled square for each black pixel
						for px := 0; px < scale; px++ {
							for py := 0; py < scale; py++ {
								img.Set(x+j*scale+px, padding+i*scale+py, color.Black)
							}
						}
					}
				}
			}
			x += scale * 10 + spacing // Space between digits
		}
	}

	return img
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	currentNum := incrementCounter()

	// Generate the image
	img := drawNumber(currentNum)

	// Set content type header
	w.Header().Set("Content-Type", "image/png")

	// Encode and send the image
	png.Encode(w, img)
}

func main() {
	// Set up HTTP route
	http.HandleFunc("/", handleRequest)

	// Start server
	fmt.Println("Server starting on :8080...")
	http.ListenAndServe(":8080", nil)
}
