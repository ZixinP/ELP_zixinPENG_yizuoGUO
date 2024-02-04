package image_handle

import (
	"image"
	"image/color"
	"image/draw"
	
	// "github.com/jeasonstudio/GaussianBlur"
	"github.com/disintegration/imaging"
)

func expandImage(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	newRect := image.Rect(0, 0, bounds.Dx()+2, bounds.Dy()+2)
	newImg := image.NewRGBA(newRect)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			newImg.Set(x+1, y+1, img.At(x, y))
		}
	}
	return newImg
}

// type Pgxel struct {
// 	X, Y  int
// 	Value uint8
// }

// func nonMaxSuppression(sobel *image.Gray) *image.Gray {
// 	bounds := sobel.Bounds()
// 	suppressed := image.NewGray(bounds)

// 	for y := bounds.Min.Y + 1; y < bounds.Max.Y-1; y++ {
// 		for x := bounds.Min.X + 1; x < bounds.Max.X-1; x++ {
// 			// Get the gradient direction for the current pgxel.
// 			gx := sobel.GrayAt(x-1, y).Y - sobel.GrayAt(x+1, y).Y
// 			gy := sobel.GrayAt(x, y-1).Y - sobel.GrayAt(x, y+1).Y

// 			// Determine the pgxels to compare based on the gradient direction.
// 			var comparePgxels [2]Pgxel
// 			if (gy <= 0 && gx > -gy) || (gy >= 0 && gx < -gy) {
// 				comparePgxels = [2]Pgxel{{X: x - 1, Y: y, Value: 0}, {X: x + 1, Y: y, Value: 0}}
// 			} else if (gx > 0 && gx <= -gy) || (gx < 0 && gx >= -gy) {
// 				comparePgxels = [2]Pgxel{{X: x, Y: y - 1, Value: 0}, {X: x, Y: y + 1, Value: 0}}
// 			} else if (gx <= 0 && gx > gy) || (gx >= 0 && gx < gy) {
// 				comparePgxels = [2]Pgxel{{X: x - 1, Y: y - 1, Value: 0}, {X: x + 1, Y: y + 1, Value: 0}}
// 			} else if (gy < 0 && gx <= gy) || (gy > 0 && gx >= gy) {
// 				comparePgxels = [2]Pgxel{{X: x + 1, Y: y - 1, Value: 0}, {X: x - 1, Y: y + 1, Value: 0}}
// 			}

// 			// Suppress the current pgxel if it's not the maximum.
// 			current := sobel.GrayAt(x, y).Y
// 			if current >= sobel.GrayAt(comparePgxels[0].X, comparePgxels[0].Y).Y &&
// 				current >= sobel.GrayAt(comparePgxels[1].X, comparePgxels[1].Y).Y {
// 				suppressed.SetGray(x, y, color.Gray{Y: current})
// 			}
// 		}
// 	}

// 	return suppressed
// }

