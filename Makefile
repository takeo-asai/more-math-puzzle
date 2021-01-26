all: linux windows
clean:
	rm build/* 
linux:
	GOOS=linux ls
windows:
	GOOS=windows ls
test:
	go test -v "github.com/takeo-asai/more-math-puzzle/..."
fmt:
	go fmt github.com/takeo-asai/more-math-puzzle/...
