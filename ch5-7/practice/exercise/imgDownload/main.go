package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const filePath = "./ch5-7/practice/exercise/imgDownload/img.jpeg"

func main() {
	url := "https://img.tsn.ua/cached/1533898791/tsn-3ad8a7940cc99f147f48233aa7502420/thumbs/585xX/ac/00/2f07e665934361372c1544e1591700ac.jpeg"

	response, e := http.Get(url)
	if e != nil {
		panic(e)
	}

	defer response.Body.Close()
	path, _ := filepath.Abs(filePath)
	file, err := os.Create(path)
	if err != nil {
		panic(e)
	}

	_, err = io.Copy(file, response.Body)
	if err != nil {
		panic(err)
	}
	file.Close()
	fmt.Println("Success!")
}
