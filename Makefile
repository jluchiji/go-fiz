
all: y


y: fiz.nn.go y.go
	go build -o fiz fiz.nn.go y.go

fiz.nn.go: fiz.l
	nex fiz.l

y.go: fiz.y
	go tool yacc fiz.y

clean:
	rm fiz.nn.go fiz y.go y.output
