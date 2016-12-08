/**
 * ast/number.go
 *
 * @author  Denis Luchkin-Zhou <wyvernzora@gmail.com>
 * @license MIT
 */

package ast
import "github.com/jluchiji/go-fiz/base"


type Call struct {
  fn    base.Func
  argv  []base.AstNode
}


func (this *Call) Resolve(fn base.Func) {
  for _, v := range this.argv {
    v.Resolve(fn)
  }
}


func (this *Call) Evaluate(context []int) int {
  return this.fn.Call(this.argv, context)
}


func NewCall(fn base.Func, argv []base.AstNode) *Call {
  return &Call{
    fn,
    argv,
  }
}
