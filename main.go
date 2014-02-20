package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/geetarista/vindinium-starter-go/vindinium"
)

var (
	client vindinium.Client
	count  int
	games  int
)

func init() {
	flag.StringVar(&client.Server, "s", "http://vindinium.org", "server")
	flag.StringVar(&client.Key, "k", "", "api key")
	flag.StringVar(&client.Mode, "m", "arena", "mode (arena or training)")
	flag.StringVar(&client.BotName, "b", "fighter", "name of bot to use")
	flag.IntVar(&count, "c", 1, "number of games or turns to play")
	flag.BoolVar(&client.RandomMap, "r", true, "use random map (useful for training against same map)")
	flag.BoolVar(&client.Debug, "d", false, "debug output")
	flag.Usage = func() {
		fmt.Printf("Usage %s [FLAGS]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
}

func main() {
	if client.Key == "" {
		fmt.Println("API key required. Pass using the -k flag. Use -h to view help.")
		os.Exit(1)
	}

	if client.Mode == "training" {
		games = 1
		client.Turns = count
	} else {
		games = count
		client.Turns = 300
	}

	client.Setup()
	for i := 0; i < games; i++ {
		if err := client.Start(); err != nil {
			panic(err.Error())
		}
		if err := client.Play(); err != nil {
			panic(err.Error())
		}
	}
	fmt.Printf("Finished %d games.", games)
}
