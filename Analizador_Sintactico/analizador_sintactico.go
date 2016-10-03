// Analizador_Sintactico project main.go
package main

import (
	"fmt"
	"os"
	"strconv"
)

type expType int
type expKind int
type stmtKind int
type nodeKind int

//Variables globales
var token *Token
var tabno = -1
var currentValType expType
var erroresW, _ = os.Create("sintactico_info.txt")
var isDecl = false
var memloc = 0

/*-------------------*/

const (
	INT expType = 1 + iota
	FLOAT
	BOOL
)

const (
	OPK expKind = 1 + iota
	IDK
	CONSTK
)

const (
	PROGRAMA stmtKind = 1 + iota
	SELECCION
	ITERACION
	REPETICION
	READ
	WRITE
	BLOQUE
	ASIGNACION
)

const (
	EXPK nodeKind = 1 + iota
	STMTK
)

type Kind struct {
	stmt stmtKind
	exp  expKind
}
type TreeNode struct {
	hijo            [3]*TreeNode
	hermano         *TreeNode
	nodekind        nodeKind
	token           *Token
	tipo            expType /* for type checking of exps */
	kind            Kind
	varType         token_types
	valInt          int
	valFloat        float64
	valBool         bool
	typeError       bool
	undeclaredError bool
	isIntType       bool
}

func newStmtNode(kind stmtKind) *TreeNode {
	var t *TreeNode
	t = new(TreeNode)
	for i := 0; i < 3; i++ {
		t.hijo[i] = nil
	}
	t.hermano = nil
	t.nodekind = STMTK
	t.kind.stmt = kind
	/*t.valBool = false
	t.typeError = false
	t.undeclaredError = false
	t.isIntType = false*/
	return t
}

func newExpNode(kind expKind) *TreeNode {
	var t *TreeNode
	t = new(TreeNode)
	for i := 0; i < 3; i++ {
		t.hijo[i] = nil
	}
	t.hermano = nil
	t.nodekind = EXPK
	t.kind.exp = kind
	t.tipo = 0
	/*t.valBool = false
	t.typeError = false
	t.undeclaredError = false
	t.isIntType = false*/
	return t
}

func syntaxError(message string) {
	fmt.Printf("\n>>> ")
	erroresW.WriteString("\n>>> ")
	fmt.Printf("Error sintatico en la linea  %d: %s", token.nline, message)
	erroresW.WriteString("Error sintactico en la linea " + strconv.Itoa(token.nline) + ": " + message)
	//Error = true
}

func match(expected token_types) {
	if token.tokenval == expected {
		if token.tokenval == TKN_ID {
			if isDecl {
				tipo := ""
				if currentValType == INT {
					tipo = "Int"
				} else if currentValType == FLOAT {
					tipo = "Float"
				} else {
					tipo = "Bool"
				}
				st_insert(token, token.nline, 0, 0, true, tipo, true, false, memloc)
				memloc++
			} else {
				l := st_lookup(token.lexema)
				if l != nil {
					st_insert(token, token.nline, 0, 0, true, "", true, false, memloc)
					memloc++
				} else {
					writerSymInfo.WriteString("Variable no declarada: " + token.lexema + " No. Linea: " + strconv.Itoa(token.nline) + "\n")
				}
			} // fin if isDecl

		} // fin if TKN_ID
		token = GetToken(reader, writer)
	} else {
		syntaxError("token inesperado -> ")
		fmt.Printf("%s\t%s", GetTknString(token.tokenval), token.lexema)
		erroresW.WriteString(GetTknString(token.tokenval) + " '" + token.lexema + "' se esperaba-> " + GetTknString(expected) + "\n")
		fmt.Printf("\n")
	}
}

func printTabulacion() {
	for i := 0; i < tabno; i++ {
		fmt.Printf("    ")
		writer2.WriteString("    ")
	}
}

