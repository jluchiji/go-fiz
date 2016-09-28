/**
 * func/user.go
 *
 * 2016 (C) Denis Luchkin-Zhou <wyvernzora@gmail.com>
 * Released under MIT license
 */
import "fmt"


/**
 * User defined function
 */
type UserFunc {
  Func

  args  []string
  body  AstNode
}


/**
 * Override the execution function for the user-defined function
 */
func (this UserFunc) call(args []AstNode, context []int) {

  /* Check the number of arguments */
  argc := len(args);
  if (argc != this.argc) { panic("FIZ_ARGNUM"); }

  /* Evaluate all arguments */
  params := [argc]int;

}
