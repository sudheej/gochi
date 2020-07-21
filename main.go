package main

import (
	"fmt"

	"github.com/sudheej/gochi/debugger"
	"github.com/sudheej/gochi/engine"
	"github.com/sudheej/gochi/intro"
)

func main() {
	intro.OpeningArt()
	fmt.Println("Testing yaml feature")
	engine.YamlParser(".gochi.yml")
	fmt.Println("Memory ", debugger.GetMemory())
	//start := time.Now()
	//engine.MavenRunner()
	//elapsed := time.Since(start)
	//fmt.Printf("Total time taken %d sec", elapsed.Milliseconds()/1000)
}
