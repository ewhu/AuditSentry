// cmd/auditsentry/main.go
package main

import (
"flag"
"log"
"os"

"auditsentry/internal/auditsentry"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := auditsentry.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
