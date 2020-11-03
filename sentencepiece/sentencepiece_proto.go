package sentencepiece

import (
	"fmt"
	"io/ioutil"

	"google.golang.org/protobuf/proto"
)

// NewSentencepieceFromFile creates sentencepiece from file.
func NewSentencepieceFromFile(filename string, lowercase bool) (Sentencepiece, error) {
	s := NewEmptySentencepiece(lowercase)
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return s, fmt.Errorf("Unable to read file : %s, err %v", filename, err)
	}
	var model ModelProto
	err = proto.Unmarshal(bytes, &model)
	if err != nil {
		return s, fmt.Errorf("Unable to read model file : %s, err %v", filename, err)
	}

	count := 0
	unknownIndex := int64(0)
	for i, piece := range model.GetPieces() {
		word := piece.GetPiece()
		if word == unknown {
			unknownIndex = int64(i)
		}
		s.insert(word, piece.GetScore(), int64(i))
		count++
	}

	s.SetUnknownIndex(unknownIndex)
	return s, nil
}
