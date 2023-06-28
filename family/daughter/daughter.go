package daughter

type Daughter struct {
	Name string
}

func (s Daughter) Data(name string) string {
	s.Name = "Daughter : " + name
	return s.Name
}