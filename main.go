package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/theladbiblegroup/cloudflarebeat/beater"
)

func main() {
	err := beat.Run("cloudflarebeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
