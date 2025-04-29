// main.go
package main

import "nookli/cmd"

func main() {
	// Entrypoint: hand off to Cobra's Execute, which will run PersistentPreRun (InitDB) first.
	cmd.Execute()
}
