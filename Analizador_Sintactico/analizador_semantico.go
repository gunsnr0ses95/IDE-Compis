/*
Proyecto Compiladores
"Analizador Semántico"
Creado por:
  - Héctor Alan López Díaz
  - Mónica Fabiola Montañez Briano
ISC 7ºA
*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

//variables Globales**************************************
var inBlock = false
var tipoActual = ""
var tem = 0
var eti = 0
var salvador = ""

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
	var p1, p2, p3 *TreeNode
	var l *BucketListRec
	switch tree.kind.stmt {
	case SELECCION:
		if tree.hermano == nil {
			salvador = "L" + strconv.Itoa(eti)
		} else {
			tree.hermano.etiqueta = "L" + strconv.Itoa(eti)
			salvador = tree.hermano.etiqueta
		}
		eti++
		p1 = tree.hijo[0]
		p2 = tree.hijo[1]
		p3 = tree.hijo[2]
		semantico(p1)
		genCuadruplos(tree.hijo[0].token.lexema, tree.hijo[0].hijo[0].token.lexema, tree.hijo[0].hijo[1].token.lexema, tree.hijo[0].temporal)
		genCuadruplos("if_false", tree.hijo[0].temporal, salvador, "_")
		semantico(p2)
		semantico(p3)
		genCuadruplos("label", salvador, "_", "_")

	case REPETICION:
		if strings.Compare(tree.etiqueta, "") == 0 {
			tree.etiqueta = "L" + strconv.Itoa(eti)
			eti++
		}
		genCuadruplos("label", tree.etiqueta, "_", "_")
		p1 = tree.hijo[0]
		p2 = tree.hijo[1]
		semantico(p1)
		semantico(p2)
		genCuadruplos(tree.hijo[1].token.lexema, tree.hijo[1].hijo[0].token.lexema, tree.hijo[1].hijo[1].token.lexema, tree.hijo[1].temporal)
		genCuadruplos("if_false", tree.hijo[1].temporal, tree.etiqueta, "_")

	case ITERACION:

		if strings.Compare(tree.etiqueta, "") == 0 {
			tree.etiqueta = "L" + strconv.Itoa(eti)
			eti++
		}
		if strings.Compare(tree.hermano.etiqueta, "") == 0 {
			tree.hermano.etiqueta = "L" + strconv.Itoa(eti)
			eti++
		}
		if tree.hermano == nil {
			salvador = "L" + strconv.Itoa(eti)
		} else {
			tree.hermano.etiqueta = "L" + strconv.Itoa(eti)
			salvador = tree.hermano.etiqueta
		}
		genCuadruplos("label", tree.etiqueta, "_", "_")
		p1 = tree.hijo[0]
		p2 = tree.hijo[1]
		semantico(p1)
		if tree.hijo[0].kind.exp == CONSTK {
			genCuadruplos("if_false", tree.hijo[0].token.lexema, salvador, "_")

		} else {
			genCuadruplos(tree.hijo[0].token.lexema, tree.hijo[0].hijo[0].token.lexema, tree.hijo[0].hijo[1].token.lexema, tree.hijo[0].temporal)
			genCuadruplos("if_false", tree.hijo[0].temporal, salvador, "_")
		}
		semantico(p2)
		genCuadruplos("label", salvador, "_", "_")
		if tree.hijo[0].kind.exp == CONSTK {
			genCuadruplos("if_true", tree.hijo[0].token.lexema, tree.etiqueta, "_")
		} else {
			genCuadruplos("if_true", tree.hijo[0].temporal, tree.etiqueta, "_")
		}

	case BLOQUE:
		inBlock = true
		semantico(tree.hijo[0])
		inBlock = false
	case ASIGNACION:
		//if !inBlock {
		l = st_lookup(tree.token.lexema)
		if l != nil {
			tipoActual = l.tipo
			semantico(tree.hijo[0])
			if strings.Compare(tree.hijo[0].temporal, "") == 0 {
				genCuadruplos("assing", tree.hijo[0].token.lexema, tree.token.lexema, "_")
			} else {
				genCuadruplos("assing", tree.hijo[0].temporal, tree.token.lexema, "_")
			}

			//Linea no copiada
			if !tree.hijo[0].typeError || !tree.hijo[0].undeclaredError {
				if (tree.hijo[0].isIntType && (strings.Compare(l.tipo, "Int") == 0)) || (!tree.hijo[0].isIntType && strings.Compare(l.tipo, "Float") == 0) || (!tree.hijo[0].isIntType && strings.Compare(l.tipo, "Bool") == 0) {
					st_insert(tree.token, tree.token.nline, tree.hijo[0].valInt, tree.hijo[0].valFloat, tree.hijo[0].valBool, l.tipo, false, true, memloc)
					memloc++
					tree.valInt = tree.hijo[0].valInt
					tree.valFloat = tree.hijo[0].valFloat
					tree.valBool = tree.hijo[0].valBool
				} else {
					if tree.hijo[0].isIntType && (strings.Compare(l.tipo, "Float") == 0) {
						tree.valFloat = float64(tree.hijo[0].valInt)
					} else {
						writerSymInfo.WriteString("Error: Tipos diferentes. Variables " + tree.token.lexema + " int=" + strconv.Itoa(l.valI) + " float=" + strconv.FormatFloat(tree.hijo[0].valFloat, 'f', -1, 64) + "\n")
						tree.valInt = l.valI
						tree.valFloat = l.valF
					}
				}
			}
		} /*else {
			tree.undeclaredError = true
			writerSymInfo.WriteString("Variable no declarada:" + tree.token.lexema + " No. Linea: " + strconv.Itoa(tree.token.nline) + "\n")
		}*/
		//}
	case READ:
		l = st_lookup(tree.token.lexema)
		if l != nil {
			l.haveVal = false
			l.valF = 0.0
			l.valI = 0
		} else {
			writerSymInfo.WriteString("Variable no declarada:" + tree.token.lexema + " No. Linea: " + strconv.Itoa(tree.token.nline) + "\n")
		}
		genCuadruplos("read", tree.token.lexema, "_", "_")

	case WRITE:
		l = st_lookup(tree.hijo[0].token.lexema)
		if l == nil {
			writerSymInfo.WriteString("Variable no declarada:" + tree.hijo[0].token.lexema + " No. Linea: " + strconv.Itoa(tree.token.nline) + "\n")
		}
		genCuadruplos("write", tree.hijo[0].token.lexema, "_", "_")
	}
}
func inExp(tree *TreeNode) {
	var p1, p2 *TreeNode
	var l *BucketListRec
	switch tree.kind.exp {
	case IDK:
		l = st_lookup(tree.token.lexema)
		if l != nil {
			if strings.Compare(l.tipo, "Int") == 0 {
				tree.isIntType = true
			} else {
				tree.isIntType = false
			}
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
		tree.temporal = "T" + strconv.Itoa(tem)
		tem++
		p1 = tree.hijo[0]
		p2 = tree.hijo[1]
		semantico(p1)
		semantico(p2)
		if tree.token.tokenval == TKN_ADD || tree.token.tokenval == TKN_MINUS || tree.token.tokenval == TKN_DIVISION || tree.token.tokenval == TKN_PRODUCT {

			if strings.Compare(tree.hijo[0].temporal, "") == 0 {
				genCuadruplos(tree.token.lexema, tree.hijo[0].token.lexema, tree.hijo[1].token.lexema, tree.temporal)
			} else {
				genCuadruplos(tree.token.lexema, tree.hijo[0].temporal, tree.hijo[1].token.lexema, tree.temporal)
			}
		}
		if !p1.typeError || !p1.undeclaredError || !p2.typeError || !p2.undeclaredError {
			if (p1.isIntType && !p2.isIntType) || (!p1.isIntType && p2.isIntType) {
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
						tree.isIntType = true
						tree.valInt = tree.hijo[0].valInt + tree.hijo[1].valInt

					case false:
						tree.valFloat = tree.hijo[0].valFloat + tree.hijo[1].valFloat

					}

				case TKN_MINUS:
					switch p1.isIntType {
					case true:
						tree.isIntType = true
						tree.valInt = tree.hijo[0].valInt - tree.hijo[1].valInt

					case false:
						tree.valFloat = tree.hijo[0].valFloat - tree.hijo[1].valFloat

					}

				case TKN_PRODUCT:
					switch p1.isIntType {
					case true:
						tree.isIntType = true
						tree.valInt = tree.hijo[0].valInt * tree.hijo[1].valInt

					case false:
						tree.valFloat = tree.hijo[0].valFloat * tree.hijo[1].valFloat
					}

				case TKN_DIVISION:
					switch p1.isIntType {
					case true:
						tree.isIntType = true
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

func printTreeSemantico(tree *TreeNode) {
	valAux := ""
	var l *BucketListRec
	tabno++
	for tree != nil {
		printTabulacion()
		if tree.nodekind == STMTK {
			switch tree.kind.stmt {
			case PROGRAMA:
				fmt.Printf("Program")
				writer3.WriteString("Program")
			case SELECCION:
				fmt.Printf("If")
				writer3.WriteString("If")
			case REPETICION:
				fmt.Printf("Repeat")
				writer3.WriteString("Repeat")
			case ASIGNACION:
				fmt.Printf("Assign to: ")
				//writer3.WriteString("Assign to: " + tree.token.lexema + "\n")
				writer3.WriteString("Assign to: ")
				l = st_lookup(tree.token.lexema)
				if l == nil {
					fmt.Printf("%s - Error: Variable no declarada", tree.token.lexema)
					writer3.WriteString("Id: " + tree.token.lexema + " - Error: Variable no declarada")
				} else {
					switch l.tipo {
					case "Int":
						valAux = strconv.Itoa(tree.valInt)
					case "Float":
						valAux = strconv.FormatFloat(tree.valFloat, 'E', -1, 64)
					case "Bool":
						valAux = strconv.FormatBool(tree.valBool)
					}
					fmt.Printf("%s -> (%s, %s)", tree.token.lexema, l.tipo, valAux)
					writer3.WriteString(tree.token.lexema + " -> (" + l.tipo + ", " + valAux + ")")
				}
			case ITERACION:
				fmt.Printf("While")
				writer3.WriteString("While")
			case READ:
				fmt.Printf("Read: %s", tree.token.lexema)
				writer3.WriteString("Read: " + tree.token.lexema)
			case WRITE:
				fmt.Printf("Write")
				writer3.WriteString("Write")
			}
			fmt.Printf(" %s\n", tree.etiqueta)
			writer3.WriteString(" " + tree.etiqueta + "\n")
		} else if tree.nodekind == EXPK {
			switch tree.kind.exp {
			case OPK:
				if tree.token.tokenval == TKN_ADD || tree.token.tokenval == TKN_MINUS || tree.token.tokenval == TKN_DIVISION || tree.token.tokenval == TKN_PRODUCT {
					if tree.isIntType {
						fmt.Printf("Op: %s -> (%s)", tree.token.lexema, strconv.Itoa(tree.valInt))
						writer3.WriteString("Op: " + tree.token.lexema + " -> (" + strconv.Itoa(tree.valInt) + ")")
					} else {
						fmt.Printf("Op: %s -> (%s)", tree.token.lexema, strconv.FormatFloat(tree.valFloat, 'E', -1, 64))
						writer3.WriteString("Op: " + tree.token.lexema + " -> (" + strconv.FormatFloat(tree.valFloat, 'E', -1, 64) + ")")
					}
				} else {
					fmt.Printf("Op: %s -> (%s)", tree.token.lexema, strconv.FormatBool(tree.valBool))
					writer3.WriteString("Op: " + tree.token.lexema + " -> (" + strconv.FormatBool(tree.valBool) + ")")
				}
				fmt.Printf(" %s\n", tree.temporal)
				writer3.WriteString(" " + tree.temporal + "\n")

			case CONSTK:
				fmt.Printf("Const: %s\n", tree.token.lexema)
				writer3.WriteString("Const: " + tree.token.lexema + "\n")

			case IDK:
				if tree.token == nil {
					goto salir
				}
				fmt.Printf("Id: ")
				writer3.WriteString("Id: ")
				l = st_lookup(tree.token.lexema)
				if l == nil {
					fmt.Printf(" - Error: Variable no declarada\n", tree.token.lexema)
					writer3.WriteString(tree.token.lexema + " - Error: Variable no declarada\n")
				} else {
					switch l.tipo {
					case "Int":
						valAux = strconv.Itoa(tree.valInt)
					case "Float":
						valAux = strconv.FormatFloat(tree.valFloat, 'E', -1, 64)
					case "Bool":
						valAux = strconv.FormatBool(tree.valBool)
					}
					fmt.Printf("%s -> (%s, %s)\n", tree.token.lexema, l.tipo, valAux)
					writer3.WriteString(tree.token.lexema + " -> (" + l.tipo + ", " + valAux + ")\n")
				}
			}
		} else {
			fmt.Printf("Unknown node kind\n")
		}
		for i := 0; i < 3; i++ {
			printTreeSemantico(tree.hijo[i])
		}
		tree = tree.hermano
	}
salir:
	tabno--
}

func genCuadruplos(d1, d2, d3, d4 string) {
	writerCuad.WriteString("(" + d1 + "," + d2 + "," + d3 + "," + d4 + ")\n")

}
