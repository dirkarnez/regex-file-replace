package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile(`(?m)^cmd = .*`)
	
	fmt.Println(re.ReplaceAllString(`[Launcher]
stream = production
auid = AutodeskInc.Fusion360
cmd = ""C:\Users\runneradmin\AppData\Local\Autodesk\webdeploy\production\1a197c6e79bef01edef1dc4f317d9f597820e633\Fusion360.exe""
global = False`, `cmd = ""Fusion360.exe""`))
}
