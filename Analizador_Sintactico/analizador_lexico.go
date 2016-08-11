// Analizador_Lexico project main.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type State int
type token_types int

const (
	//Palabras reservadas
	TKN_PROGRAM token_types = 1 + iota
	TKN_IF
	TKN_ELSE
	TKN_FI
	TKN_DO
	TKN_UNTIL
	TKN_WHILE
	TKN_READ
	TKN_WRITE
	TKN_FLOAT
	TKN_INT
	TKN_BOOL
	TKN_NOT
	TKN_AND
	TKN_OR
	TKN_TRUE
	TKN_FALSE

	//Simbolos especiales
	TKN_ADD
	TKN_MINUS
	TKN_PRODUCT
	TKN_DIVISION
	TKN_EXPONENTIAL
	TKN_LTHAN
	TKN_LETHAN
	TKN_GTHAN
	TKN_GETHAN
	TKN_EQUAL
	TKN_NEQUAL
	TKN_ASSIGN
	TKN_SEMICOLON
	TKN_COMMA
	TKN_LPARENT
	TKN_RPARENT
	TKN_LBRACE
	TKN_RBRACE
	TKN_COMMENT
	TKN_MLCOMMENT

	//Identificadores y numeros

	TKN_ID
	TKN_NUM

	//Error
	TKN_ERROR

	//Fin de archivo

	TKN_EOF
)

const (
	IN_START State = 1 + iota
	IN_ID
	IN_NUM
	IN_LPARENT
	IN_RPARENT
	IN_SEMICOLON
	IN_COMMA
	IN_EQU
	IN_NEQU
	IN_ADD
	IN_MINUS
	IN_EOF
	IN_ERROR
	IN_DONE
	IN_LESS
	IN_GR
	IN_COMMENT_OR_DIVISION
	IN_COMMENT
	IN_MLCOMMENT
	IN_MLCOMMENTERROR
	IN_END_OF_MLCOMMENT
	IN_NUM_OR_OPERATOR
	IN_DEC_POIN
	IN_EXPONENTIAL
)

type Token struct {
	tokenval token_types
	lexema   string
	nline    int
}

var ReserveWords = []Token{
	Token{TKN_IF, "if", 0},
	Token{TKN_NOT, "not", 0},
	Token{TKN_AND, "and", 0},
	Token{TKN_OR, "or", 0},
	Token{TKN_ELSE, "else", 0},
	Token{TKN_FI, "fi", 0},
	Token{TKN_DO, "do", 0},
	Token{TKN_UNTIL, "until", 0},
	Token{TKN_WHILE, "while", 0},
	Token{TKN_READ, "read", 0},
	Token{TKN_WRITE, "write", 0},
	Token{TKN_FLOAT, "float", 0},
	Token{TKN_INT, "int", 0},
	Token{TKN_BOOL, "bool", 0},
	Token{TKN_PROGRAM, "program", 0},
	Token{TKN_TRUE, "true", 0},
	Token{TKN_FALSE, "false", 0},
}

var nline = 0
var ncol = 0
var n = 0
var decimal_point_flag = false
var writer, _ = os.Create("tokens_output.txt")
var writerE, _ = os.Create("tokens_info.txt")
var writer2, _ = os.Create("sintactico_output.txt")
var buffer = ""
var reader *bufio.Reader
var tokens = false

