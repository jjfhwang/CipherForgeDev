// cmd/cipherforgedev/main.go
package main

import (
"flag"
"log"
"os"

"cipherforgedev/internal/cipherforgedev"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := cipherforgedev.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
