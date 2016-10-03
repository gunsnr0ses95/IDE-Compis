// analizador_semantico
package main

//variable,_ := strconv.ParseFloat("3.2",64)
import (
	"strconv"
	"strings"
)

//variables Globales**************************************
var inBlock = false
var tipoActual = ""

//********************************************************

func semantico(tree *TreeNode) {
	if tree != nil {
		switch tree.nodekind {
		case STMTK:
			inStmt(tree)
		case EXPK:
			inExp(tree)
		} // fin switch
		semantico(tree.hermano)
	}
}

func inStmt(tree *TreeNode) {
	var p1, p2 *TreeNode
	var l *BucketListRec
	switch tree.kind.stmt {
	case SELECCION:
		semantico(tree.hijo[0])
	case REPETICION:
		p1 = tree.hijo[0]
		p2 = tree.hijo[1]
		semantico(p1)
		semantico(p2)
	case ITERACION:
		p1 = tree.hijo[0]
		p2 = tree.hijo[1]
		semantico(p1)
		semantico(p2)
	case BLOQUE:
		inBlock = true
		semantico(tree.hijo[0])
		inBlock = false
	case ASIGNACION:
		if !inBlock {
			l = st_lookup(tree.token.lexema)
			tipoActual = l.tipo
			semantico(tree.hijo[0])
			//Linea no copiada
			if !tree.hijo[0].typeError || tree.hijo[0].undeclaredError {
				if (tree.hijo[0].isIntType && (strings.Compare(l.tipo, "Int") == 0)) || (!tree.hijo[0].isIntType && strings.Compare(l.tipo, "Float") == 0) {
					st_insert(tree.token, tree.token.nline, tree.hijo[0].valInt, tree.hijo[0].valFloat, tree.hijo[0].valBool, l.tipo, false, true, memloc)
					memloc++
					tree.valInt = tree.hijo[0].valInt
					tree.valFloat = tree.hijo[0].valFloat
				} else {
					writerSymInfo.WriteString("Error: Tipos diferentes. Variables " + tree.token.lexema + " int=" + strconv.Itoa(l.valI) + " float=" + strconv.FormatFloat(tree.hijo[0].valFloat, 'f', -1, 64) + "\n")
					tree.valInt = l.valI
					tree.valFloat = l.valF

				}
			}
		}
	case READ:
		l = st_lookup(tree.token.lexema)
		if l != nil {
			l.haveVal = false
			l.valF = 0.0
			l.valI = 0
		} else {
			writerSymInfo.WriteString("Variable no declarada:" + tree.token.lexema + " No. Linea: " + strconv.Itoa(tree.token.nline) + "\n")
		}
	}
}
func inExp(tree *TreeNode) {
	var p1, p2 *TreeNode
	var l *BucketListRec
	switch tree.kind.exp {
	case IDK:
		l = st_lookup(tree.token.lexema)
		if strings.Compare(l.tipo, "Int") == 0 {
			tree.isIntType = true
		} else {
			tree.isIntType = false
		}

		if l != nil {
			if l.haveVal {
				//tieneVal = true ** Variable no puesta
				if tree.isIntType {
					tree.valInt = l.valI

				} else {
					tree.valFloat = l.valF
				}
			} // tieneval removido
		} else {
			tree.undeclaredError = true
			writerSymInfo.WriteString("Variable no declarada:" + tree.token.lexema + " No. Linea: " + strconv.Itoa(tree.token.nline) + "\n")
		} /* IdK */
	case OPK:
		p1 = tree.hijo[0]
		p2 = tree.hijo[1]
		semantico(p1)
		semantico(p2)
		if !p1.typeError || !p1.undeclaredError || !p1.typeError || !p1.undeclaredError {
			if !(p1.isIntType && p2.isIntType) {
				// tree.typeError = true
				//Console.WriteLine("Tipos diferentes")
				tree.isIntType = false
				switch tree.token.tokenval {
				case TKN_ADD:
					switch p1.isIntType {
					case true:
						tree.valFloat = float64(tree.hijo[0].valInt) + tree.hijo[1].valFloat

					case false:
						tree.valFloat = tree.hijo[0].valFloat + float64(tree.hijo[1].valInt)

					}

				case TKN_MINUS:
					switch p1.isIntType {
					case true:
						tree.valFloat = float64(tree.hijo[0].valInt) - tree.hijo[1].valFloat

					case false:
						tree.valFloat = tree.hijo[0].valFloat - float64(tree.hijo[1].valInt)

					}

				case TKN_PRODUCT:
					switch p1.isIntType {
					case true:
						tree.valFloat = float64(tree.hijo[0].valInt) * tree.hijo[1].valFloat

					case false:
						tree.valFloat = tree.hijo[0].valFloat * float64(tree.hijo[1].valInt)

					}

				case TKN_DIVISION:
					switch p1.isIntType {
					case true:
						tree.valFloat = float64(tree.hijo[0].valInt) / tree.hijo[1].valFloat

					case false:
						tree.valFloat = tree.hijo[0].valFloat / float64(tree.hijo[1].valInt)

					}

				case TKN_LTHAN:
					switch p1.isIntType {
					case true:
						if float64(tree.hijo[0].valInt) < tree.hijo[1].valFloat {
							tree.valBool = true
						} else {
							tree.valBool = false
						}

					case false:
						if tree.hijo[0].valFloat < float64(tree.hijo[1].valInt) {
							tree.valBool = true
						} else {
							tree.valBool = false
						}

					}

				case TKN_LETHAN:
					switch p1.isIntType {
					case true:
						if float64(tree.hijo[0].valInt) <= tree.hijo[1].valFloat {
							tree.valBool = true
						} else {
							tree.valBool = false
						}

					case false:
						if tree.hijo[0].valFloat <= float64(tree.hijo[1].valInt) {
							tree.valBool = true
						} else {
							tree.valBool = false
						}

					}

				case TKN_GTHAN:
					switch p1.isIntType {
					case true:
						if float64(tree.hijo[0].valInt) > tree.hijo[1].valFloat {
							tree.valBool = true
						} else {
							tree.valBool = false
						}

					case false:
						if tree.hijo[0].valFloat > float64(tree.hijo[1].valInt) {
							tree.valBool = true
						} else {
							tree.valBool = false
						}

					}

				case TKN_GETHAN:
					switch p1.isIntType {
					case true:
						if float64(tree.hijo[0].valInt) >= tree.hijo[1].valFloat {
							tree.valBool = true
						} else {
							tree.valBool = false
						}

					case false:
						if tree.hijo[0].valFloat >= float64(tree.hijo[1].valInt) {
							tree.valBool = true
						} else {
							tree.valBool = false
						}

					}

				case TKN_EQUAL:
					switch p1.isIntType {
					case true:
						if float64(tree.hijo[0].valInt) == tree.hijo[1].valFloat {
							tree.valBool = true
						} else {
							tree.valBool = false
						}

					case false:
						if tree.hijo[0].valFloat == float64(tree.hijo[1].valInt) {
							tree.valBool = true
						} else {
							tree.valBool = false
						}

					}

				case TKN_NEQUAL:
					switch p1.isIntType {
					case true:
						if float64(tree.hijo[0].valInt) != tree.hijo[1].valFloat {
							tree.valBool = true
						} else {
							tree.valBool = false
						}

					case false:
						if tree.hijo[0].valFloat != float64(tree.hijo[1].valInt) {
							tree.valBool = true
						} else {
							tree.valBool = false
						}
					}

				} /* case op */
			} else {
				switch tree.token.tokenval {
				case TKN_ADD:
					switch p1.isIntType {
					case true:
						tree.valInt = tree.hijo[0].valInt + tree.hijo[1].valInt

					case false:
						tree.valFloat = tree.hijo[0].valFloat + tree.hijo[1].valFloat

					}

				case TKN_MINUS:
					switch p1.isIntType {
					case true:
						tree.valInt = tree.hijo[0].valInt - tree.hijo[1].valInt

					case false:
						tree.valFloat = tree.hijo[0].valFloat - tree.hijo[1].valFloat

					}

				case TKN_PRODUCT:
					switch p1.isIntType {
					case true:
						tree.valInt = tree.hijo[0].valInt * tree.hijo[1].valInt

					case false:
						tree.valFloat = tree.hijo[0].valFloat * tree.hijo[1].valFloat

					}

				case TKN_DIVISION:
					switch p1.isIntType {
					case true:
						tree.valInt = tree.hijo[0].valInt / tree.hijo[1].valInt

					case false:
						tree.valFloat = tree.hijo[0].valFloat / tree.hijo[1].valFloat

					}

				case TKN_LTHAN:
					switch p1.isIntType {
					case true:
						if tree.hijo[0].valInt < tree.hijo[1].valInt {
							tree.valBool = true
						} else {
							tree.valBool = false
						}

					case false:
						if tree.hijo[0].valFloat < tree.hijo[1].valFloat {
							tree.valBool = true
						} else {
							tree.valBool = false
						}

					}

				case TKN_LETHAN:
					switch p1.isIntType {
					case true:
						if tree.hijo[0].valInt <= tree.hijo[1].valInt {
							tree.valBool = true
						} else {
							tree.valBool = false
						}

					case false:
						if tree.hijo[0].valFloat <= tree.hijo[1].valFloat {
							tree.valBool = true
						} else {
							tree.valBool = false
						}

					}

				case TKN_GTHAN:
					switch p1.isIntType {
					case true:
						if tree.hijo[0].valInt > tree.hijo[1].valInt {
							tree.valBool = true
						} else {
							tree.valBool = false
						}

					case false:
						if tree.hijo[0].valFloat > tree.hijo[1].valFloat {
							tree.valBool = true
						} else {
							tree.valBool = false
						}

					}

				case TKN_GETHAN:
					switch p1.isIntType {
					case true:
						if tree.hijo[0].valInt >= tree.hijo[1].valInt {
							tree.valBool = true
						} else {
							tree.valBool = false
						}

					case false:
						if tree.hijo[0].valFloat >= tree.hijo[1].valFloat {
							tree.valBool = true
						} else {
							tree.valBool = false
						}

					}

				case TKN_EQUAL:
					switch p1.isIntType {
					case true:
						if tree.hijo[0].valInt == tree.hijo[1].valInt {
							tree.valBool = true
						} else {
							tree.valBool = false
						}

					case false:
						if tree.hijo[0].valFloat == tree.hijo[1].valFloat {
							tree.valBool = true
						} else {
							tree.valBool = false
						}

					}

				case TKN_NEQUAL:
					switch p1.isIntType {
					case true:
						if tree.hijo[0].valInt != tree.hijo[1].valInt {
							tree.valBool = true
						} else {
							tree.valBool = false
						}

					case false:
						if tree.hijo[0].valFloat != tree.hijo[1].valFloat {
							tree.valBool = true
						} else {
							tree.valBool = false
						}

					}
				} /* case op */
			} //fin else
		} else { //alguno de los dos hijos tiene alugun tipo de error
			if p1.typeError || p2.typeError {
				tree.typeError = true
			}
			if p1.undeclaredError || p2.undeclaredError {
				tree.undeclaredError = true
			}
		} /* OpK */
	}

}
