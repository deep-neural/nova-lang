package main // main.go

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Println("Usage: compiler <input-file> [output-file]")
        os.Exit(1)
    }

    inputFile := os.Args[1]
    outputFile := "output.ll"
    
    if len(os.Args) > 2 {
        outputFile = os.Args[2]
    }

    // Read input file
    source, err := ioutil.ReadFile(inputFile)
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        os.Exit(1)
    }

    // Create tokenizer
    tokenizer := NewTokenizer(string(source))

    // Create parser
    parser := NewParser(tokenizer)

    // Parse program
    program := parser.Parse()
    
    // Check for errors
    if len(parser.Errors()) > 0 {
        fmt.Println("Parser errors:")
        for _, err := range parser.Errors() {
            fmt.Println(err)
        }
        os.Exit(1)
    }

    // Create code generator
    moduleName := filepath.Base(inputFile)
    generator := NewCodeGenerator(moduleName)

    // Generate LLVM IR
    module, err := generator.Generate(program)
    if err != nil {
        fmt.Printf("Code generation error: %v\n", err)
        os.Exit(1)
    }

    // Save to file
    err = SaveModule(module, outputFile)
    if err != nil {
        fmt.Printf("Error saving output: %v\n", err)
        os.Exit(1)
    }

    fmt.Printf("Successfully compiled %s to %s\n", inputFile, outputFile)
}