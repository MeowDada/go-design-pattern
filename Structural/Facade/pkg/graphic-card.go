package pc

import "fmt"

type GraphicCard struct {}

func (gc *GraphicCard) Render() {
	fmt.Println("Graphic card is rendering")
}