func printTree(tree *TreeNode) {
	tabno++
	for tree != nil {
		printTabulacion()
		if tree.nodekind == STMTK {
			switch tree.kind.stmt {
			case PROGRAMA:
				fmt.Printf("Program\n")
				writer2.WriteString("Program\n")
			case SELECCION:
				fmt.Printf("If\n")
				writer2.WriteString("If\n")
			case REPETICION:
				fmt.Printf("Repeat\n")
				writer2.WriteString("Repeat\n")
			case ASIGNACION:
				fmt.Printf("Assign to: %s\n", tree.token.lexema)
				writer2.WriteString("Assign to: " + tree.token.lexema + "\n")
			case ITERACION:
				fmt.Printf("While\n")
				writer2.WriteString("While\n")
			case READ:
				fmt.Printf("Read: %s\n", tree.token.lexema)
				writer2.WriteString("Read: " + tree.token.lexema + "\n")
			case WRITE:
				fmt.Printf("Write\n")
				writer2.WriteString("Write\n")
			}
		} else if tree.nodekind == EXPK {
			switch tree.kind.exp {
			case OPK:
				fmt.Printf("Op: ")
				writer2.WriteString("Op: ")
				fmt.Printf("%s\t%s\n", GetTknString(tree.token.tokenval), tree.token.lexema)
				writer2.WriteString(GetTknString(tree.token.tokenval) + tree.token.lexema + "\n")
			case CONSTK:
				fmt.Printf("Const: %s\n", tree.token.lexema)
				writer2.WriteString("Const: " + tree.token.lexema + "\n")

			case IDK:
				if tree.token == nil {
					goto salir
				}
				fmt.Printf("Id: %s", tree.token.lexema)
				writer2.WriteString("Id: " + tree.token.lexema)
				if tree.tipo != 0 {
					fmt.Printf("-%s\n", tipoToString(tree.tipo))
					writer2.WriteString("-" + tipoToString(tree.tipo) + "\n")
				} else {
					fmt.Println()
					writer2.WriteString("\n")
				}
			}
		} else {
			fmt.Printf("Unknown node kind\n")
		}
		for i := 0; i < 3; i++ {
			printTree(tree.hijo[i])
		}
		tree = tree.hermano
	}
salir:
	tabno--
}

func tipoToString(tipo expType) string {
	tipoStr := ""
	switch tipo {
	case INT:
		tipoStr = "Int"
	case FLOAT:
		tipoStr = "Float"
	case BOOL:
		tipoStr = "Bool"
	}
	return tipoStr
}

//factor → ( expresión ) | numero | identificador
func factor() *TreeNode {
	var t *TreeNode
	t = nil
	switch token.tokenval {
	case TKN_NUM:
		t = newExpNode(CONSTK)
		if (t != nil) && (token.tokenval == TKN_NUM) {
			t.token = token
			t.valInt, _ = strconv.Atoi(token.lexema)
			t.isIntType = true
		}
		match(TKN_NUM)

	case TKN_NUM_FLOAT:
		t = newExpNode(CONSTK)
		if (t != nil) && (token.tokenval == TKN_NUM_FLOAT) {
			t.token = token
			t.valFloat, _ = strconv.ParseFloat(token.lexema, 64)
			t.isIntType = false
		}
		match(TKN_NUM_FLOAT)

	case TKN_ID:
		t = newExpNode(IDK)
		if (t != nil) && (token.tokenval == TKN_ID) {
			t.token = token
		}
		match(TKN_ID)
	case TKN_TRUE:
		t = newExpNode(IDK)
		if t != nil && token.tokenval == TKN_TRUE {
			t.token = token
		}
		match(TKN_TRUE)

	case TKN_FALSE:
		t = newExpNode(IDK)
		if t != nil && token.tokenval == TKN_FALSE {
			t.token = token
		}
		match(TKN_FALSE)

	case TKN_LPARENT:
		match(TKN_LPARENT)
		t = exp()
		match(TKN_RPARENT)

	default:
		syntaxError("token inesperado -> ")

		token = GetToken(reader, writer)

	}
	return t
}

func exp() *TreeNode {
	var t *TreeNode
	t = term_and()
	for token.tokenval == TKN_OR {
		var p *TreeNode
		p = newExpNode(OPK)
		if p != nil {
			p.hijo[0] = t
			p.token = token
			t = p
		}
		match(TKN_OR)
		if t != nil {
			t.hijo[1] = term_and()
		}
	}
	return t
}

func term_and() *TreeNode {
	var t *TreeNode
	t = igualdad()
	for token.tokenval == TKN_AND {
		var p *TreeNode
		p = newExpNode(OPK)
		if p != nil {
			p.hijo[0] = t
			p.token = token
			t = p
		}
		match(TKN_AND)
		if t != nil {
			t.hijo[1] = igualdad()
		}
	}
	return t
}

