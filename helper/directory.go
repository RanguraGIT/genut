package helper

import (
	"os"
	"path/filepath"
	"strings"
)

type directory_config struct {
	path    string
	skip    []string
	process []string
}

var (
	mockdir string
	walker  string
)

// function to set directory config
func NewDirectoryConfig(path string, skip []string, process []string) *directory_config {
	return &directory_config{
		path:    path,
		skip:    skip,
		process: process,
	}
}

// function to get path of the file
func (d *directory_config) walker(target string) {
	filepath.Walk(d.path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			if d.IsContain(d.skip, info.Name()) {
				return filepath.SkipDir
			} else if info.Name() == target {
				walker = path
			}
		}
		return nil
	})
}

// function to check if the want is in the target list
func (d *directory_config) IsContain(target []string, want string) bool {
	for _, t := range target {
		if t == want {
			return true
		}
	}
	return false
}

// function to check if the some of the want is in the target list
func (d *directory_config) IsContainAny(target []string, want []string) bool {
	for _, t := range target {
		for _, w := range want {
			if t == w {
				return true
			}
		}
	}
	return false
}

// function to check if the want is in the path list
func (d *directory_config) GetPath(path string, want string) string {
	workdir := strings.TrimPrefix(strings.ReplaceAll(path, d.path, ""), string(os.PathSeparator))
	service := strings.Split(workdir, string(os.PathSeparator))[0]
	microse := strings.Split(workdir, string(os.PathSeparator))[1]

	switch want {
	// returning wanted path
	case "workdir":
		return workdir
	case "service":
		return service
	case "microservice":
		return microse
	}

	return ""
}

// function to check if the want is in the target list
func (d *directory_config) GetPackage(target []string, want string) string {
	for _, t := range strings.Split(want, string(filepath.Separator)) {
		for _, w := range target {
			if t == w {
				return t
			}
		}
	}
	return ""
}

// function to set mock directory
func (d *directory_config) GetMockDirectory(packages string, services string, others bool, types bool) string {
	i := filepath.Join("internal", "infrastructure", "mocks")
	r := filepath.Join("internal", "repository", "mocks")
	s := filepath.Join("internal", "service", "mocks")
	o := filepath.Join("internal", "others", "mocks")
	p := filepath.Join("pkg", "mocks")

	switch packages {
	// returning package path
	case "infrastructure":
		mockdir = i
	case "service", "usecase":
		mockdir = s
	case "repository":
		mockdir = r
	case "pkg":
		mockdir = p
	default:
		if others {
			mockdir = o
		}
	}

	if packages != "pkg" {
		// types true = microservice, false = service
		if types {
			mockdir = filepath.Join("services", services, mockdir)
		} else {
			d.walker("internal")
			mockdir = filepath.Join(walker, mockdir)
		}
	}

	return mockdir
}
