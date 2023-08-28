package helper

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/RanguraGIT/genut/config"
)

type file_config struct {
	workingtable string
	filePath     string
}

// func to set file config
func NewFileConfig(workingtable string, filePath string) *file_config {
	return &file_config{
		workingtable: workingtable,
		filePath:     filePath,
	}
}

// function to check if file exist
func GenConfig() bool {
	wd, _ := Worktable()
	files := filepath.Join(wd, ".genut.yml")
	file, err := os.Create(files)
	if err != nil {
		return false
	}

	defer file.Close()

	file.WriteString("directories:\n")
	file.WriteString(" # Skip is used for skiping folder who didn't want to generate\n")
	file.WriteString(" skip:\n")
	file.WriteString("  - vendor\n")
	file.WriteString("  - mocks\n")
	file.WriteString("  - utils\n")
	file.WriteString("  - helpers\n")
	file.WriteString(" # Process is used for processing folder who want to generate\n")
	file.WriteString(" process:\n")
	file.WriteString("  - service\n")
	file.WriteString("  - repository\n")
	file.WriteString("  - usecase\n")
	file.WriteString("  - pkg\n")
	file.WriteString("  - infrastructure\n")
	file.WriteString(" # Others is used for processing folder who want to generate\n")
	file.WriteString(" others: false\n")

	fmt.Println("Config file has been generated")

	return true
}

// function to check if the file contains an interface
func (c *file_config) IsContain() bool {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, c.filePath, nil, parser.AllErrors)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return false
	}

	for _, decl := range node.Decls {
		if genDecl, isGenDecl := decl.(*ast.GenDecl); isGenDecl {
			if genDecl.Tok == token.TYPE {
				for _, spec := range genDecl.Specs {
					if typeSpec, isTypeSpec := spec.(*ast.TypeSpec); isTypeSpec {
						if _, isInterface := typeSpec.Type.(*ast.InterfaceType); isInterface {
							return true
						}
					}
				}
			}
		}
	}

	return false
}

// function to generate filename
func (c *file_config) GenFilename(files string, packages string) string {
	directory := strings.SplitAfter(files, packages)

	mockname := directory[0]
	if len(directory) > 1 {
		mockname = directory[1]
	}

	file := strings.ReplaceAll(mockname, string(filepath.Separator), "_")
	filename := strings.TrimSuffix(file, ".go")
	return filename[1:]
}

// function to generate wrapper file
func (c *file_config) GenWrapperFile(packet string, waobj map[string]string) error {
	skip, process, others := config.NewConfig(c.workingtable).LoadConfig()
	dir := NewDirectoryConfig(c.workingtable, skip, process)

	workdir := dir.GetPath(c.filePath, "workdir")
	service := dir.GetPath(c.filePath, "service")

	if service == "services" {
		service = strings.Split(workdir, string(os.PathSeparator))[1]
		mockdir = dir.GetMockDirectory(packet, service, others, true)
	} else {
		mockdir = dir.GetMockDirectory(packet, service, others, false)
	}

	packet = "Wrapper"
	dirs := filepath.Join(mockdir, "wrapper.go")
	wrapperFile, err := os.Create(dirs)
	if err != nil {
		return err
	}
	defer wrapperFile.Close()

	wrapperFile.WriteString("package mocks\n\n")
	wrapperFile.WriteString("import \"github.com/golang/mock/gomock\"\n\n")
	wrapperFile.WriteString(fmt.Sprintf("type %s struct {\n", packet))

	for mock := range waobj {
		wrapperFile.WriteString(fmt.Sprintf("\t*%s\n", mock))
	}

	wrapperFile.WriteString("}\n\n")
	wrapperFile.WriteString(fmt.Sprintf("func Init(ctrl *gomock.Controller) %s {\n", packet))
	wrapperFile.WriteString(fmt.Sprintf("\treturn %s{\n", packet))

	for mock := range waobj {
		wrapperFile.WriteString(fmt.Sprintf("\t\t%s: New%s(ctrl),\n", mock, mock))
	}

	wrapperFile.WriteString("\t}\n}\n")
	return nil
}

// function to get struct name
func (c *file_config) GetStruct(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	defer file.Close()

	structName := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if strings.HasPrefix(line, "type ") && strings.HasSuffix(line, " struct {") {
			typeName := strings.TrimSuffix(strings.TrimPrefix(line, "type "), " struct {")
			if !strings.Contains(typeName, "Recorder") {
				structName = typeName
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	if structName == "" {
		return "", fmt.Errorf("struct name not found in file: %s", path)
	}

	return structName, nil
}
