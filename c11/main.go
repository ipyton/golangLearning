package main

type A interface {
	SayHello()
	Say()
}

type B struct {
}

func (this *B) SayHello() {

}
func (this B) Say() {

}
func main() {
	var b B
	var a A = &b
	a.SayHello()
}