func LookUpReservedWords(tok *Token, s string) {
	for i := 0; i < len(ReserveWords); i++ {
		if strings.Compare(s, ReserveWords[i].lexema) == 0 {
			tok.tokenval = ReserveWords[i].tokenval
			tok.lexema = ReserveWords[i].lexema
			goto EndFunction
		}
	}
	tok.lexema = s
	tok.tokenval = TKN_ID
EndFunction:
}
func GetChar(reader *bufio.Reader) rune {
	var c rune
	if !(ncol < n) {
		linea, _, err := reader.ReadLine()
		if err == nil {
			buffer = string(linea)
			n = len(buffer)
			ncol = 0
			nline++
			for n == 0 {
				linea, _, err = reader.ReadLine()
				buffer = string(linea)
				n = len(buffer)
				ncol = 0
				nline++
			}

		} else {
			return '$' //End of file
		}
	}
	c = rune(buffer[ncol])
	ncol++
	return c
}
func UnGetChar() {
	ncol--
}
func isDelimit(c rune) bool {
	if c == ' ' || c == '\t' || c == '\n' {
		return true
	}
	return false
}

func GetToken(readerFile *bufio.Reader, writer *os.File) *Token {
	c := ' '
	state := IN_START
	var token = new(Token)
	for state != IN_DONE {
		switch state { //Selection of state
		case IN_START:
			{
				c = GetChar(readerFile)
				for isDelimit(c) { //While the character is a delimiter
					c = GetChar(readerFile)
				}
				if unicode.IsLetter(c) {
					state = IN_ID
					token.lexema += string(c)
				} else if unicode.IsDigit(c) {
					state = IN_NUM
					token.lexema += string(c)
				} else if c == '(' {
					token.tokenval = TKN_LPARENT
					state = IN_DONE
					token.lexema += string(c)
				} else if c == ')' {
					token.tokenval = TKN_RPARENT
					state = IN_DONE
					token.lexema += string(c)
				} else if c == '}' {
					token.tokenval = TKN_RBRACE
					state = IN_DONE
					token.lexema += string(c)
				} else if c == '{' {
					token.tokenval = TKN_LBRACE
					state = IN_DONE
					token.lexema += string(c)
				} else if c == ';' {
					token.tokenval = TKN_SEMICOLON
					state = IN_DONE
					token.lexema += string(c)
				} else if c == ',' {
					token.tokenval = TKN_COMMA
					state = IN_DONE
					token.lexema += string(c)
				} else if c == '=' {
					state = IN_EQU
					token.tokenval = TKN_ASSIGN
					token.lexema += string(c)
				} else if c == '!' {
					state = IN_NEQU
					token.lexema += string(c)
				} else if c == '+' {
					token.tokenval = TKN_ADD
					state = IN_NUM_OR_OPERATOR
					token.lexema += string(c)
				} else if c == '-' {
					token.tokenval = TKN_MINUS
					state = IN_NUM_OR_OPERATOR
					token.lexema += string(c)
				} else if c == '*' {
					token.tokenval = TKN_PRODUCT
					state = IN_DONE
					token.lexema += string(c)
				} else if c == '/' {
					token.tokenval = TKN_DIVISION
					state = IN_COMMENT_OR_DIVISION
					token.lexema += string(c)
				} else if c == '^' {
					token.tokenval = TKN_EXPONENTIAL
					state = IN_DONE
					token.lexema += string(c)
				} else if c == '<' {
					token.tokenval = TKN_LTHAN
					state = IN_LESS
					token.lexema += string(c)
				} else if c == '>' {
					token.tokenval = TKN_GTHAN
					state = IN_GR
					token.lexema += string(c)
				} else if c == '$' {
					token.tokenval = TKN_EOF
					state = IN_DONE
					token.lexema += string(c)
				} else {
					token.tokenval = TKN_ERROR
					state = IN_ERROR
				}
			}
		case IN_COMMENT_OR_DIVISION:
			{
				c = GetChar(readerFile)
				if c == '/' {
					token.tokenval = TKN_COMMENT
					state = IN_COMMENT
					token.lexema += string(c)
				} else if c == '*' {
					token.tokenval = TKN_MLCOMMENT
					state = IN_MLCOMMENT
					token.lexema += string(c)
				} else {
					state = IN_DONE
					UnGetChar()
				}
			}
		case IN_COMMENT:
			{
				token.lexema = ""
				c = GetChar(readerFile)
				if n == ncol {
					//token.lexema += string(c)
					state = IN_START
					break
				}
				//token.lexema += string(c)
			}
		case IN_MLCOMMENT:
			{
				c = GetChar(readerFile)
				if c == '$' {
					state = IN_MLCOMMENTERROR
				}
				if c == '*' {
					state = IN_END_OF_MLCOMMENT
				}
				token.lexema += string(c)
			}
		case IN_END_OF_MLCOMMENT:
			{
				c = GetChar(readerFile)
				if c == '$' {
					state = IN_MLCOMMENTERROR
				}
				if c == '/' {
					token.lexema = ""
					state = IN_START
				} else {
					state = IN_MLCOMMENT
				}
				//token.lexema += string(c)
			}
		case IN_MLCOMMENTERROR:
			{
				writer.WriteString("The multiline comment has not been closed\n")
				state = IN_DONE
			}
		case IN_NUM_OR_OPERATOR:
			{
				c = GetChar(readerFile)
				previous_char_to_symbol := buffer[ncol-3]
				if unicode.IsLetter(rune(previous_char_to_symbol)) || unicode.IsDigit(rune(previous_char_to_symbol)) {
					UnGetChar()
					state = IN_DONE
				} else {
					token.lexema += string(c)
					state = IN_NUM
				}
			}
		case IN_NUM:
			{
				c = GetChar(readerFile)
				token.lexema += string(c)
				if !unicode.IsDigit(c) {
					if c == '.' && !decimal_point_flag {
						decimal_point_flag = true
						break
					}
					decimal_point_flag = false
					token.tokenval = TKN_NUM
					state = IN_DONE
					token.lexema = string(token.lexema[0 : len(token.lexema)-1])
					UnGetChar()
				}
			}
		case IN_LESS:
			{
				c = GetChar(readerFile)
				if c == '=' { //pudiera ser el operador <=
					token.lexema += string(c)
					token.tokenval = TKN_LETHAN
				} else { //o solo ser <
					UnGetChar()
				}
				state = IN_DONE
				break
			}
		case IN_GR:
			{
				c = GetChar(readerFile)
				if c == '=' { //pudiera ser el operador >=
					token.lexema += string(c)
					token.tokenval = TKN_GETHAN
				} else { //o solo ser >
					UnGetChar()
				}
				state = IN_DONE
				break
			}
		case IN_NEQU:
			{
				c = GetChar(readerFile)
				if c == '=' {
					token.lexema += string(c)
					token.tokenval = TKN_NEQUAL
				}
				state = IN_DONE
			}
		case IN_EQU:
			{
				c = GetChar(readerFile)
				if c == '=' {
					token.lexema += string(c)
					token.tokenval = TKN_EQUAL
				} else {
					UnGetChar()
				}
				state = IN_DONE
			}
		case IN_ID:
			{
				c = GetChar(readerFile)
				token.lexema += string(c)
				if !(unicode.IsLetter(c) || unicode.IsDigit(c)) {
					token.tokenval = TKN_ID
					state = IN_DONE
					UnGetChar()
					token.lexema = string(token.lexema[0 : len(token.lexema)-1])
					LookUpReservedWords(token, token.lexema)
				}
				break
			}
		default:
			{
				token.tokenval = TKN_ERROR
				state = IN_DONE
				token.lexema += string(c)
			}
		} //end switch
		token.nline = nline
	} //end while

	if token.tokenval == TKN_ERROR && !tokens {
		writerE.WriteString("Line:" + strconv.Itoa(nline) + "," + strconv.Itoa(ncol) + " Error:Character " + string(c) + " does not match any token.\n")
	}
	return token
}

