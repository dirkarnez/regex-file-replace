package main

import (
	"fmt"
	"regexp"
)

var (
	regexp string
	filepath string 
	replacement string
)


func main() {
	flag.StringVar(&regexp, "regexp", "", "regexp")
	flag.StringVar(&filepath, "filepath", "", "filepath")
	flag.StringVar(&replacement, "replacement", "", "replacement")
	flag.Parse()
	
	re := regexp.MustCompile(fmt.Sprintf(`(?m)%s`, regexp)) // ^cmd = .*
	content := re.ReplaceAllString(ReadFileToString(filepath), replacement)
	fmt.Println(content) // cmd = ""Fusion360.exe""
	err := os.WriteFile(filepath, []byte(content), 0644)
}

func ReadFileToString(path string) (string, error) {
	b, err := os.ReadFile(path) // just pass the file name
	if err != nil {
		return err
	}
	return string(b), nil
}
