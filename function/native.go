/**
 * fn/native.go
 *
 * @author  Denis Luchkin-Zhou <wyvernzora@gmail.com>
 * @license MIT
 */

package function
import "github.com/jluchiji/go-fiz/base"


type Native struct {
  args  []string
  body  func([]base.AstNode, []int) int
}


func (this *Native) GetArgs() []string {
  return this.args
}


func (this *Native) Call(args []base.AstNode, context []int) int {
  if (len(args) != len(this.args)) {
    panic("FIZ_ARGNUM")
  }
  return this.body(args, context)
}


func NewNative(args []string, body func([]base.AstNode, []int) int) *Native {
  return &Native{
    args,
    body,
  }
}
