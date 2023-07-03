package algos

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Perimetre() float64 {
	return 2 * (r.Height + r.Width)
}
