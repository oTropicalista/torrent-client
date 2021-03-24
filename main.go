package main

// TODO
// - abrir .torrent
// - baixar e madnar para arquivo

import (
	"log"
	"os"

	"github.com/oTropicalista/torrent-client/torrentfile"
)

func main() {
	inPath := os.Args[1]
	outPath := os.Args[2]

	tf, err := torrentfile.Open(inPath)
	if err != nil {
		log.Fatal(err)
	}
}