func GetTknString(tkn token_types) string {
	switch tkn {
	case TKN_PROGRAM:
		return "TKN_PROGRAM"
	case TKN_IF:
		return "TKN_IF"
	case TKN_ELSE:
		return "TKN_ELSE"
	case TKN_FI:
		return "TKN_FI"
	case TKN_DO:
		return "TKN_DO"
	case TKN_UNTIL:
		return "TKN_UNTIL"
	case TKN_WHILE:
		return "TKN_WHILE"
	case TKN_READ:
		return "TKN_READ"
	case TKN_WRITE:
		return "TKN_WRITE"
	case TKN_FLOAT:
		return "TKN_FLOAT"
	case TKN_INT:
		return "TKN_INT"
	case TKN_BOOL:
		return "TKN_BOOL"
	case TKN_NOT:
		return "TKN_NOT"
	case TKN_AND:
		return "TKN_AND"
	case TKN_OR:
		return "TKN_OR"
	case TKN_ADD:
		return "TKN_ADD"
	case TKN_MINUS:
		return "TKN_MINUS"
	case TKN_PRODUCT:
		return "TKN_PRODUCT"
	case TKN_DIVISION:
		return "TKN_DIVISION"
	case TKN_EXPONENTIAL:
		return "TKN_EXPONENTIAL"
	case TKN_LTHAN:
		return "TKN_LTHAN"
	case TKN_LETHAN:
		return "TKN_LETHAN"
	case TKN_GTHAN:
		return "TKN_GTHAN"
	case TKN_GETHAN:
		return "TKN_GETHAN"
	case TKN_EQUAL:
		return "TKN_EQUAL"
	case TKN_NEQUAL:
		return "TKN_NEQUAL"
	case TKN_ASSIGN:
		return "TKN_ASSIGN"
	case TKN_SEMICOLON:
		return "TKN_SEMICOLON"
	case TKN_COMMA:
		return "TKN_COMMA"
	case TKN_LPARENT:
		return "TKN_LPARENT"
	case TKN_RPARENT:
		return "TKN_RPARENT"
	case TKN_LBRACE:
		return "TKN_LBRACE"
	case TKN_RBRACE:
		return "TKN_RBRACE"
	case TKN_COMMENT:
		return "TKN_COMMENT"
	case TKN_MLCOMMENT:
		return "TKN_MLCOMMENT"
	case TKN_ID:
		return "TKN_ID"
	case TKN_TRUE:
		return "TKN_TRUE"
	case TKN_FALSE:
		return "TKN_FALSE"
	case TKN_NUM:
		return "TKN_NUM"
	case TKN_ERROR:
		return "TKN_ERROR"
	case TKN_EOF:
		return "TKN_EOF"
	}
	return ""
}

