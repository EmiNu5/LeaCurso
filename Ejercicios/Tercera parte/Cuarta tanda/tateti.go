package main

import "fmt"

func main() {

	var player1 string
	var player2 string
	var x int
	var y int
	var tablero [3][3]string
	const minValor int = 1
	const maxValor int = 3
	var casillaCorrecta bool
	var i int
	var ganador bool
	var repetir string
	var gameover bool

	var jugadorActual string
	var fichaActual string

	fmt.Println("♦♦♦- BIENVENIDO -♦♦♦ \n")
	fmt.Println("Player1 Ingresa tu nombre: ")
	fmt.Scan(&player1)
	fmt.Println("Player2 ingresa tu nombre: ")
	fmt.Scan(&player2)
	fmt.Printf("%s vs %s \n\n", player1, player2)

	for {

		/*DIBUJITO VACIO*/
		for x = 0; x < len(tablero); x++ {
			for y = 0; y < len(tablero); y++ {

				if y < 2 {
					fmt.Print("   |")
				} else {
					fmt.Println("   ")
				}
			}
			if x < 2 {
				fmt.Println("-----------")
			}
		}
		for !gameover { /*Entrada y parametros */

			if i%2 == 0 {
				jugadorActual = player1
				fichaActual = "X"
			} else {
				jugadorActual = player2
				fichaActual = "O"
			}

			casillaCorrecta = false
			for !casillaCorrecta {
				fmt.Printf(" %s ingresa la (X,Y) donde ira la ficha X\n", jugadorActual)
				fmt.Scanf("\n%d,%d", &x, &y)
				if (x < minValor || x > maxValor) || (y < minValor || y > maxValor) {
					fmt.Println("Valores fuera de rango, prueba numeros del 1 al 3\n")
				} else if tablero[x-1][y-1] == "X" || tablero[x-1][y-1] == "O" {
					fmt.Println("Ouchh casilla ocupada prueba otra \n")
				} else {
					tablero[x-1][y-1] = fichaActual
					casillaCorrecta = true
				}
			}
			i++
			/*CASOS DE GANADOR*/
			if tablero[0][0] == tablero[0][1] && tablero[0][0] == tablero[0][2] {
				if tablero[0][0] == "X" || tablero[0][0] == "O" {
					ganador = true
				}
			} else if tablero[0][0] == tablero[1][0] && tablero[0][0] == tablero[2][0] {
				if tablero[0][0] == "X" || tablero[0][0] == "O" {
					ganador = true
				}

			} else if tablero[0][0] == tablero[1][1] && tablero[0][0] == tablero[2][2] {
				if tablero[0][0] == "X" || tablero[0][0] == "O" {
					ganador = true
				}
			}

			if tablero[1][1] == tablero[0][1] && tablero[1][1] == tablero[2][1] {
				if tablero[1][1] == "X" || tablero[1][1] == "O" {
					ganador = true
				}
			} else if tablero[1][1] == tablero[0][2] && tablero[1][1] == tablero[2][0] {
				if tablero[1][1] == "X" || tablero[1][1] == "O" {
					ganador = true
				}
			} else if tablero[1][1] == tablero[1][0] && tablero[1][1] == tablero[1][2] {
				if tablero[1][1] == "X" || tablero[1][1] == "O" {
					ganador = true
				}
			}

			if tablero[2][2] == tablero[2][1] && tablero[2][2] == tablero[2][0] {
				if tablero[2][2] == "X" || tablero[2][2] == "O" {
					ganador = true
				}
			} else if tablero[2][2] == tablero[1][2] && tablero[2][2] == tablero[0][2] {
				if tablero[2][2] == "X" || tablero[2][2] == "O" {
					ganador = true
				}

			}

			/*DIBUJITO*/
			fmt.Printf(" %1s | %1s | %1s\n", tablero[0][0], tablero[0][1], tablero[0][2])
			fmt.Println("-----------")
			fmt.Printf(" %1s | %1s | %1s\n", tablero[1][0], tablero[1][1], tablero[1][2])
			fmt.Println("-----------")
			fmt.Printf(" %1s | %1s | %1s\n", tablero[2][0], tablero[2][1], tablero[2][2])

			if ganador {
				fmt.Print("Has ganado ")
				if i%2 != 0 {
					fmt.Print(player1)
				} else {
					fmt.Print(player2)
				}
			} else if i == 9 {
				fmt.Print("Es un empate")
			}
			for ganador || i == 9 {
				fmt.Println("\n Quieres volver a jugar? s/n ")
				fmt.Scan(&repetir)
				if repetir == "s" || repetir == "n" {
					break
				}
			}

			if repetir == "s" {

				tablero = [3][3]string{}
				gameover = false
				ganador = false
				repetir = ""
				i = 0
				break
			} else if repetir == "n" {
				fmt.Println("Hasta la proxima")
				gameover = true
			}
		}
		if gameover {
			break
		}
	}
}
