/**
 * base/func.go
 *
 * @author  Denis Luchkin-Zhou <wyvernzora@gmail.com>
 * @license MIT
 */

package base



/**
 * Base type for all Fiz functions
 */
type Func interface {
  GetArgs() []string
  Call(args []AstNode, context []int) int
}
