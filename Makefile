.PHONY: test bench clean

sentencepiece/sentencepiece_model.pb.go: sentencepiece/sentencepiece_model.proto
	protoc --go_out=. $<


cmd/dumpspm/dumpspm: cmd/dumpspm/main.go
	cd cmd/dumpspm && go build

test:
	go test -cover -coverprofile=c.out ./sentencepiece && go tool cover -html=c.out -o coverage.html

bench:
	go test -benchmem ./sentencepiece -bench Benchmark.*

clean:
	rm -f *.out coverage.html cmd/dumpspm/dumpspm