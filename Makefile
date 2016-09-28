
all: y


y: fiz.nn.go y.go
	go build -o y fiz.nn.go y.go

fiz.nn.go:
	nex fiz.l

y.go:
	go tool yacc fiz.y

clean:
	rm fiz.nn.go y y.go y.output
