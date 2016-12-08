/**
 * function/builtin/inc.go
 *
 * @author  Denis Luchkin-Zhou <wyvernzora@gmail.com>
 * @license MIT
 */

package builtin
import "github.com/jluchiji/go-fiz/base"


func Inc(args []base.AstNode, context []int) int {
  return args[0].Evaluate(context) + 1
}
