package main

import (
	"flag"
	"log"

	"go.i3wm.org/i3"
)

func main() {

	separator := flag.String("separator", "|", "Symbol used to separate icons")
	config := flag.String("config", "", "JSON configuration file with map of icons")
	flag.Parse()

	if *config == "" {
		log.Fatal("JSON configuration file was not provided")
	}

	c := Config{separator: *separator}
	c.init(*config)
	recv := i3.Subscribe(i3.WindowEventType)
	for recv.Next() {
		tree, err := i3.GetTree()
		if err != nil {
			log.Println(err)
		}
		c.iterate(tree.Root)
	}
}
