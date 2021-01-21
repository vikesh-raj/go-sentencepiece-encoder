package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/vikesh-raj/go-sentencepiece-encoder/sentencepiece"
	"google.golang.org/protobuf/proto"
)

func dumpWords(filename string) error {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Unable to read file : %s, err %v", filename, err)
	}
	var model sentencepiece.ModelProto
	err = proto.Unmarshal(bytes, &model)
	if err != nil {
		return fmt.Errorf("Unable to read model file : %s, err %v", filename, err)
	}

	count := 0
	for i, piece := range model.GetPieces() {
		word := piece.GetPiece()
		fmt.Println(word, piece.GetScore(), "(", piece.GetType(), ")", i)
		count++
	}

	fmt.Println("Total words :", count)
	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Expected Sentencepiece model as argument.\n")
		os.Exit(1)
	}
	modelFile := os.Args[1]
	if err := dumpWords(modelFile); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
