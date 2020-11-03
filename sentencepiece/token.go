package sentencepiece

// Token holds a unit of a tokenized word
type Token struct {
	ID   int64
	Text string
}

type tokenOffset struct {
	id    int64
	text  string
	start int
	end   int
}
