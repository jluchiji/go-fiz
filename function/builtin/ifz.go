/**
 * function/builtin/ifz.go
 *
 * @author  Denis Luchkin-Zhou <wyvernzora@gmail.com>
 * @license MIT
 */

package builtin
import "github.com/jluchiji/go-fiz/base"


func Ifz(args []base.AstNode, context []int) int {
  cond := args[0].Evaluate(context)
  var value int
  if (cond == 0) {
    value = args[1].Evaluate(context)
  } else {
    value = args[2].Evaluate(context)
  }
  return value
}
