package main

import (
  "fmt"
  "os"
)

func main() {
  var rutaArchivo string
  fmt.Println("Input your dictionary location:")
  fmt.Scan(&rutaArchivo)

  archivoAbierto, err_parser := os.Open(rutaArchivo)
  if err_parser != nil {
    fmt.Println("Error opening the dictionary!", err_parser)
  return
  }
defer archivoAbierto.Close() 
}
