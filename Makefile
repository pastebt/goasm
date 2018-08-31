
all: goasm

clean:
	@rm -f goasm
	@

goasm: *.go
	@GOOS=js GOARCH=wasm go build goasm.go table.go date.go chart.go
	@

