package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"

	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

var (
	regex       string
	filepath    string
	replacement string
)

func main() {
	flag.StringVar(&regex, "regex", "", "regexp")
	flag.StringVar(&filepath, "filepath", "", "filepath")
	flag.StringVar(&replacement, "replacement", "", "replacement")
	flag.Parse()

	golangRegex := fmt.Sprintf(`(?m)%s`, regex)
	fmt.Println("received regex: ", golangRegex)
	fmt.Println("received filepath: ", filepath)
	fmt.Println("received replacement: ", replacement)

	re := regexp.MustCompile(golangRegex) // ^cmd = .*
	content, err := ReadFileUTF16(filepath)
	if err != nil {
		log.Fatal(err)
	}
	content = re.ReplaceAll(content, []byte(replacement))
	fmt.Println(string(content)) // cmd = ""Fusion360.exe""
	err = WriteUTF16(filepath, content)
	if err != nil {
		log.Fatal(err)
	}
}

func ReadFileToString(path string) (string, error) {
	b, err := os.ReadFile(path) // just pass the file name
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// Similar to ioutil.ReadFile() but decodes UTF-16.  Useful when
// reading data from MS-Windows systems that generate UTF-16BE files,
// but will do the right thing if other BOMs are found.
func ReadFileUTF16(filename string) ([]byte, error) {

	// Read the file into a []byte:
	raw, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Make an tranformer that converts MS-Win default to UTF8:
	win16be := unicode.UTF16(unicode.BigEndian, unicode.ExpectBOM)
	// Make a transformer that is like win16be, but abides by BOM:
	utf16bom := unicode.BOMOverride(win16be.NewDecoder())

	// Make a Reader that uses utf16bom:
	unicodeReader := transform.NewReader(bytes.NewReader(raw), utf16bom)

	// decode and print:
	decoded, err := io.ReadAll(unicodeReader)
	return decoded, err
}

func WriteUTF16(filename string, content []byte) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := unicode.UTF16(unicode.LittleEndian, unicode.ExpectBOM).NewEncoder()
	writer := transform.NewWriter(file, encoder)
	// Write the byte slice to the writer
	if _, err := writer.Write(content); err != nil {
		return err
	}

	// Flush the writer
	if err := writer.Close(); err != nil {
		return err
	}
	return nil
}
