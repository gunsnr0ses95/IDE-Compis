// analizador_semantico
package main

//variable,_ := strconv.ParseFloat("3.2",64)
import (
	"fmt"
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
			if l != nil {
				tipoActual = l.tipo
				semantico(tree.hijo[0])
				//Linea no copiada
				if !tree.hijo[0].typeError || !tree.hijo[0].undeclaredError {
					if (tree.hijo[0].isIntType && (strings.Compare(l.tipo, "Int") == 0)) || (!tree.hijo[0].isIntType && strings.Compare(l.tipo, "Float") == 0) || (!tree.hijo[0].isIntType && strings.Compare(l.tipo, "Bool") == 0) {
						st_insert(tree.token, tree.token.nline, tree.hijo[0].valInt, tree.hijo[0].valFloat, tree.hijo[0].valBool, l.tipo, false, true, memloc)
						memloc++
						tree.valInt = tree.hijo[0].valInt
						tree.valFloat = tree.hijo[0].valFloat
						tree.valBool = tree.hijo[0].valBool
					} else {
						writerSymInfo.WriteString("Error: Tipos diferentes. Variables " + tree.token.lexema + " int=" + strconv.Itoa(l.valI) + " float=" + strconv.FormatFloat(tree.hijo[0].valFloat, 'f', -1, 64) + "\n")
						tree.valInt = l.valI
						tree.valFloat = l.valF

					}
				}
			} /*else {
				tree.undeclaredError = true
				writerSymInfo.WriteString("Variable no declarada:" + tree.token.lexema + " No. Linea: " + strconv.Itoa(tree.token.nline) + "\n")
			}*/
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
		p1 = tree.hijo[0]
		p2 = tree.hijo[1]
		semantico(p1)
		semantico(p2)
		if !p1.typeError || !p1.undeclaredError || !p2.typeError || !p2.undeclaredError {
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

func printTreeSemantico(tree *TreeNode) {
	valAux := ""
	var l *BucketListRec
	tabno++
	for tree != nil {
		printTabulacion()
		if tree.nodekind == STMTK {
			switch tree.kind.stmt {
			case PROGRAMA:
				fmt.Printf("Program\n")
				writer3.WriteString("Program\n")
			case SELECCION:
				fmt.Printf("If\n")
				writer3.WriteString("If\n")
			case REPETICION:
				fmt.Printf("Repeat\n")
				writer3.WriteString("Repeat\n")
			case ASIGNACION:
				fmt.Printf("Assign to: ")
				//writer3.WriteString("Assign to: " + tree.token.lexema + "\n")
				writer3.WriteString("Assign to: ")
				l = st_lookup(tree.token.lexema)
				if l == nil {
					fmt.Printf("%s - Error: Variable no declarada\n", tree.token.lexema)
					writer3.WriteString("Id: " + tree.token.lexema + " - Error: Variable no declarada\n")
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
			case ITERACION:
				fmt.Printf("While\n")
				writer3.WriteString("While\n")
			case READ:
				fmt.Printf("Read: %s\n", tree.token.lexema)
				writer3.WriteString("Read: " + tree.token.lexema + "\n")
			case WRITE:
				fmt.Printf("Write\n")
				writer3.WriteString("Write\n")
			}
		} else if tree.nodekind == EXPK {
			switch tree.kind.exp {
			case OPK:
				if tree.token.tokenval == TKN_ADD || tree.token.tokenval == TKN_MINUS || tree.token.tokenval == TKN_DIVISION || tree.token.tokenval == TKN_PRODUCT {
					if tree.isIntType {
						fmt.Printf("Op: %s -> (%s)\n", tree.token.lexema, strconv.Itoa(tree.valInt))
						writer3.WriteString("Op: " + tree.token.lexema + " -> (" + strconv.Itoa(tree.valInt) + ")\n")
					} else {
						fmt.Printf("Op: %s -> (%s)\n", tree.token.lexema, strconv.FormatFloat(tree.valFloat, 'E', -1, 64))
						writer3.WriteString("Op: " + tree.token.lexema + " -> (" + strconv.FormatFloat(tree.valFloat, 'E', -1, 64) + ")\n")
					}
				} else {
					fmt.Printf("Op: %s -> (%s)\n", tree.token.lexema, strconv.FormatBool(tree.valBool))
					writer3.WriteString("Op: " + tree.token.lexema + " -> (" + strconv.FormatBool(tree.valBool) + ")\n")
				}

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
				/*if tree.tipo != 0 {
					fmt.Printf("-%s\n", tipoToString(tree.tipo))
					writer2.WriteString("-" + tipoToString(tree.tipo) + "\n")
				} else {
					fmt.Println()
					writer2.WriteString("\n")
				}*/ //no entendi esto pero lo comento por si acaso jaja
			}
		} else { //falta una parte del codigo de la tipa pero no lo copie porque nosotros no tenemos ningun nodekind tipo deck
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
