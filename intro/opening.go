package intro

import (
	"fmt"

	"github.com/common-nighthawk/go-figure"
)

//OpeningArt would render the opening page in CLI
func OpeningArt() {
	fmt.Println("")
	figure.NewFigure("GOCHI", "banner3-D", true).Print()
}
