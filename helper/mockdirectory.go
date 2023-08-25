package helper

import (
	"fmt"
	"os"
	"path/filepath"
)

type mock_directory_config struct {
	path string
}

var (
	mockDir string
	walker  string
)

func NewMockDirectoryConfig(path string) *mock_directory_config {
	return &mock_directory_config{
		path: path,
	}
}

func (m *mock_directory_config) GetMockMicroDir(packages string, microservice string, others bool) string {
	switch packages {
	case "infrastructure":
		mockDir = filepath.Join("services", microservice, "internal", "infrastructure", "mocks")
	case "service", "usecase":
		mockDir = filepath.Join("services", microservice, "internal", "service", "mocks")
	case "repository":
		mockDir = filepath.Join("services", microservice, "internal", "repository", "mocks")
	case "pkg":
		mockDir = filepath.Join(microservice, "mocks")
	case "others":
		if others {
			mockDir = filepath.Join("services", microservice, "internal", "others", "mocks")
		} else {
			fmt.Println(m.path, "is not in the process directory.")
		}
	}

	return mockDir
}

func (m *mock_directory_config) GetMockDir(skip []string, packages string, microservice string, others bool) string {
	filepath.Walk(m.path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			if NewContainsConfig(skip, info.Name()).Contains() {
				return filepath.SkipDir
			} else if info.Name() == "internal" {
				walker = path
			}
		}

		return nil
	})

	switch packages {
	case "infrastructure":
		mockDir = filepath.Join(walker, "infrastructure", "mocks")
	case "service", "usecase":
		mockDir = filepath.Join(walker, "service", "mocks")
	case "repository":
		mockDir = filepath.Join(walker, "repository", "mocks")
	case "pkg":
		mockDir = filepath.Join(walker, "mocks")
	case "others":
		if others {
			mockDir = filepath.Join(walker, "others", "mocks")
		} else {
			fmt.Println(m.path, "is not in the process directory.")
		}
	}

	return mockDir
}
