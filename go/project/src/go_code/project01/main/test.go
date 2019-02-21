package main
//首字母小写的student
type student struct {
	Name string
	//首字母小写的score
	score float64
}
func NewStudent(n string, s float64) *student {
	return &student{
		Name: n,
		score: s,
	}
}
func (student *student) GetScore() float64 {
	return student.score
}