package language

type JapaneseGreeting struct {}

func (jg *JapaneseGreeting) Hello() string {
	return "こんにちは"
}

func (jg *JapaneseGreeting) Thanks() string {
	return "ありがとう"
}