func handle_image(img image.Image, Parameter int) image.Image {

	img = imaging.Blur(img, 1)
	img = expandImage(img)

	// Convert the image to grayscale.
	gray := image.NewGray(img.Bounds())
	draw.Draw(gray, img.Bounds(), img, image.Point{}, draw.Src)

	// Apply the Sobel operator.
	sobel := image.NewGray(img.Bounds())

	type Tuple struct {
		X, Y uint8
	}
	// gardient_map := make([][]int, gray.Bounds().Max.X)
	// for i := range gardient_map {
	// 	gardient_map[i] = make([]int, gray.Bounds().Max.Y)
	// }

	gxgy_map := make([][]Tuple, gray.Bounds().Max.X)
	for i := range gxgy_map {
		gxgy_map[i] = make([]Tuple, gray.Bounds().Max.Y)
	}

	Gup := make([][]uint8, gray.Bounds().Max.X)
	for i := range Gup {
		Gup[i] = make([]uint8, gray.Bounds().Max.Y)
	}

	Gdown := make([][]uint8, gray.Bounds().Max.X)
	for i := range Gdown {
		Gdown[i] = make([]uint8, gray.Bounds().Max.Y)
	}

	for y := 1; y < gray.Bounds().Max.Y-1; y++ {
		for x := 1; x < gray.Bounds().Max.X-1; x++ {
			gx := -gray.GrayAt(x-1, y-1).Y - 2*gray.GrayAt(x-1, y).Y - gray.GrayAt(x-1, y+1).Y +
				gray.GrayAt(x+1, y-1).Y + 2*gray.GrayAt(x+1, y).Y + gray.GrayAt(x+1, y+1).Y
			gy := -gray.GrayAt(x-1, y-1).Y - 2*gray.GrayAt(x, y-1).Y - gray.GrayAt(x+1, y-1).Y +
				gray.GrayAt(x-1, y+1).Y + 2*gray.GrayAt(x, y+1).Y + gray.GrayAt(x+1, y+1).Y
			gxgy_map[x][y] = Tuple{gx, gy}
			sobel.SetGray(x, y, color.Gray{Y: uint8(sqrt(float64(gx*gx + gy*gy)))})

			// if (gy <= 0 && gx > -gy) || (gy >= 0 && gx < -gy) {
			// 	gardient_map[x][y] = 1
			// } else if (gx > 0 && gx <= -gy) || (gx < 0 && gx >= -gy) {
			// 	gardient_map[x][y] = 2
			// } else if (gx <= 0 && gx > gy) || (gx >= 0 && gx < gy) {
			// 	gardient_map[x][y] = 3
			// } else if (gy < 0 && gx <= gy) || (gy > 0 && gx >= gy) {
			// 	gardient_map[x][y] = 4
			// } else {
			// 	gardient_map[x][y] = 0
			// }
		
		}
	}

	var image_end image.Image
	for y := 1; y < gray.Bounds().Max.Y-1; y++ {
		for x := 1; x < gray.Bounds().Max.X-1; x++ {
			gx := gxgy_map[x][y].X
			gy := gxgy_map[x][y].Y
			if (gy <= 0 && gx > -gy) || (gy >= 0 && gx < -gy) {
				t := abs(gy / gx)
				Gup[x][y] = sobel.GrayAt(x, y+1).Y*(1-t) + sobel.GrayAt(x-1, y+1).Y*t
				Gdown[x][y] = sobel.GrayAt(x, y-1).Y*(1-t) + sobel.GrayAt(x+1, y-1).Y*t
			} else if (gx > 0 && gx <= -gy) || (gx < 0 && gx >= -gy) {
				t := abs(gx / gy)
				Gup[x][y] = sobel.GrayAt(x-1, y).Y*(1-t) + sobel.GrayAt(x-1, y+1).Y*t
				Gdown[x][y] = sobel.GrayAt(x+1, y).Y*(1-t) + sobel.GrayAt(x+1, y-1).Y*t
			} else if (gx <= 0 && gx > gy) || (gx >= 0 && gx < gy) {
				t := abs(gx / gy)
				Gup[x][y] = sobel.GrayAt(x-1, y).Y*(1-t) + sobel.GrayAt(x-1, y-1).Y*t
				Gdown[x][y] = sobel.GrayAt(x+1, y).Y*(1-t) + sobel.GrayAt(x+1, y+1).Y*t
			} else if (gy < 0 && gx <= gy) || (gy > 0 && gx >= gy) {
				t := abs(gy / gx)
				Gup[x][y] = sobel.GrayAt(x, y-1).Y*(1-t) + sobel.GrayAt(x-1, y-1).Y*t
				Gdown[x][y] = sobel.GrayAt(x, y+1).Y*(1-t) + sobel.GrayAt(x+1, y+1).Y*t
			}
		}

		// Suppress non-maximum pgxels.
		bounds := sobel.Bounds()
		suppressed := image.NewGray(bounds)
		for y := 1; y < gray.Bounds().Max.Y-1; y++ {
			for x := 1; x < gray.Bounds().Max.X-1; x++ {
				if sobel.GrayAt(x, y).Y >= Gup[x][y] && sobel.GrayAt(x, y).Y >= Gdown[x][y] {
					suppressed.SetGray(x, y, color.Gray{Y: sobel.GrayAt(x, y).Y})
				} else {
					suppressed.SetGray(x, y, color.Gray{Y: 0})
				}
			}
		}
		image_end = suppressed
	}
	return image_end
}

func sqrt(x float64) int {
	if x < 0 {
		return -int(-x)
	}
	return int(x)
}

func abs(x uint8) uint8 {
	if x < 0 {
		return -x
	}
	return x
}


