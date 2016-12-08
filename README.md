# Go-FIZ
The FIZ is a very LISP-y scripting language that is basically useless. This is a a Go
reimplementation of [jluchiji/fiz][0], which was originally written in C++ as a part of CS252
Operating Systems course at Purdue University.

Built using [blynn/nex][1] instead of Flex/Bison.


## Setup
```
$ go get github.com/blynn/nex
$ go get github.com/jluchiji/go-fiz
$ cd $GOPATH/src/github.com/jluchiji/go-fiz && make install
```


## FIZ Language

### Builtin Functions
| Function | Description |
| -------- | ----------- |
| `inc`    | Increases the argument by 1. |
| `dec`    | Decreases the argument by 1. Prints an error on attempt to decrease 0. |
| `ifz`    | If first argument is not 0, evaluates and returns the second argument; otherwise evaluates and returns the third argument. |
| `halt`   | Accepts no arguments. Terminates the interpreter. |

**Note**: New builtin functions can be added easily if needed. For details, check `function/registry.go`.

### Example Scripts
*The following scripts were taken from the course web-page*

```
(dec
  (ifz 0 4 (inc 1))
)                            ; interpreter prints 3
```

```
(define (add x y)
  (ifz y x
    (add
      (inc x)
      (dec y)
    )
  )
)
(add 1 2)                    ; interpreter prints 3
```


## Testing

This implementation has been tested against the test suite included with the original
[jluchiji/fiz][0] implementation and passed all tests.


## License
### The MIT License

Copyright (c) 2015-16 Denis Luchkin-Zhou

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.


[0]: https://github.com/jluchiji/fiz
[1]: http://crypto.stanford.edu/~blynn/nex/
