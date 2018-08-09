
all: goasm

clean:
	@rm -f goasm
	@

goasm: goasm.go
	@#GOROOT="/home/yma/workspace/golang/go1.9.4"; PATH="/usr/local/bin:/usr/bin:/bin:/usr/local/sbin:/usr/sbin:/home/yma/bin:/home/yma/workspace/golang/go1.9.4/bin"; go version; go build fure_hide.go
	@GOROOT="/home/yma/workspace/golang/go1.11beta3"; PATH="/home/yma/workspace/golang/go1.11beta3/bin"; go version; GOOS=js GOARCH=wasm go build goasm.go table.go
	@

t: t.go
	@#GOROOT="/home/yma/workspace/golang/go1.9.4"; PATH="/usr/local/bin:/usr/bin:/bin:/usr/local/sbin:/usr/sbin:/home/yma/bin:/home/yma/workspace/golang/go1.9.4/bin"; go version; go build fure_hide.go
	@GOROOT="/home/yma/workspace/golang/go1.11beta3"; PATH="/home/yma/workspace/golang/go1.11beta3/bin"; go version; GOOS=js GOARCH=wasm go build -o test.wasm t.go
	@
