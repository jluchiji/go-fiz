/**
 * func/base.go
 *
 * 2016 (C) Denis Luchkin-Zhou <wyvernzora@gmail.com>
 * Released under MIT license
 */
import "fmt"


/**
 * Base class for all fiz functions
 */
type Func {
  argc   int
  name   string
}


/**
 * Virtual method for calling fiz function
 */
func (n Func) call(argv []AstNode, context []int) int { return 0; }