// Version of using goroutine for each pgxel in handle_image function
/*
package image_handle

import (
	"image"
	"image/color"
	"image/draw"
	"github.com/disintegration/imaging"
)

func expandImage(img image.Image) *image.RGBA {
	bounds := img.Bounds()
	newRect := image.Rect(0, 0, bounds.Dx()+2, bounds.Dy()+2)
	newImg := image.NewRGBA(newRect)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			newImg.Set(x+1, y+1, img.At(x, y))
		}
	}
	return newImg
}

func handle_image(img image.Image, Parameter int) image.Image {

	img = imaging.Blur(img, 1)
	img = expandImage(img)

	// Convert the image to grayscale.
	gray := image.NewGray(img.Bounds())
	draw.Draw(gray, img.Bounds(), img, image.Point{}, draw.Src)

	// Apply the Sobel operator.
	sobel := image.NewGray(img.Bounds())

	type Tuple struct {
		X, Y uint8
	}

	gxgy_map := make([][]Tuple, gray.Bounds().Max.X)
	for i := range gxgy_map {
		gxgy_map[i] = make([]Tuple, gray.Bounds().Max.Y)
	}

	Gup := make([][]uint8, gray.Bounds().Max.X)
	for i := range Gup {
		Gup[i] = make([]uint8, gray.Bounds().Max.Y)
	}

	Gdown := make([][]uint8, gray.Bounds().Max.X)
	for i := range Gdown {
		Gdown[i] = make([]uint8, gray.Bounds().Max.Y)
	}

	wg_1 := sync.WaitGroup
	for y := 1; y < gray.Bounds().Max.Y-1; y++ {
		for x := 1; x < gray.Bounds().Max.X-1; x++ {
			wg_1.Add(1)
			go func(x, y int , sobel *image.Gray, gxgy_map [][]Tuple) {
				defer wg_1.Done()
				gx := -gray.GrayAt(x-1, y-1).Y - 2*gray.GrayAt(x-1, y).Y - gray.GrayAt(x-1, y+1).Y +
					gray.GrayAt(x+1, y-1).Y + 2*gray.GrayAt(x+1, y).Y + gray.GrayAt(x+1, y+1).Y
				gy := -gray.GrayAt(x-1, y-1).Y - 2*gray.GrayAt(x, y-1).Y - gray.GrayAt(x+1, y-1).Y +
					gray.GrayAt(x-1, y+1).Y + 2*gray.GrayAt(x, y+1).Y + gray.GrayAt(x+1, y+1).Y
				gxgy_map[x][y] = Tuple{gx, gy}
				sobel.SetGray(x, y, color.Gray{Y: uint8(sqrt(float64(gx*gx + gy*gy)))})
			}(x, y, sobel, gxgy_map)	
		}
	}
	wg_1.Wait()

	wg_2 := sync.WaitGroup
	var image_end image.Image
	for y := 1; y < gray.Bounds().Max.Y-1; y++ {
		for x := 1; x < gray.Bounds().Max.X-1; x++ {
			wg_2.Add(1)
			go func(x, y int, gxgy_map [][]Tuple, sobel *image.Gray, Gup, Gdown [][]uint8) {
				defer wg_2.Done()
				gx := gxgy_map[x][y].X
				gy := gxgy_map[x][y].Y
				if (gy <= 0 && gx > -gy) || (gy >= 0 && gx < -gy) {
					t := abs(gy / gx)
					Gup[x][y] = sobel.GrayAt(x, y+1).Y*(1-t) + sobel.GrayAt(x-1, y+1).Y*t
					Gdown[x][y] = sobel.GrayAt(x, y-1).Y*(1-t) + sobel.GrayAt(x+1, y-1).Y*t
				} else if (gx > 0 && gx <= -gy) || (gx < 0 && gx >= -gy) {
					t := abs(gx / gy)
					Gup[x][y] = sobel.GrayAt(x-1, y).Y*(1-t) + sobel.GrayAt(x-1, y+1).Y*t
					Gdown[x][y] = sobel.GrayAt(x+1, y).Y*(1-t) + sobel.GrayAt(x+1, y-1).Y*t
				} else if (gx <= 0 && gx > gy) || (gx >= 0 && gx < gy) {
					t := abs(gx / gy)
					Gup[x][y] = sobel.GrayAt(x-1, y).Y*(1-t) + sobel.GrayAt(x-1, y-1).Y*t
					Gdown[x][y] = sobel.GrayAt(x+1, y).Y*(1-t) + sobel.GrayAt(x+1, y+1).Y*t
				} else if (gy < 0 && gx <= gy) || (gy > 0 && gx >= gy) {
					t := abs(gy / gx)
					Gup[x][y] = sobel.GrayAt(x, y-1).Y*(1-t) + sobel.GrayAt(x-1, y-1).Y*t
					Gdown[x][y] = sobel.GrayAt(x, y+1).Y*(1-t) + sobel.GrayAt(x+1, y+1).Y*t
				}
			}(x, y, gxgy_map, sobel, Gup, Gdown)
		}
	
		// Suppress non-maximum pgxels.
		bounds := sobel.Bounds()
		suppressed := image.NewGray(bounds)
		for y := 1; y < gray.Bounds().Max.Y-1; y++ {
			for x := 1; x < gray.Bounds().Max.X-1; x++ {
				if sobel.GrayAt(x, y).Y >= Gup[x][y] && sobel.GrayAt(x, y).Y >= Gdown[x][y] {
					suppressed.SetGray(x, y, color.Gray{Y: sobel.GrayAt(x, y).Y})
				} else {
					suppressed.SetGray(x, y, color.Gray{Y: 0})
				}
			}
		}	
		image_end = suppressed
	}
	wg_2.Wait()
	return image_end
}



func sqrt(x float64) int {
	if x < 0 {
		return -int(-x)
	}
	return int(x)
}

func abs(x uint8) uint8 {
	if x < 0 {
		return -x
	}
	return x
}

*/
