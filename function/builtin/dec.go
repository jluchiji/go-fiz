/**
 * function/builtin/dec.go
 *
 * @author  Denis Luchkin-Zhou <wyvernzora@gmail.com>
 * @license MIT
 */

package builtin
import "github.com/jluchiji/go-fiz/base"


func Dec(args []base.AstNode, context []int) int {
  result := args[0].Evaluate(context) - 1;
  if (result < 0) {
    panic("FIZ_DEC_ZERO")
  }
  return result
}
