package services

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type service struct {
	project  string
	services string
}

// NewServiceConfig is a function to create new service config
func NewServiceConfig(project, services string) *service {
	return &service{
		project:  project,
		services: services,
	}
}

// Create is a function to create new service
func (s *service) Create() {
	if err := s.directory(); err != nil {
		return
	}

	if err := s.files(); err != nil {
		return
	}
}

// directory is a function to create directory for service
func (s *service) directory() error {
	directory := []string{
		"cmd",
		"domain",
		"domain/application",
		"domain/entity",
		"domain/repository",
		"domain/service",
		"domain/usecase",
		"domain/value_object",
		"internal",
		"internal/delivery",
		"internal/delivery/http",
		"internal/delivery/request",
		"internal/delivery/response",
		"internal/infrastucture",
		"internal/repository",
		"internal/repository/mysql",
		"internal/repository/mysql/mapper",
		"internal/repository/mysql/model",
		"internal/service",
		"internal/usecase",
		"middleware",
		"testdata",
	}

	if _, err := os.Stat(filepath.Join(s.project, "services", s.services)); os.IsNotExist(err) {
		if err := os.Mkdir(filepath.Join(s.project, "services", s.services), 0755); err != nil {
			fmt.Printf("Error when creating directory %s\n", s.services)
			return err
		}
	}

	for _, dir := range directory {
		if err := os.Mkdir(filepath.Join(s.project, "services", s.services, dir), 0755); err != nil {
			fmt.Printf("Error when creating directory %s\n", dir)
			return err
		}
	}
	return nil
}

// files is a function to create files for service
func (s *service) files() error {
	files := map[string]string{}

	for file, types := range files {
		types := strings.Split(types, "|")
		if _, err := os.Create(filepath.Join(s.project, "services", s.services, file)); err != nil {
			fmt.Printf("Error when creating file %s\n", file)
			return err
		}

		err := s.calls(file, types[0], types[1])
		if err != nil {
			return err
		}
	}

	return nil
}

// calls is a function to call function to create file content
func (s *service) calls(file, types, content string) error {
	// switch types {
	// case "single":
	// 	if err := s.single(file, content); err != nil {
	// 		return err
	// 	}
	// case "multiple":
	// 	if err := s.multiple(file, content); err != nil {
	// 		return err
	// 	}
	// }
	return nil
}

// dockerfile
// check apakah service hanya ada 1 ataukah service ada banyak
// jika service hanya ada 1 maka dockerfile akan dibuat untuk single service
// jika service ada banyak maka dockerfile akan dibuat untuk multiple service
// apabila service dari 1 lalu diinstall service lain maka akan memindahkan dockerfile ke multiple service
