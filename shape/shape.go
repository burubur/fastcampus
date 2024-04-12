package shape

type Rectangle struct {
	Multipoint []Point
}

type Point struct {
	X float64
	Y float64
}

func CreateRectangleDefault() Rectangle {
	return Rectangle{
		[]Point{
			{0, 0},
			{0, 5},
			{5, 0},
			{5, 5},
		},
	}
}

func (r *Rectangle) Scale(ratio float64) *Rectangle {
	r.Multipoint = []Point{
		{r.Multipoint[0].X, r.Multipoint[0].Y},
		{r.Multipoint[1].X, r.Multipoint[1].Y * ratio},
		{r.Multipoint[2].X * ratio, r.Multipoint[2].Y},
		{r.Multipoint[3].X, r.Multipoint[3].Y * ratio},
	}

	return r
}
