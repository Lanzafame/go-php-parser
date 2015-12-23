package main

// Tree is the representation of a single parsed file.
type Tree struct {
	Name      string    // name of the file represented by the tree.
	ParseName string    // name of the top-level file during parsing, for error messages.
	Root      *ListNode // top-level root of the tree.
	text      string    // text parsed to create the file.
	// Parsing only; cleared after parse.
	funcs     []map[string]interface{}
	lex       *lexer
	token     [3]item // three-token lookahead for parser.
	peekCount int
	vars      []string // variables defined at the moment.
}
