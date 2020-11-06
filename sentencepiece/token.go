package sentencepiece

// Token holds a unit of a tokenized word
type Token struct {
	ID   int32
	Text string
}

type tokenOffset struct {
	id    int32
	text  string
	start int
	end   int
}
