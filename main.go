package main

import (
	"flag"
	"log"
)

func main() {
	t := mustToken()

	// tgClient = telegram.New(token)

	// fetcher = fetcher.New(tgClient )

	// processor = processor.New(tgClient)

	// consumer.Start(fetcher, processor)
}

func mustToken() string {
	token := flag.String()
}
