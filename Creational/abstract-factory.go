package main

import (
	"fmt"
)

type NoteBook interface {
	Brand() string
}

type NoteBookFactory interface {
	ManufactureProduct() NoteBook
}

type AsusNoteBook struct {}

func (nb *AsusNoteBook) Brand() string {
	return "ASUS"
}

type AsusFactory struct {}

func (f *AsusFactory) ManufactureProduct() NoteBook {
	return new(AsusNoteBook)
}

type DellNoteBook struct {}

func (nb *DellNoteBook) Brand() string {
	return "DELL"
}

type DellFactory struct {}

func (f *DellFactory) ManufactureProduct() NoteBook {
	return new(DellNoteBook)
}

type AppleNoteBook struct {}

func (nb *AppleNoteBook) Brand() string {
	return "Apple"
}

type AppleFactory struct {}

func (f *AppleFactory) ManufactureProduct() NoteBook {
	return new(AppleNoteBook)
}

func buildFactory(manufacture string) NoteBookFactory {
	switch manufacture {
	case "Asus":
		return new(AsusFactory)
	case "Dell":
		return new(DellFactory)
	case "Apple":
		return new(AppleFactory)
	}
	return nil
}

func main() {

	if factory := buildFactory("Asus"); factory != nil {
		notebook := factory.ManufactureProduct()
		fmt.Println(notebook.Brand())
	}

	if factory := buildFactory("Dell"); factory != nil {
		notebook := factory.ManufactureProduct()
		fmt.Println(notebook.Brand())
	}

	if factory := buildFactory("Apple"); factory != nil {
		notebook := factory.ManufactureProduct()
		fmt.Println(notebook.Brand())
	}
}