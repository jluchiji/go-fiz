/**
 * fiz.nex
 *
 * 2016 (C) Denis Luchkin-Zhou <wyvernzora@gmail.com>
 * Released under MIT license
 */

%{
  package main
  import (
    "fmt"
    "github.com/jluchiji/go-fiz/base"
    "github.com/jluchiji/go-fiz/ast"
    "github.com/jluchiji/go-fiz/function"
  )
%}


%union {
  str_val   string
  num_val   int
  node_val  base.AstNode
  id_list   []string
  node_list []base.AstNode
}

%token OP
%token CP
%token DEF

%token <str_val>   ID
%token <num_val>   NUM
%token <str_val>   XID

%type  <node_val>  expr
%type  <node_list> exprs
%type  <id_list>   idlist

%%

statements:
  statement | statement statements;

statement:
  OP DEF OP ID idlist CP expr CP  { function.Register($4, function.NewFiz($5, $7))  } |
  expr                            { fmt.Println($1.Evaluate([]int{ }));             } ;

idlist:
  ID idlist                       { $$ = append([]string{ $1 }, $2...)        } |
  /* EMPTY */                     { $$ = []string{ }                          } ;

expr:
  OP ID exprs CP                  {
                                    f, ok := function.Find($2)
                                    if (!ok) {
                                      panic("FIZ_UNDEF_FUNC")
                                    }
                                    $$ = ast.NewCall(f, $3)
                                  } |
  ID                              { $$ = ast.NewVariable($1)                  } |
  XID                             { panic("FIZ_NOT_IMPLEMENTED")              } |
  NUM                             { $$ = ast.NewNumber($1)                    } ;

exprs:
  expr exprs                      { $$ = append([]base.AstNode{ $1 }, $2...)  } |
  /* EMPTY */                     { $$ = []base.AstNode{ }                    } ;

%%
