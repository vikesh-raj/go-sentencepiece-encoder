Sentence Piece Encoder
======================

This is pure go implementation of the sentencepiece encoder.
Create an encoder for the given sentencepiece model and then use
use the `Tokenize` function to split the input text into tokens.

Example:

```go

import "github.com/vikesh-raj/go-sentencepiece-encoder/sentencepiece"

text := "This is a sample text"
spm, _ := sentencepiece.NewSentencepieceFromFile("spm.model", false)
tokens := spm.Tokenize(text)

```
