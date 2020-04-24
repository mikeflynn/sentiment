package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/mikeflynn/sentiment"
)

func writeTrainingFile(dir string, text string) error {
	filename := fmt.Sprintf("cli_%d.txt", time.Now().Unix())

	data := []byte(text)
	err := ioutil.WriteFile("./datasets/train/"+dir+"/"+filename, data, 0644)

	return err
}

func main() {
	mode := flag.String("mode", "", "positive, negative, or build")
	text := flag.String("text", "", "Text to add as trainig for positive or negative sentiment.")

	flag.Parse()

	switch *mode {
	case "positive":
		if *text == "" {
			fmt.Println("Please enter the text of this training document.")
			os.Exit(1)
		}

		err := writeTrainingFile("pos", *text)
		if err != nil {
			fmt.Println("Training file could not be created.")
			os.Exit(1)
		}
	case "negative":
		if *text == "" {
			fmt.Println("Please enter the text of this training document.")
			os.Exit(1)
		}

		err := writeTrainingFile("neg", *text)
		if err != nil {
			fmt.Println("Training file could not be created.")
			os.Exit(1)
		}
	case "build":
		sentiment.Train()
		fmt.Println("Please run > go-bindata -prefix \"/tmp/.sentiment/\" -pkg sentiment /tmp/.sentiment/model.json")
	default:
		fmt.Println("You must select a mode: positive, negative, or build.")
		os.Exit(1)
	}

	os.Exit(0)
}
