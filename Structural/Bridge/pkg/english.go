package language

type EnglishGreeting struct {}

func (eg *EnglishGreeting) Hello() string {
	return "Hello"
}

func (eg *EnglishGreeting) Thanks() string {
	return "Thanks"
}