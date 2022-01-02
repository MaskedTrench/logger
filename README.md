
# Go-Logger

Simple usage is:

```go
package main

import "github.com/MaskedTrench/logger"

func main() {
    l, _ := logger.BuildLogger(logger.DebugLvl) // Builds logger and error
    l.EnableFiles(true)                         // Enables files
    l.Inform("Hello World")                     // Sends inform message to file and console
}
```

## At the moment Logger contains

- Levels
- Pretty-print option
- Automatic folder output
- Parralel output: file and console

## So in next update there will be

- Safe exit
- Print formulas
- Auto 'hello world' printing
