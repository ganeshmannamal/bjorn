package pair

import "testing"

func TestCompare(t *testing.T) {
	image1File := "../../testdata/image1.png"
	image2File := "../../testdata/image2.png"
	image3File := "../../testdata/image2.jpg"
	image4File := "../../testdata/image4.png"

	p, err := NewImagePair(image1File, image2File)

	if err != nil {
		t.Errorf("could not create image pair %v", err)
	}

	p.Compare()

	if p.Score != 0 {
		t.Errorf("Score was incorrect, got : %0.2f, expected %0.2f", p.Score, 0.0)
	}

	p1, err := NewImagePair(image1File, image3File)

	if err != nil {
		t.Errorf("could not create image pair %v", err)
	}

	p1.Compare()

	if p1.Score != 0 {
		t.Errorf("Score was incorrect, got : %0.2f, expected %0.2f", p.Score, 0.0)
	}

	p2, err := NewImagePair(image1File, image4File)

	if err != nil {
		t.Errorf("could not create image pair %v", err)
	}

	p2.Compare()

	if p2.Score == 0 {
		t.Errorf("Score was incorrect, got : %0.2f, expected > %0.2f", p.Score, 0.0)
	}
}
