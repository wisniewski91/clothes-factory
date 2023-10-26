package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	name         string
	capital      string
	capitalS     string
	nameS        string
	tableName    string
	internalPath string
	viewPath     string
)

func generateModule(moduleName string, tableName string) {
	setNames(moduleName, tableName)
	generateFolderStructure()
	generateTypeFile()
	generateRepositoryFile()
	generateServiceFile()
	generateHandlersFile()
	generateRoutesFile()
	generateCreateFile()
	generateEditFile()
	generateIndexFile()
}

func main() {
	moduleName := os.Args[1]
	tableName := os.Args[2]
	generateModule(moduleName, tableName)
}

func generateTypeFile() {

	f, err := os.Create(filepath.Join(internalPath, "type.go"))
	if err != nil {
		log.Fatal(err)
	}

	templateFile, err := os.ReadFile("./console/type_template.txt")
	if err != err {
		log.Fatal(err)
	}

	fileGo := string(templateFile)
	fileGo = strings.ReplaceAll(fileGo, "{{name}}", name)
	fileGo = strings.ReplaceAll(fileGo, "{{nameS}}", nameS)
	fileGo = strings.ReplaceAll(fileGo, "{{capitalName}}", capital)
	fileGo = strings.ReplaceAll(fileGo, "{{capitalNameS}}", capitalS)
	fileGo = strings.ReplaceAll(fileGo, "{{tableName}}", tableName)

	os.WriteFile(filepath.Join(internalPath, "type.go"), []byte(fileGo), 0777)
	defer f.Close()

}

func generateServiceFile() {

	f, err := os.Create(filepath.Join(internalPath, "service.go"))
	if err != nil {
		log.Fatal(err)
	}

	templateFile, err := os.ReadFile("./console/service_template.txt")
	if err != err {
		log.Fatal(err)
	}

	fileGo := string(templateFile)
	fileGo = strings.ReplaceAll(fileGo, "{{name}}", name)
	fileGo = strings.ReplaceAll(fileGo, "{{nameS}}", nameS)
	fileGo = strings.ReplaceAll(fileGo, "{{capitalName}}", capital)
	fileGo = strings.ReplaceAll(fileGo, "{{capitalNameS}}", capitalS)

	os.WriteFile(filepath.Join(internalPath, "service.go"), []byte(fileGo), 0777)
	defer f.Close()
}

func generateRepositoryFile() {
	f, err := os.Create(filepath.Join(internalPath, "repository.go"))
	if err != nil {
		log.Fatal(err)
	}

	templateFile, err := os.ReadFile("./console/repository_template.txt")
	if err != err {
		log.Fatal(err)
	}

	fileGo := string(templateFile)
	fileGo = strings.ReplaceAll(fileGo, "{{name}}", name)
	fileGo = strings.ReplaceAll(fileGo, "{{nameS}}", nameS)
	fileGo = strings.ReplaceAll(fileGo, "{{capitalName}}", capital)
	fileGo = strings.ReplaceAll(fileGo, "{{capitalNameS}}", capitalS)

	os.WriteFile(filepath.Join(internalPath, "repository.go"), []byte(fileGo), 0777)
	defer f.Close()
}

func generateHandlersFile() {
	f, err := os.Create(filepath.Join(internalPath, "handlers.go"))
	if err != nil {
		log.Fatal(err)
	}

	templateFile, err := os.ReadFile("./console/handlers_template.txt")
	if err != err {
		log.Fatal(err)
	}

	fileGo := string(templateFile)
	fileGo = strings.ReplaceAll(fileGo, "{{name}}", name)
	fileGo = strings.ReplaceAll(fileGo, "{{nameS}}", nameS)
	fileGo = strings.ReplaceAll(fileGo, "{{capitalName}}", capital)
	fileGo = strings.ReplaceAll(fileGo, "{{capitalNameS}}", capitalS)

	os.WriteFile(filepath.Join(internalPath, "handlers.go"), []byte(fileGo), 0777)
	defer f.Close()
}

func generateRoutesFile() {
	f, err := os.Create(filepath.Join(internalPath, "routes.go"))
	if err != nil {
		log.Fatal(err)
	}

	templateFile, err := os.ReadFile("./console/routes_template.txt")
	if err != err {
		log.Fatal(err)
	}

	fileGo := string(templateFile)
	fileGo = strings.ReplaceAll(fileGo, "{{nameS}}", nameS)

	os.WriteFile(filepath.Join(internalPath, "routes.go"), []byte(fileGo), 0777)
	defer f.Close()
}

func generateIndexFile() {
	f, err := os.Create(filepath.Join(viewPath, "index.html"))
	if err != nil {
		log.Fatal(err)
	}

	templateFile, err := os.ReadFile("./console/templates/views/index.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileHtml := string(templateFile)
	os.WriteFile(filepath.Join(viewPath, "index.html"), []byte(fileHtml), 0777)
	defer f.Close()
}
func generateCreateFile() {
	f, err := os.Create(filepath.Join(viewPath, "modal", "create.html"))
	if err != nil {
		log.Fatal(err)
	}

	templateFile, err := os.ReadFile("./console/templates/views/create.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileHtml := string(templateFile)
	os.WriteFile(filepath.Join(viewPath, "modal", "create.html"), []byte(fileHtml), 0777)
	defer f.Close()
}

func generateEditFile() {
	f, err := os.Create(filepath.Join(viewPath, "modal", "edit.html"))
	if err != nil {
		log.Fatal(err)
	}

	templateFile, err := os.ReadFile("./console/templates/views/edit.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileHtml := string(templateFile)
	os.WriteFile(filepath.Join(viewPath, "modal", "edit.html"), []byte(fileHtml), 0777)
	defer f.Close()
}

func generateFolderStructure() {
	err := os.Mkdir(internalPath, 0777)
	if err != nil {
		log.Fatal(err)
	}
	err = os.Mkdir(viewPath, 0777)
	if err != nil {
		log.Fatal(err)
	}
	err = os.Mkdir(filepath.Join(viewPath, "modal"), 0777)
	if err != nil {
		log.Fatal(err)
	}
}

func setNames(moduleName string, table string) {
	name = moduleName
	caser := cases.Title(language.English)
	capital = caser.String(moduleName)
	capitalS = capital + "s"
	nameS = name + "s"
	tableName = table
	name = moduleName
	internalPath = filepath.Join("internal", nameS)
	viewPath = filepath.Join("views", "admin", "modules", nameS)
}
