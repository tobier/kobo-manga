package main

import (
	log "github.com/sirupsen/logrus"

	"github.com/tobier/kobo-manga/cmd"
)

func main() {
    log.SetFormatter(&log.TextFormatter{
        ForceColors: true,
    })

    if err := cmd.Run(); err != nil {
        log.Fatal(err.Error())
    }
}