package mocking

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/RanguraGIT/genut/config"
	"github.com/RanguraGIT/genut/helper"
	"github.com/RanguraGIT/genut/helper/checker"
	"github.com/RanguraGIT/genut/helper/installer"
)

type mockgen struct {
	skip         []string
	process      []string
	others       bool
	workingtable string
}

type wrapper struct {
	skip         []string
	process      []string
	others       bool
	workingtable string
}

var (
	con     = config.NewConfig(wt())
	mock    = configMockgen()
	modir   = helper.NewDirectoryConfig(mock.workingtable, mock.skip, mock.process)
	mockdir string
	wrap    = configWrapper()
	wrdir   = helper.NewDirectoryConfig(wrap.workingtable, wrap.skip, wrap.process)
	wrobj   = make(map[string]string)
	wrtemp  string
)

// function to get current working directory
func wt() string {
	wt, _ := helper.Worktable()
	return wt
}

// function to set mockgen config
func configMockgen() *mockgen {
	wt, _ := helper.Worktable()
	skip, process, others := con.LoadConfig()

	return &mockgen{
		skip:         skip,
		process:      process,
		others:       others,
		workingtable: wt,
	}
}

// function to set wrapper config
func configWrapper() *wrapper {
	wt, _ := helper.Worktable()
	skip, process, others := con.LoadMockConfig()

	return &wrapper{
		skip:         skip,
		process:      process,
		others:       others,
		workingtable: wt,
	}
}

// function to generate mockgen
func GenMockgen() {
	walkErr := filepath.Walk(mock.workingtable, func(path string, info os.FileInfo, walkerr error) error {
		fil := helper.NewFileConfig(wt(), path)
		if walkerr != nil {
			return walkerr
		}

		if info.IsDir() {
			if modir.IsContain(mock.skip, info.Name()) {
				return filepath.SkipDir
			}
		} else {
			if modir.IsContainAny(strings.Split(path, string(filepath.Separator)), mock.process) && fil.IsContain() && strings.HasSuffix(info.Name(), ".go") {
				if err := mock.genMockgenInterface(path); err != nil {
					return err
				}
			} else if mock.others && strings.HasSuffix(info.Name(), ".go") && fil.IsContain() {
				if err := mock.genMockgenInterface(path); err != nil {
					return err
				}
			}
		}

		return nil
	})

	if walkErr != nil {
		fmt.Println("Walk Error:", walkErr)
	}

	err := GenWrapper()

	if err != nil {
		fmt.Println("Generate Error:", err)
	}
}

// function to generate mockgen interface
func (m *mockgen) genMockgenInterface(path string) error {
	if !checker.Mockgen() {
		installer.Mockgen()
	}

	fil := helper.NewFileConfig(m.workingtable, path)
	workdir := modir.GetPath(path, "workdir")
	service := modir.GetPath(path, "service")
	packages := modir.GetPackage(m.process, workdir)

	if service == "services" {
		service = modir.GetPath(path, "microservice")
		mockdir = modir.GetMockDirectory(packages, service, m.others, true)
	} else {
		mockdir = modir.GetMockDirectory(packages, service, m.others, false)
	}

	filebase := filepath.Base(path)
	filename := strings.TrimSuffix(filebase, ".go")

	if filename == "interface" {
		filename = fil.GenFilename(workdir, packages)
	}

	if mockdir != "" {
		cmd := exec.Command("mockgen", "-source="+path, "-destination="+filepath.Join(mockdir, "mock_"+filename+".go"), "-package=mocks")
		_, err := cmd.Output()
		if err != nil {
			return err
		}

		fmt.Printf("Mock generated for:\n")
		fmt.Printf("%s => %s\n", workdir, filepath.Join(mockdir, "mock_"+filename+".go"))
	}

	return nil
}

// function to generate wrapper
func GenWrapper() error {
	walkErr := filepath.Walk(wrap.workingtable, func(path string, info os.FileInfo, walkerr error) error {
		if walkerr != nil {
			return walkerr
		}

		if info.IsDir() {
			if wrdir.IsContain(wrap.skip, info.Name()) {
				return filepath.SkipDir
			}
		} else {
			if wrdir.IsContainAny(strings.Split(path, string(filepath.Separator)), wrap.process) && strings.HasSuffix(info.Name(), ".go") && info.Name() != "wrapper.go" {
				if err := wrap.genWrapperContent(path); err != nil {
					return err
				}
			}
		}
		return nil
	})

	if walkErr != nil {
		fmt.Println("Walk Error:", walkErr)
	}
	return nil
}

// function to generate wrapper content
func (w *wrapper) genWrapperContent(path string) error {
	fmt.Printf("Wrapper generated for:\n")

	files := helper.NewFileConfig(w.workingtable, path)
	process := config.NewConfig(wt()).GetProcess()
	packages := modir.GetPackage(process, path)

	if wrtemp == "" {
		wrtemp = packages
	}

	if wrtemp != packages || wrobj == nil {
		wrobj = make(map[string]string)
		wrtemp = packages
	}

	structs, structsErr := files.GetStruct(path)
	if structsErr != nil {
		return structsErr
	}

	wrobj[structs] = packages

	// function to generate wrapper file
	fileErr := files.GenWrapperFile(packages, wrobj)
	if fileErr != nil {
		return fileErr
	}
	return nil
}
