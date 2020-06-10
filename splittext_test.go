package gofpdf_test

import (
	"github.com/jung-kurt/gofpdf"
	"testing"
)

func TestFpdf_SplitText_ShouldNotCreateNewLineWhenCharacterIsZeroWidthCharacter(t *testing.T) {
	var pdf *gofpdf.Fpdf
	pdf = gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	// Load font that contains zero-width character
	pdf.AddUTF8Font("THSarabun", "", "./font/THSarabun.ttf")
	pdf.SetFont("THSarabun", "", 14)

	// " ี" and " ้" are zero-width character
	// this string `กี้` is very short (compare to 100mm width), SplitText should return one-length slice
	lines := pdf.SplitText("กี้", 100.00)
	if len(lines) != 1 {
		t.Errorf("SplitText should return one-length slice, got %v length slice", len(lines))
	}
}