func igualdad() *TreeNode {
	var t *TreeNode
	t = relacion()
	for token.tokenval == TKN_EQUAL || token.tokenval == TKN_NEQUAL {
		var p *TreeNode
		p = newExpNode(OPK)
		if p != nil {
			p.hijo[0] = t
			p.token = token
			t = p
		}
		match(token.tokenval)
		if t != nil {
			t.hijo[1] = relacion()
		}
	}
	return t
}

func relacion() *TreeNode {
	var t *TreeNode
	t = exp_suma()
	for token.tokenval == TKN_LETHAN || token.tokenval == TKN_LTHAN || token.tokenval == TKN_GETHAN || token.tokenval == TKN_GTHAN {
		var p *TreeNode
		p = newExpNode(OPK)
		if p != nil {
			p.hijo[0] = t
			p.token = token
			t = p
		}
		match(token.tokenval)
		if t != nil {
			t.hijo[1] = exp_suma()
		}
	}
	return t
}

func exp_suma() *TreeNode {
	var t *TreeNode
	t = termino()
	for token.tokenval == TKN_ADD || token.tokenval == TKN_MINUS {
		var p *TreeNode
		p = newExpNode(OPK)
		if p != nil {
			p.hijo[0] = t
			p.token = token
			t = p
		}
		match(token.tokenval)
		if t != nil {
			t.hijo[1] = termino()
		}
	}
	return t
}

func termino() *TreeNode {
	var t *TreeNode
	t = unario()
	for token.tokenval == TKN_DIVISION || token.tokenval == TKN_PRODUCT {
		var p *TreeNode
		p = newExpNode(OPK)
		if p != nil {
			p.hijo[0] = t
			p.token = token
			t = p

		}
		match(token.tokenval)
		if t != nil {
			t.hijo[1] = unario()
		}
	}
	return t
}

func unario() *TreeNode {
	var t *TreeNode
	t = nil
	switch token.tokenval {
	case TKN_MINUS:
	case TKN_ADD:
		t = newExpNode(OPK)
		t.token = token
		match(token.tokenval)
		t.hijo[0] = factor()
	case TKN_NOT:
		t = newExpNode(OPK)
		t.token = token
		match(TKN_NOT)
		t.hijo[0] = factor()
	default:
		t = factor()
	}
	return t
}

func program() *TreeNode {
	var t *TreeNode
	t = newStmtNode(PROGRAMA)
	match(TKN_PROGRAM)
	match(TKN_LBRACE)
	if token.tokenval != TKN_EOF {
		t.hijo[0] = lista_declaracion()
		if t.hijo[0] == nil {
			t.hijo[0] = lista_sentencias()
		} else {
			t.hijo[1] = lista_sentencias()
		}

	}
	match(TKN_RBRACE)
	return t
}

func lista_declaracion() *TreeNode {
	var t *TreeNode
	t = nil
	for token.tokenval != TKN_WRITE && token.tokenval != TKN_READ && token.tokenval != TKN_IF && token.tokenval != TKN_DO && token.tokenval != TKN_WHILE && token.tokenval != TKN_LBRACE && token.tokenval != TKN_RBRACE && token.tokenval != TKN_ID {
		t = declaracion()
		match(TKN_SEMICOLON)
		t.hermano = lista_declaracion()
	}
	return t
}

func declaracion() *TreeNode {
	var t *TreeNode
	t = nil
	isDecl = true
	if token.tokenval == TKN_INT || token.tokenval == TKN_FLOAT || token.tokenval == TKN_BOOL {
		switch token.tokenval {
		case TKN_INT:
			currentValType = INT
		case TKN_FLOAT:
			currentValType = FLOAT
		case TKN_BOOL:
			currentValType = BOOL
		}
		match(token.tokenval)
	}
	t = lista_id()
	isDecl = false
	return t
}

func lista_id() *TreeNode {
	var t *TreeNode
	t = newExpNode(IDK)
	t.tipo = currentValType
	if (t != nil) && (token.tokenval != TKN_ID) {
		syntaxError("Declaración erronea")
		for (token.tokenval != TKN_COMMA) && (token.tokenval != TKN_SEMICOLON) {
			token = GetToken(reader, writer)
		}
	}
	if (t != nil) && (token.tokenval == TKN_ID) {
		t.token = token
		match(TKN_ID)
	}
	if token.tokenval == TKN_COMMA {
		match(TKN_COMMA)
		t.hijo[0] = lista_id()
	}
	return t
}

