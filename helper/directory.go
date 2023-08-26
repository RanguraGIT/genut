package helper

import (
	"fmt"
	"os"
	"path/filepath"
)

type mock_directory_config struct {
	path string
	skip []string
}

var (
	mockDir string
	walker  string
)

func NewMockDirectoryConfig(path string, skip []string) *mock_directory_config {
	return &mock_directory_config{
		path: path,
		skip: skip,
	}
}

func (m *mock_directory_config) setWalker(target string) {
	filepath.Walk(m.path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			if NewContainsConfig(m.skip, info.Name()).Contains() {
				return filepath.SkipDir
			} else if info.Name() == target {
				walker = path
			}
		}
		return nil
	})
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
		mockDir = filepath.Join("pkg", "mocks")
	case "others":
		if others {
			mockDir = filepath.Join("services", microservice, "internal", "others", "mocks")
		} else {
			fmt.Println(m.path, "is not in the process directory.")
		}
	}
	return mockDir
}

func (m *mock_directory_config) GetMockDir(packages string, microservice string, others bool) string {
	m.setWalker("internal")
	switch packages {
	case "infrastructure":
		mockDir = filepath.Join(walker, "infrastructure", "mocks")
	case "service", "usecase":
		mockDir = filepath.Join(walker, "service", "mocks")
	case "repository":
		mockDir = filepath.Join(walker, "repository", "mocks")
	case "pkg":
		mockDir = filepath.Join("pkg", "mocks")
	case "others":
		if others {
			mockDir = filepath.Join(walker, "others", "mocks")
		} else {
			fmt.Println(m.path, "is not in the process directory.")
		}
	}
	return mockDir
}
