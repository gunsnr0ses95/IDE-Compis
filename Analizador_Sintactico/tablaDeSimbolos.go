package main

import (
	"strconv"
	"strings"
)

//Variables globales**************************************
var SHIFT = 4
var SIZE = 211
var hashTable [211]*BucketListRec

//********************************************************

type BucketListRec struct {
	token   *Token
	lines   *LineListRec
	valI    int
	valF    float64
	valB    bool
	tipo    string
	haveVal bool
	next    *BucketListRec
	memloc  int
}

type LineListRec struct {
	lineno int
	next   *LineListRec
}

func newLineListRec(lineno int) *LineListRec {
	var llr *LineListRec
	llr = new(LineListRec)
	llr.lineno = lineno
	llr.next = nil
	return llr
}

func newBucketListRec(token *Token, next *BucketListRec, lines *LineListRec, valI int, valF float64, valB bool, tipo string, haveVal bool, memloc int) *BucketListRec {
	var blr *BucketListRec
	blr = new(BucketListRec)
	blr.token = token
	blr.lines = lines
	blr.valI = valI
	blr.valF = valF
	blr.valB = valB
	blr.tipo = tipo
	blr.haveVal = haveVal
	blr.next = next
	blr.memloc = memloc
	return blr
}
func hash(key string) int {
	temp := 0
	for i := 0; i < len(key); i++ {
		temp = ((temp << 4) + int(key[i])) % SIZE
	}
	return temp
}

func st_insert(token *Token, lineno int, valI int, valF float64, valB bool, tipo string, isDec bool, haveVal bool, memloc int) {
	h := hash(token.lexema)
	l := hashTable[h]
	for l != nil && strings.Compare(token.lexema, l.token.lexema) != 0 {
		l = l.next
	}
	if l == nil { //Variable que no esta en la tabla
		list := newLineListRec(lineno)
		l = newBucketListRec(token, hashTable[h], list, valI, valF, valB, tipo, haveVal, memloc)
		hashTable[h] = l
	} else { // esta en la tabla, de modo que solo se agrega el numero de linea
		t := l.lines
		if strings.Compare(l.tipo, "Int") == 0 {
			l.valI = valI
		} else if strings.Compare(l.tipo, "Float") == 0 {
			l.valF = valF
		} else {
			l.valB = valB
		}
		l.haveVal = haveVal
		for t.next != nil {
			t = t.next
		}
		if lineno != 0 {
			t.next = newLineListRec(lineno)
		}
	}
}

func st_lookup(name string) *BucketListRec {
	h := hash(name)
	l := hashTable[h]
	for l != nil && strings.Compare(name, l.token.lexema) != 0 {
		l = l.next
	}
	if l == nil {
		return nil
	} else {
		return l
	}
}
func printSymTab() {
	writerSym.WriteString("Nombre\tLoc. Mem\tTipo\t\tValor\t\tNo Linea\n")
	strWriter := ""
	for i := 0; i < SIZE; i++ {
		if hashTable[i] != nil {
			l := hashTable[i]
			for l != nil {
				t := l.lines
				if len(l.token.lexema) > 2 {
					writerSym.WriteString(l.token.lexema + "\t\t\t\t" + strconv.Itoa(l.memloc) + "\t\t\t" + l.tipo)
				} else {
					writerSym.WriteString(l.token.lexema + "\t\t\t\t\t\t" + strconv.Itoa(l.memloc) + "\t\t\t" + l.tipo)
				}
				if len(l.tipo) < 4 {
					writerSym.WriteString("\t\t\t")
				} else {
					writerSym.WriteString("\t\t")
				}
				switch l.tipo {
				case "Int":
					strWriter = strconv.Itoa(l.valI)
				case "Float":
					strWriter = strconv.FormatFloat(l.valF, 'f', -1, 64)
				case "Bool":
					strWriter = strconv.FormatBool(l.valB)
				}
				writerSym.WriteString(strWriter)
				if len(strWriter) < 3 {
					writerSym.WriteString("\t\t\t\t")
				} else {
					writerSym.WriteString("\t\t")
				}
				for t != nil {
					writerSym.WriteString(strconv.Itoa(t.lineno))
					t = t.next
					if t != nil {
						writerSym.WriteString(", ")
					}
				}
				writerSym.WriteString("\n")
				l = l.next
			}
		}
	}
}
