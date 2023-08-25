package genut

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/RanguraGIT/genut/config"
	"github.com/RanguraGIT/genut/helper"
	"github.com/RanguraGIT/genut/helper/checker"
)

var (
	skip     []string
	process  []string
	others   bool
	mockDir  string
	mockName string
	isMicro  bool
)

// function to generate mockgen
func GenMockgen() {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error wd:", err)
		return
	}

	skip, process, others = config.NewDirectoriesConfig(cwd).LoadConfig()
	isMicro = false

	err = filepath.Walk(cwd, func(path string, info os.FileInfo, walkerr error) error {
		if walkerr != nil {
			return walkerr
		}

		if info.IsDir() {
			if helper.NewContainsConfig(skip, info.Name()).Contains() {
				return filepath.SkipDir
			}
		} else {
			if helper.NewContainsConfig(process, path).ContainsAny() && strings.HasSuffix(info.Name(), ".go") && helper.NewInterfaceConfig(path).Contains() {
				if err := generateMockForInterface(path, cwd); err != nil {
					return err
				}
			} else if others && strings.HasSuffix(info.Name(), ".go") && helper.NewInterfaceConfig(path).Contains() {
				if err := generateMockForInterface(path, cwd); err != nil {
					return err
				}
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
	}
}

// function to generate mock for interface
func generateMockForInterface(origin string, root string) error {
	if !checker.Mockgen() {
		cmd := exec.Command("go", "install", "github.com/golang/mock/mockgen")
		out, err := cmd.CombinedOutput()
		if err != nil {
			if strings.Contains(string(out), "malformed module path") {
				return errors.New("go install malformed module path.")
			}
			return errors.New("installing mockgen, You may need to install manually!")
		}
	}

	microservicePath := strings.ReplaceAll(origin, root, "")
	microservicePath = strings.TrimPrefix(microservicePath, string(os.PathSeparator))
	microservice := strings.Split(microservicePath, string(os.PathSeparator))[0]

	packages := helper.NewContainsConfig(process, microservicePath).GetDir()

	if microservice == "services" {
		isMicro = true
		microservice = strings.Split(microservicePath, string(os.PathSeparator))[1]
		mockDir = helper.NewMockDirectoryConfig(root).GetMockMicroDir(packages, microservice, others)
	} else {
		mockDir = helper.NewMockDirectoryConfig(root).GetMockDir(skip, packages, microservice, others)
	}

	fileBase := filepath.Base(origin)
	fileName := strings.TrimSuffix(fileBase, ".go")

	if fileName == "interface" {
		fileName = helper.NewFilenameConfig(microservicePath, packages).Generate()
	}

	if mockDir != "" {
		cmd := exec.Command("mockgen", "-source="+origin, "-destination="+filepath.Join(mockDir, "mock_"+fileName+".go"), "-package=mocks")
		_, err := cmd.Output()
		if err != nil {
			return err
		}

		fmt.Println("Mock generated for:", origin, "to", filepath.Join(mockDir, "mock_"+fileName+".go"))
	}

	return nil
}
