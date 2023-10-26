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

func generateFolderStructure() {
	err := os.Mkdir(internalPath, 0777)
	if err != nil {
		log.Fatal(err)
	}
	err = os.Mkdir(viewPath, 0777)
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
