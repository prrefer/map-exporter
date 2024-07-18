package main

import (
	"compress/bzip2"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func getMap(mapName string) string {
	if !strings.ContainsRune(mapName, '_') {
		return fmt.Sprintf("bhop_%s", mapName)
	}

	return mapName
}

func downloadMap(mapName string) (err error) {
	mapUrl := fmt.Sprintf("http://main.fastdl.me/maps/%s.bsp.bz2", mapName)

	resp, err := http.Get(mapUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return errors.New("map not found")
	}

	output, err := os.Create(fmt.Sprintf("%s.bsp", mapName))
	if err != nil {
		return err
	}
	defer output.Close()

	fmt.Println("writing...")
	reader := bzip2.NewReader(resp.Body)
	if _, err := io.Copy(output, reader); err != nil {
		return err
	}

	return nil
}

func main() {
	fmt.Println("map-exporter | https://github.com/prrefer/map-exporter")
	for {
		fmt.Print("map: ")
		var input string
		fmt.Scan(&input)

		mapName := getMap(input)
		if err := downloadMap(mapName); err != nil {
			fmt.Printf("could not retrieve map: %s\n", err.Error())
		}
	}
}
