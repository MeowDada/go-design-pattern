package language

type ChineseGreeting struct {}

func (cg *ChineseGreeting) Hello() string {
	return "你好"
}

func (ch *ChineseGreeting) Thanks() string {
	return "多謝"
}