
package area

type Circletype interface {

	Area() float64
	Circumference() float64
}

type Myvalue struct {
	Radius float64
}

func (m Myvalue) Area() float64 {

	return 3.14*m.Radius * m.Radius;
}

func (m Myvalue) Circumference() float64 {

	return 2*3.14* m.Radius;
}