func lexico() {
	var file *os.File
	if len(os.Args) == 2 { // ingreso los argumentos correctamente
		file, _ = os.Open(os.Args[1])
		reader = bufio.NewReader(file)
		var token *Token
		token = GetToken(reader, writer)
		for TKN_EOF != token.tokenval {
			writer.WriteString(token.lexema + "\t\t" + strconv.Itoa(token.nline) + "\t\t" + GetTknString(token.tokenval) + "\r\n")
			token = GetToken(reader, writer)
		}
		writer.WriteString(token.lexema + "\t\t" + strconv.Itoa(token.nline) + "\t\t" + GetTknString(token.tokenval) + "\r\n")

	} else {
		fmt.Println("Error, ingresar <programa> <archivo>")
	}
	writer.Close()
	file.Close()
	tokens = true
}

func reset() {
	var file *os.File
	if len(os.Args) == 2 { // ingreso los argumentos correctamente
		file, _ = os.Open(os.Args[1])
		reader = bufio.NewReader(file)
	}
	writer, _ = os.Create("tokens_output2.txt")
	nline = 0
	ncol = 0
	n = 0
}

func main() {
	lexico()
	if len(os.Args) == 2 {
		reset()
		var arbol *TreeNode
		arbol = sintactico()
		fmt.Println("Error aqui")
		printTree(arbol)
	}
}
