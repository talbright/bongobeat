package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/talbright/bongobeat/beater"
)

func main() {
	err := beat.Run("bongobeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
