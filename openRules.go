package main

import (
	"os"
)

func openRules(rulesFileName string) {
	os.OpenFile(rulesFileName, os.O_RDWR|os.O_CREATE, 0644)
}
