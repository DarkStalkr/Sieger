// 1. This is the first part of the program, load the dictionary into memory

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
  // Ask user input dictionary
  //var archivo string
  //fmt.Println("Input your dictionary location:")
  //fmt.Scan(archivo)

  archivo, err_parser := os.Open("diccionario1.txt")
  if err_parser != nil {
    fmt.Println("Error al abrir el archivo!", err_parser)
    return
  }
  defer archivo.Close() // Cierra el archivo
  
  //Scanner linea por linea
  scanner := bufio.NewScanner(archivo)

  // Lee cada linea
  lineNum := 1
  for scanner.Scan() {
    linea := scanner.Text()
    fmt.Println("Linea %d: %s\n", lineNum, linea)
    lineNum++

    //Procesamiento de los parametros aquí



  }
  //Verifica errores durante escaneo
  if err_parser := scanner.Err(); err_parser != nil {
    fmt.Println("Error durante lectura", err_parser)
  } 

}
