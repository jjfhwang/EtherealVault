// cmd/etherealvault/main.go
package main

import (
"flag"
"log"
"os"

"etherealvault/internal/etherealvault"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := etherealvault.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
