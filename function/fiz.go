/**
 * fn/fiz.go
 *
 * @author  Denis Luchkin-Zhou <wyvernzora@gmail.com>
 * @license MIT
 */

package function
import "github.com/jluchiji/go-fiz/base"


type Fiz struct {
  args  []string
  body  base.AstNode
}


func (this *Fiz) GetArgs() []string {
  return this.args
}


func (this *Fiz) Call(args []base.AstNode, context []int) int {
  /* Check the number of arguments */
  if (len(args) != len(this.args)) {
    panic("FIZ_ARGNUM")
  }
  /* Evaluate all parameters */
  params := make([]int, len(args))
  for i, v := range args {
    params[i] = v.Evaluate(context)
  }
  /* Evaluate function body */
  return this.body.Evaluate(params)
}


func NewFiz(args []string, body base.AstNode) *Fiz {
  f := &Fiz{
    args,
    body,
  }
  body.Resolve(f)
  return f
}
