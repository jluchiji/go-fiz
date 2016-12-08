/define/                       { return DEF }
/[$][_a-zA-Z][_a-zA-Z0-9]*/    {
                                 lval.str_val = yylex.Text()
                                 return XID;
                               }
/[_a-zA-Z][_a-zA-Z0-9]*/       {
                                 lval.str_val = yylex.Text()
                                 return ID;
                               }
/\(/                           { return OP; }
/\)/                           { return CP; }
/0|[1-9][0-9]*/                {
                                 lval.num_val, _ = strconv.Atoi(yylex.Text())
                                 return NUM;
                               }
/[ \t\n]/                      { /* Discard */ }
/;.*/                          { /* Discard */ }
//

package main
import (
  "os"
  "strconv"
  "github.com/jluchiji/go-fiz/function"
)

func main() {
  function.Init()
  yyParse(NewLexer(os.Stdin))
}