func lista_sentencias() *TreeNode {
	var t *TreeNode
	t = nil
	for token.tokenval != TKN_RBRACE && token.tokenval != TKN_EOF && token.tokenval != TKN_UNTIL && token.tokenval != TKN_ELSE && token.tokenval != TKN_FI {
		t = sentencia()
		if token.tokenval == TKN_IF || token.tokenval == TKN_WRITE || token.tokenval == TKN_LBRACE || token.tokenval == TKN_DO || token.tokenval == TKN_ID || token.tokenval == TKN_READ || token.tokenval == TKN_WHILE {
			if t != nil {
				t.hermano = lista_sentencias()
			}
		}
	}
	return t
}

func sentencia() *TreeNode {
	var t *TreeNode
	t = nil
	switch token.tokenval {
	case TKN_IF:
		t = seleccion()
	case TKN_ID:
		t = asignacion()
	case TKN_READ:
		t = sent_read()
	case TKN_WRITE:
		t = sent_write()
	case TKN_DO:
		t = repeticion()
	case TKN_WHILE:
		t = iteracion()
	case TKN_LBRACE:
		t = bloque()
	default:
		syntaxError("Token inesperado->")
		token = GetToken(reader, writer)
	}
	return t
}

func seleccion() *TreeNode {
	var t *TreeNode
	t = newStmtNode(SELECCION)
	match(TKN_IF)
	match(TKN_LPARENT)
	if t != nil {
		t.hijo[0] = exp()
	}
	match(TKN_RPARENT)
	if t != nil {
		t.hijo[1] = bloque()
	}
	if token.tokenval == TKN_ELSE {
		match(TKN_ELSE)
		if t != nil {
			t.hijo[2] = bloque()
		}
	}
	match(TKN_FI)
	return t
}

func asignacion() *TreeNode {
	var t *TreeNode
	t = newStmtNode(ASIGNACION)
	if (t != nil) && (token.tokenval == TKN_ID) {
		t.token = token
	}
	match(TKN_ID)
	match(TKN_ASSIGN)
	if t != nil {
		t.hijo[0] = exp()
	}
	match(TKN_SEMICOLON)
	return t
}

func sent_read() *TreeNode {
	var t *TreeNode
	t = newStmtNode(READ)
	match(TKN_READ)
	if (t != nil) && (token.tokenval == TKN_ID) {
		t.token = token
	}
	match(TKN_ID)
	match(TKN_SEMICOLON)
	return t
}

func sent_write() *TreeNode {
	var t *TreeNode
	t = newStmtNode(WRITE)
	match(TKN_WRITE)
	if t != nil {
		t.hijo[0] = exp()
	}
	match(TKN_SEMICOLON)
	return t
}

func repeticion() *TreeNode {
	var t *TreeNode
	t = newStmtNode(REPETICION)
	match(TKN_DO)
	if t != nil {
		t.hijo[0] = bloque()
	}
	match(TKN_UNTIL)
	match(TKN_LPARENT)
	if t != nil {
		t.hijo[1] = exp()
	}
	match(TKN_RPARENT)
	match(TKN_SEMICOLON)
	return t
}

func iteracion() *TreeNode {
	var t *TreeNode
	t = newStmtNode(ITERACION)
	match(TKN_WHILE)
	match(TKN_LPARENT)
	if t != nil {
		t.hijo[0] = exp()
	}
	match(TKN_RPARENT)
	if t != nil {
		t.hijo[1] = bloque()
	}
	return t
}

func bloque() *TreeNode {
	var t *TreeNode
	t = newStmtNode(BLOQUE)
	match(TKN_LBRACE)
	if t != nil {
		t.hijo[0] = lista_sentencias()
	}
	match(TKN_RBRACE)
	return t
}

func sintactico() *TreeNode {
	var t *TreeNode
	token = GetToken(reader, writer)
	t = program()
	if token.tokenval != TKN_EOF {
		syntaxError("Code ends before file\n")
	}
	return t
}
