package main

import (
	"fmt"
	"time"

	"github.com/sudheej/gochi/engine"
	"github.com/sudheej/gochi/intro"
)

func main() {
	intro.OpeningArt()
	fmt.Println("Testing yaml feature")
	engine.YamlParser(".gochi.yml")
	start := time.Now()
	engine.MavenRunner()
	elapsed := time.Since(start)
	fmt.Printf("Total time taken %d sec", elapsed.Milliseconds()/1000)
}
