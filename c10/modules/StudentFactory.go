package modules

// uppercase and lower case
type student struct {
	Name  string
	Score float64
}

func NewStudent(n string, s float64) *student {
	return &student{
		Name:  n,
		Score: s,
	}
}

func (s *student) GetScore() float64 {
	return s.Score
}
