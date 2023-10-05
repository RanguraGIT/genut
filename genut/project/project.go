package project

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/RanguraGIT/genut/genut/services"
)

type project struct {
	name     string
	services string
	version  string
}

// NewProjectConfig is a function to create new project config
func NewProjectConfig(name string, services string, version string) *project {
	return &project{
		name:     name,
		services: services,
		version:  version,
	}
}

// Create is a function to create new project
func (p *project) Create() {
	if _, err := os.Stat(filepath.Join(p.name)); !os.IsNotExist(err) {
		fmt.Printf("Project %s already exist\n", p.name)
		return
	}

	if err := os.Mkdir(filepath.Join(p.name), 0755); err != nil {
		fmt.Printf("Error when creating directory %s\n", p.name)
		return
	}

	if err := p.directory(); err != nil {
		return
	}

	if err := p.files(); err != nil {
		return
	}

	services.NewServiceConfig(p.name, p.services).Create()
}

// directory is a function to create directory
func (p *project) directory() error {
	directory := []string{
		"bin",
		"deploy",
		"deploy/cloud",
		"deploy/config",
		"deploy/local",
		"helper",
		"helper/dotenv",
		"migration",
		"migration/db",
		"migration/src",
		"pkg",
		"pkg/http_client",
		"services",
		"task-definition",
		"task-definition/dev",
		"task-definition/nft",
		"task-definition/preprod",
		"task-definition/preprod/stage1",
		"task-definition/preprod/stage2",
		"task-definition/preprod/stage3",
		"task-definition/preprod/stage4",
		"task-definition/preprod/stage5",
		"task-definition/prod",
	}

	for _, dir := range directory {
		if err := os.Mkdir(filepath.Join(p.name, dir), 0755); err != nil {
			fmt.Printf("Error when creating directory %s\n", dir)
			return err
		}
	}
	return nil
}

// files is a function to create files
func (p *project) files() error {
	files := map[string]string{
		"bin/.gitignore":                   "gitignore|default",
		"helper/dotenv/dotenv.go":          "helper|dotenv",
		"migration/.env.sample":            "env|migration",
		"pkg/http_client/client_http.go":   "http_client|client_http",
		"pkg/http_client/external_http.go": "http_client|external_http",
		"pkg/http_client/interface.go":     "http_client|interface",
		"pkg/http_client/request_http.go":  "http_client|request_http",
		"go.mod":                           "go.mod|default",
		".gitignore":                       "gitignore|root",
		".env.sample":                      "env|default",
	}

	for file, types := range files {
		types := strings.Split(types, "|")
		if _, err := os.Create(filepath.Join(p.name, file)); err != nil {
			fmt.Printf("Error when creating file %s\n", file)
			return err
		}

		err := p.calls(file, types[0], types[1])
		if err != nil {
			return err
		}
	}

	return nil
}

// calls is a function to call function based on class and type
func (p *project) calls(path string, class string, types string) error {
	switch class {
	case "helper":
		if err := p.helper(path, types); err != nil {
			return err
		}
	case "http_client":
		if err := p.http_client(path, types); err != nil {
			return err
		}
	case "gitignore":
		if err := p.gitignore(path, types); err != nil {
			return err
		}
	case "go.mod":
		if err := p.go_mod(path); err != nil {
			return err
		}
	case "env":
		if err := p.env(path, types); err != nil {
			return err
		}
	// case "dockerfile":
	// 	if err := p.dockerfile(path, types); err != nil {
	// 		return err
	// 	}
	default:
		return errors.New("Type not found: " + types)
	}
	return nil
}

// helper is a function to create helper
func (p *project) helper(target string, types string) error {
	var helper []string

	file, err := os.Create(filepath.Join(p.name, target))
	if err != nil {
		return errors.New("Error when creating directory " + target)
	}

	switch types {
	case "dotenv":
		helper = []string{
			"package dotenv",
			"import (",
			"	\"os\"",
			"	\"strconv\"",
			")",
			"",
			"// GetString read env variable or use a fallback string value",
			"func GetString(variable string, fallback string) string {",
			"	res := os.Getenv(variable)",
			"	if res != \"\" {",
			"		return res",
			"	}",
			"",
			"	return fallback",
			"}",
			"",
			"// GetInt read env variable or use a fallback integer value",
			"func GetInt(variable string, fallback int) int {",
			"	res := os.Getenv(variable)",
			"	if res != \"\" {",
			"		resInt, err := strconv.Atoi(res)",
			"		if err == nil {",
			"			return resInt",
			"		}",
			"	}",
			"",
			"	return fallback",
			"}",
			"",
			"// GetBool read env variable or use a fallback boolean value",
			"func GetBool(variable string, fallback bool) bool {",
			"	res := os.Getenv(variable)",
			"	if res != \"\" {",
			"		resBool, err := strconv.ParseBool(res)",
			"		if err == nil {",
			"			return resBool",
			"		}",
			"	}",
			"",
			"	return fallback",
			"}",
		}
	default:
		helper = []string{}
	}

	defer file.Close()

	for _, helper := range helper {
		file.WriteString(helper + "\n")
	}

	return nil
}

