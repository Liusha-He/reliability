package basics

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	Width, Height int
}

func (r *Rectangle) GetWidth() int {
	return r.Width
}

func (r *Rectangle) SetWidth(width int) {
	r.Width = width
}

func (r *Rectangle) GetHeight() int {
	return r.Height
}

func (r *Rectangle) SetHeight(height int) {
	r.Height = height
}

func UsedSized(sized Sized) (int, int) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetHeight() * sized.GetWidth()

	return expectedArea, actualArea
}

type Square struct {
	Size int
}

func (s Square) Rectangle() Rectangle {
	return Rectangle{s.Size, s.Size}
}
