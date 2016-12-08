/**
 * function/builtin/halt.go
 *
 * @author  Denis Luchkin-Zhou <wyvernzora@gmail.com>
 * @license MIT
 */

package builtin
import "github.com/jluchiji/go-fiz/base"


func Halt(args []base.AstNode, context []int) int {
  panic("FIZ_HALT")
}
