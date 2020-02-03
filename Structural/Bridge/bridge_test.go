package bridge

import (
	"testing"
)

func TestChineseViaEmail(t *testing.T) {
	chinese := NewChinese()
	viaEmail := NewViaEmail(chinese)
	viaEmail.Greet()
}

func TestChineseViaTalk(t *testing.T) {
	chinese := NewChinese()
	viaTalk := NewViaTalk(chinese)
	viaTalk.Greet()
}

func TestEnglishViaEmail(t *testing.T) {
	english := NewEnglish()
	viaEmail := NewViaEmail(english)
	viaEmail.Greet()
}

func TestEnglishViaTalk(t *testing.T) {
	english := NewEnglish()
	viaTalk := NewViaTalk(english)
	viaTalk.Greet()
}