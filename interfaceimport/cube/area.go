
package area

type Cubetype interface {

	Area() int64
	Permiter() int64
}

type Myvalue struct {
	Side int64
}

func (s Myvalue) Area() int64 {

	return s.Side*s.Side*s.Side;
}

func (s Myvalue) Permiter() int64{

	return  6*s.Side;
}



