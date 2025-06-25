package testgigacode

type GigaTest1 struct {
	Test1   string
	Test2   string
	Number  int
	Boolean bool
	Rune    rune
}

func (g *GigaTest1) GetTest1() string {
	if g.Test1 == "" {
		return "Test1 is empty"
	} else {
		return g.Test1
	}
}
