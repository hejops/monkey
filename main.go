package main

import "monkey/repl"

// an implementation of the Monkey programming language, as specified by
// Thorsten Ball in "Writing an Interpreter in Go". the language features:
//
// •  C-like syntax
// •  variable bindings
// •  integers and booleans (no other numeric types!)
// •  arithmetic expressions
// •  built-in functions
// •  first-class and higher-order functions
// •  closures
// •  a string data structure
// •  an array data structure
// •  a hash data structure
//
// (no comments?)

// note: functions are defined js/lua style:
// let foo = fn() {...}

func main() {
	repl.Start()
}
