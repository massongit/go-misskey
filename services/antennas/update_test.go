package antennas_test

import (
	"log"
	"os"

	"github.com/yitsushi/go-misskey"
)

func ExampleService_Update() {
	client := misskey.NewClient("https://slippy.xyz", os.Getenv("MISSKEY_TOKEN"))

	antenna, err := client.Antennas().Show("8dbpybhulw")
	if err != nil {
		log.Println(err)
		return
	}

	update := antenna.Antenna()
	update.Keywords = append(update.Keywords, []string{"addition"})

	_, err = client.Antennas().UpdateAntenna(&update)
	if err != nil {
		log.Println(err)
		return
	}
}
