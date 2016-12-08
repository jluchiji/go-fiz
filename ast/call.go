/**
 * ast/number.go
 *
 * @author  Denis Luchkin-Zhou <wyvernzora@gmail.com>
 * @license MIT
 */

package ast
import (
  "github.com/jluchiji/go-fiz/base"
  "github.com/jluchiji/go-fiz/function"
)


type Call struct {
  name  string
  argv  []base.AstNode
}


func (this *Call) Resolve(fn base.Func) {
  for _, v := range this.argv {
    v.Resolve(fn)
  }
}


func (this *Call) Evaluate(context []int) int {
  fn, ok := function.Find(this.name)
  if (!ok) {
    panic("FIZ_UNDEF_FUNC")
  }
  return fn.Call(this.argv, context)
}


func NewCall(name string, argv []base.AstNode) *Call {
  return &Call{
    name,
    argv,
  }
}
