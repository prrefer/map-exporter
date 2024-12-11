package main

import (
	"compress/bzip2"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func getMapName(mapName string) string {
	if !strings.ContainsRune(mapName, '_') {
		return fmt.Sprintf("bhop_%s", mapName)
	}

	return mapName
}

func downloadMap(mapName string) (err error) {
	mapUrl := fmt.Sprintf("http://main.fastdl.me/maps/%s.bsp.bz2", mapName)
	fmt.Println(mapUrl)

	fmt.Println("gettting " + mapName)
	resp, err := http.Get(mapUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Println("failure getting " + mapName)
		return
	}

	fmt.Println("creating...")
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
	if len(os.Args) < 2 {
		log.Fatalln("no arguments")
	}

	mapNames := strings.Split(os.Args[1], ",")
	for _, arg := range mapNames {
		mapName := getMapName(arg)
		err := downloadMap(mapName)
		if err != nil {
			log.Fatalln(err.Error())
		}
	}
}