// http_client is a function to create http client
func (p *project) http_client(target string, types string) error {
	var http_client []string

	file, err := os.Create(filepath.Join(p.name, target))
	if err != nil {
		return errors.New("Error when creating directory " + target)
	}

	switch types {
	case "client_http":
		http_client = []string{
			"package httpclient",
			"",
			"import (",
			"	\"net/http\"",
			")",
			"",
			"type httpClient struct {",
			"	clientHTTP *http.Client",
			"}",
			"",
			"func InitHttpClient(client *http.Client) HTTPClientTemplate {",
			"	return &httpClient{",
			"		clientHTTP: client,",
			"	}",
			"}",
			"",
			"func (h *httpClient) Do(req *http.Request) (*http.Response, error) {",
			"	return h.clientHTTP.Do(req)",
			"}",
		}
	case "external_http":
		http_client = []string{
			"package httpclient",
			"",
			"import (",
			"	\"crypto/tls\"",
			"	\"net/http\"",
			"	\"net/url\"",
			"	\"time\"",
			"",
			fmt.Sprintf("	\"%s/helper/dotenv\"", p.name),
			")",
			"",
			"// InitClient initialize basic http client",
			"func InitClient() *http.Client {",
			"	tr := &http.Transport{",
			"		TLSClientConfig: &tls.Config{",
			"			InsecureSkipVerify: true,",
			"		},",
			"	}",
			"",
			"	if dotenv.GetString(\"APP_ENV\", \"\") != \"local\" && dotenv.GetBool(\"IS_USE_PROXY\", false) {",
			"		proxyURL, _ := url.Parse(dotenv.GetString(\"PROXY_DUNIAGAMES\", \"\"))",
			"		tr.Proxy = http.ProxyURL(proxyURL)",
			"	}",
			"",
			"	client := &http.Client{",
			"		Transport: tr,",
			"		Timeout:   10 * time.Second,",
			"	}",
			"	return client",
			"}",
			"",
			"// UseProxy initialize http client using proxy",
			"// func UseProxy() *http.Client {",
			"// 	tr := &http.Transport{",
			"// 		TLSClientConfig: &tls.Config{",
			"// 			InsecureSkipVerify: true,",
			"// 		},",
			"// 	}",
			"",
			"// 	if dotenv.APPENV() != \"local\" {",
			"// 		proxyURL, _ := url.Parse(dotenv.PROXYDUNIAGAMES())",
			"// 		tr.Proxy = http.ProxyURL(proxyURL)",
			"// 	}",
			"// 	client := &http.Client{",
			"// 		Transport: tr,",
			"// 		Timeout:   10 * time.Second,",
			"// 	}",
			"// 	return client",
			"// }",
		}
	case "interface":
		http_client = []string{
			"package httpclient",
			"",
			"import (",
			"	\"io\"",
			"	\"net/http\"",
			")",
			"",
			"// HTTPRequestTemplate Http request template wrapper",
			"type HTTPRequestTemplate interface {",
			"	NewRequest(method string, url string, body io.Reader) (*http.Request, error)",
			"}",
			"",
			"// HTTPClientTemplate HTTP client wrapper",
			"type HTTPClientTemplate interface {",
			"	Do(req *http.Request) (*http.Response, error)",
			"}",
		}
	case "request_http":
		http_client = []string{
			"package httpclient",
			"",
			"import (",
			"	\"io\"",
			"	\"net/http\"",
			")",
			"",
			"type httpRequest struct {",
			"}",
			"",
			"func InitHttpRequest() HTTPRequestTemplate {",
			"	return &httpRequest{}",
			"}",
			"",
			"func (*httpRequest) NewRequest(method string, url string, body io.Reader) (*http.Request, error) {",
			"	return http.NewRequest(method, url, body)",
			"}",
		}
	default:
		http_client = []string{}
	}

	defer file.Close()

	for _, http_client := range http_client {
		file.WriteString(http_client + "\n")
	}

	return nil
}

