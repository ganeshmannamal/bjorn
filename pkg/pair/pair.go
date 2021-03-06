package pair

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"time"
)

// Pair is a construct of 2 images to be compared, along with Score and time taken for comparison
// Score is a value between 0 to 1 indicating similarity of images.
// 0 indicates the images are identical, 1 indicates images are of different sizes
type Pair struct {
	Image1 image.Image
	Image2 image.Image
	Score  float64
	Time   float64
}

func NewImagePair(image1Path, image2Path string) (*Pair, error) {
	p := &Pair{}

	image1, err := readImage(image1Path)

	if err != nil {
		return nil, err
	}
	p.Image1 = image1

	image2, err := readImage(image2Path)

	if err != nil {
		return nil, err
	}
	p.Image2 = image2

	return p, nil
}

// Compare uses simple pixel based comparison to determine similarity between images.
// If the images are of different sizes, they are considered different with a Score of 1.
// Red, Green and Blue values of each pixel is compared to calculate the difference in pixel.
// Ratio of the sum of all pixel differences and Total no of pixels is user to determine the Score
func (p *Pair) Compare() {
	defer p.elapsed()() // deferred call to get execution time of Compare func
	if p.Image1.Bounds() != p.Image2.Bounds() {
		p.Score = 1
		return
	}

	bounds := p.Image2.Bounds()
	var sum int64
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r1, g1, b1, _ := p.Image1.At(x, y).RGBA()
			r2, g2, b2, _ := p.Image2.At(x, y).RGBA()
			sum += diff(r1, r2)
			sum += diff(g1, g2)
			sum += diff(b1, b2)
		}
	}

	nPixels := (bounds.Max.X - bounds.Min.X) * (bounds.Max.Y - bounds.Min.Y)

	p.Score = float64(sum) / (float64(nPixels) * 0xffff * 3)

	if p.Score < 0.01 {
		p.Score = 0
	}
}

// Elapsed helps calculate the time taken for each comparison
// A deferred call at the start of Compare() helps calculate the time taken
func (p *Pair) elapsed() func() {
	start := time.Now()
	return func() {
		p.Time = time.Since(start).Seconds()
	}
}

func diff(a, b uint32) int64 {
	if a > b {
		return int64(a - b)
	}
	return int64(b - a)
}

func readImage(file string) (image.Image, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return img, nil
}
