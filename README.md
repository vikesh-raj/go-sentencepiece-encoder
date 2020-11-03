Sentence Piece Encoder
======================

This is pure go implementation of the sentencepiece encoder.
It takes a sentencepiece model and tokenizes it.

Example:

```go

import "github.com/vikesh-raj/go-sentencepiece-encoder/sentencepiece"

text := "This is a sample text"
spm, _ := sentencepiece.NewSentencepieceFromFile("spm.model", false)
tokens := spm.Tokenize(text)

```
