all: gomain.o goasm.asm goasm64.asm cgo

# get go assembler

gomain.o:
	@echo ":: Generates obj file as main.o -> gomain.o"
	go tool compile gomain.go

goasm.asm:
	@echo ":: Generates assembly from go -> goasm.asm"
	go tool compile -S gomain.go > goasm.asm

goasm64.asm:
	@echo ":: Generates assembly x86_64 from go -> goasm64.asm"
	GOOS=linux GOARCH=amd64 go tool compile -S gomain.go > goasm64.asm
	# go build -gcflags -S gomain.go > goasm64-2.asm # alternative

# cgo

cgo:
	env CGO_ENABLED=1 GOOS=linux go build cgo.go
	# variant with linking .o file:
	# CGO_ENABLED=1 GOOS=linux go build -buildmode=plugin -o path/to/module.so test.go

# clean

clean:
	rm gomain.o
	rm goasm.asm
	rm goasm64.asm
	rm goasm64-2.asm
	rm cgo
