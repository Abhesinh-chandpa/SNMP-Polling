package main

import (
 "fmt"
 "os"

 "github.com/Abhesinh-chandpa/SNMP-Polling/internal/app"
)

func main() {
 if err := app.Run(os.Args[1:], os.Stdout); err != nil {
  fmt.Fprintln(os.Stderr, "error:", err)
  os.Exit(1)
 }
}