// gitignore is a function to create gitignore
func (p *project) gitignore(target string, types string) error {
	var (
		file      *os.File
		err       error
		gitignore []string
	)

	switch types {
	case "root":
		gitignore = []string{
			"vendor",
			"**/.env",
		}

		file, err = os.Create(filepath.Join(p.name, ".gitignore"))
		if err != nil {
			return errors.New("Error when creating directory " + target)
		}

	default:
		gitignore = []string{
			"*",
			"!.gitignore",
		}

		file, err = os.Create(filepath.Join(p.name, target))
		if err != nil {
			return errors.New("Error when creating directory " + target)
		}

		defer file.Close()
	}

	for _, gitignore := range gitignore {
		file.WriteString(gitignore + "\n")
	}

	return nil
}

// go_mod is a function to create go.mod
func (p *project) go_mod(target string) error {
	var go_mod []string

	file, err := os.Create(filepath.Join(p.name, target))
	if err != nil {
		return errors.New("Error when creating directory " + target)
	}

	go_mod = []string{
		"module " + p.name,
		"",
		fmt.Sprintf("go %s", p.version),
	}

	defer file.Close()

	for _, go_mod := range go_mod {
		file.WriteString(go_mod + "\n")
	}

	return nil
}

// env is a function to create env
func (p *project) env(target string, types string) error {
	var env []string

	file, err := os.Create(filepath.Join(p.name, target))
	if err != nil {
		return errors.New("Error when creating directory " + target)
	}

	switch types {
	case "migration":
		env = []string{
			"# DBMate Config",
			"DATABASE_URL=mysql://uname:pwd@127.0.0.1:3306/db_name",
			"DBMATE_MIGRATIONS_DIR=./src",
			"DBMATE_SCHEMA_FILE=./db/schema.sql",
		}
	default:
		env = []string{
			"# App Config",
			"# local, staging, or production",
			"APP_ENV=local",
			"APP_TIMEZONE=Asia/Jakarta",
			"APP_PORT=:8080",
			"PREFIX=/v1",
			"IS_USE_HTTPS=false",
			"APP_TLS_CERT_FILENAME=",
			"APP_TLS_KEY_FILENAME=",
			"URL_PRODUCT_IMAGE_DG=",
			"IS_NFT=false",
			"PROXY_DUNIAGAMES=",
			"IS_USE_PROXY=false",
			"",
			"# Forgerock Config",
			"FORGEROCK_CONFIG_FRALG=",
			"FORGEROCK_CONFIG_TIMEOUT=",
			"FORGEROCK_CONFIG_DOMAIN=",
			"FORGEROCK_CONFIG_CLIENT_ID=",
			"",
			"# MySQL Config",
			"MYSQL_HOST=localhost",
			"MYSQL_PORT=3306",
			"MYSQL_NAME=username",
			"MYSQL_PASS=password",
			"MYSQL_DB_NAME=db_name",
		}
	}

	defer file.Close()

	for _, env := range env {
		file.WriteString(env + "\n")
	}
	return nil
}

// func (p *project) dockerfile(target string, types string) error {
// 	var dockerfile []string

// 	file, err := os.Create(filepath.Join(p.name, target))
// 	if err != nil {
// 		return errors.New("Error when creating directory " + target)
// 	}

// 	switch types {
// 	case "cloud":
// 		dockerfile = []string{
// 			"FROM ___",
// 			"",
// 			"WORKDIR /app",
// 			"",
// 			"COPY bin/___ /app",
// 			"",
// 			"ENV SERVICENAME ___",
// 			"ENV APP_ENV /app",
// 			"ENV GO111MODULE on",
// 			"ENV CGO_ENABLED=0",
// 			"ENV GOOS=linux",
// 			"ENV GOARCH=amd64",
// 			"",
// 			"CMD [\"/app/___\"]",
// 		}
// 	default:
// 		dockerfile = []string{
// 			"FROM ___",
// 			"",
// 			"WORKDIR /app",
// 			"",
// 			"COPY bin/___ /app",
// 			"",
// 			"ENV SERVICENAME ___",
// 			"ENV APP_ENV /app",
// 			"ENV GO111MODULE on",
// 			"ENV CGO_ENABLED=0",
// 			"ENV GOOS=linux",
// 			"ENV GOARCH=amd64",
// 			"",
// 			"CMD [\"/app/___\"]",
// 		}
// 	}

// 	defer file.Close()

// 	for _, dockerfile := range dockerfile {
// 		file.WriteString(dockerfile + "\n")
// 	}
// 	return nil
// }
