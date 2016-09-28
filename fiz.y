/**
 * fiz.nex
 *
 * 2016 (C) Denis Luchkin-Zhou <wyvernzora@gmail.com>
 * Released under MIT license
 */

%{
  package main
  import "fmt"
%}


%union {
  str_val string
  num_val int
  id_list []string
}

%token OP
%token CP
%token DEF

%token <str_val>   ID
%token <num_val>   NUM
%token <str_val>   XID

%type  <id_list>   idlist

%%

statements:
  statement | statement statements;

statement:
  OP DEF OP ID idlist CP expr CP  { fmt.Println("statement 1");               } |
  expr                            { fmt.Println("statement 2");               } ;

idlist:
  ID idlist                       { fmt.Println("idlist 1")                   } |
  /* EMPTY */                     { $$ = []string{ }                          } ;

expr:
  OP ID exprs CP                  { fmt.Println("expr 1");                    } |
  ID                              { fmt.Println("expr 2");                    } |
  XID                             { fmt.Println("expr 3");                    } |
  NUM                             { fmt.Println("expr 4");                    } ;

exprs:
  expr exprs                      { fmt.Println("exprs 1");                   } |
  /* EMPTY */                     { fmt.Println("exprs 2");                   } ;

%%
