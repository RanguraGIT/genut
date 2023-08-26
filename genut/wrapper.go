package genut

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/RanguraGIT/genut/config"
	"github.com/RanguraGIT/genut/helper"
)

type wrapper struct {
	skip         []string
	process      []string
	others       bool
	workingtable string
}

var (
	mockObjects map[string]string
	packageTemp string
)

func GenWrapper() error {
	wrap := configWrapper()
	walkErr := filepath.Walk(wrap.workingtable, func(path string, info os.FileInfo, walkerr error) error {
		if walkerr != nil {
			return walkerr
		}

		if info.IsDir() {
			if helper.NewContainsConfig(wrap.skip, info.Name()).Contains() {
				return filepath.SkipDir
			}
		} else {
			if helper.NewContainsConfig(wrap.process, path).ContainsAny() && strings.HasSuffix(info.Name(), ".go") && info.Name() != "wrapper.go" {
				err := wrap.genWrapperContent(path)
				if err != nil {
					return err
				}
			}
		}

		return nil
	})

	if walkErr != nil {
		fmt.Println("Error:", walkErr)
	}

	return nil
}

func configWrapper() *wrapper {
	wt, _ := helper.Worktable()
	skip, process, others := config.NewDirectoriesConfig(wt).LoadMockConfig()

	return &wrapper{
		skip:         skip,
		process:      process,
		others:       others,
		workingtable: wt,
	}
}

func (w *wrapper) genWrapperContent(path string) error {
	fmt.Println("Processing:", path)

	process := config.NewDirectoriesConfig(w.workingtable).GetProcess()
	packages := helper.NewContainsConfig(process, path).GetDir()
	if packageTemp == "" {
		packageTemp = packages
	}

	structs, structsErr := structName(path)
	if structsErr != nil {
		return structsErr
	}

	if mockObjects == nil || packageTemp != packages {
		packageTemp = packages
		mockObjects = make(map[string]string)
	}

	mockObjects[structs] = packages

	fileErr := w.file(path, packages)
	if fileErr != nil {
		return fileErr
	}

	return nil
}

func (w *wrapper) file(path string, packet string) error {
	microservicePath := strings.ReplaceAll(filepath.Dir(path), w.workingtable, "")
	microservicePath = strings.TrimPrefix(microservicePath, string(os.PathSeparator))
	microservice := strings.Split(microservicePath, string(os.PathSeparator))[0]

	if microservice == "services" {
		microservice = strings.Split(microservicePath, string(os.PathSeparator))[1]
	}

	dir := filepath.Join(w.workingtable, microservicePath, "wrapper.go")
	wrapperFile, err := os.Create(dir)
	if err != nil {
		return err
	}
	defer wrapperFile.Close()

	wrapperFile.WriteString("package mocks\n\n")
	wrapperFile.WriteString("import \"github.com/golang/mock/gomock\"\n\n")
	wrapperFile.WriteString(fmt.Sprintf("type %s struct {\n", packet))

	for mock := range mockObjects {
		wrapperFile.WriteString(fmt.Sprintf("\t*%s\n", mock))
	}

	wrapperFile.WriteString("}\n\n")
	wrapperFile.WriteString(fmt.Sprintf("func Init(ctrl *gomock.Controller) %s {\n", packet))
	wrapperFile.WriteString(fmt.Sprintf("\treturn %s{\n", packet))

	for mock := range mockObjects {
		wrapperFile.WriteString(fmt.Sprintf("\t\t%s: New%s(ctrl),\n", mock, mock))
	}

	wrapperFile.WriteString("\t}\n}\n")
	return nil
}

func structName(path string) (string, error) {
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
