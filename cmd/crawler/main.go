package main

import (
	"log"

	"flag"

	"github.com/baramsivaramireddy/single-site-crawler/internal/phase0"
)

func main() {

	log.Println("Starting Crawler")

	var baseURL = flag.String("url", "", "Site we are going to ")
	flag.Parse()

	phase0.Run(*baseURL)
}
