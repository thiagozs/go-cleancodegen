package usecase

import (
	"embed"
	"fmt"
	"io/fs"
	"time"

	"github.com/lucasepe/spinner"

	"github.com/thiagozs/go-cleancodegen/internal/app/config"
	"github.com/thiagozs/go-cleancodegen/pkg/utils"
	"github.com/thiagozs/go-xutils"
	"github.com/thiagozs/go-xutils/files"
)

// embed the templates folder into the binary
//
//go:embed templates
var templates embed.FS

func CreateScaffold() error {

	u := xutils.New()
	f := u.Files()

	if config.Cfg.Path == "" {
		config.Cfg.Path = "/tmp"
	}

	if config.Cfg.Project == "" {
		config.Cfg.Project = "project"
	}

	config.Cfg.Project = utils.RemovePrefixSuffix(config.Cfg.Project, "/")

	projectPath := config.Cfg.Path + "/" + config.Cfg.Project

	if f.DirectoryExist(projectPath) {
		return fmt.Errorf("path %s not exists", config.Cfg.Path)
	}

	s := spinner.StartNew("Task 1: Processing...")
	for _, folder := range config.Folders {
		s.SetText(projectPath + "/" + folder)

		if err := f.CreateDirAll(projectPath + "/" + folder); err != nil {
			return err
		}
		time.Sleep(time.Duration(100) * time.Millisecond)
	}
	s.Stop()
	fmt.Println("✓ Task 1: Completed")

	s = spinner.StartNew("Task 2: Processing...")

	for file, tmpl := range config.Files {
		s.SetText(projectPath + "/" + file)

		tmplContent, err := ReadTemplates(f, tmpl)
		if err != nil {
			return err
		}

		if err := f.SaveFile(projectPath+"/"+file, tmplContent); err != nil {
			return err
		}
		time.Sleep(time.Duration(100) * time.Millisecond)
	}
	s.Stop()
	fmt.Println("✓ Task 2: Completed")

	return nil
}

func ReadTemplates(f *files.Files, filePath string) ([]byte, error) {
	data, err := fs.ReadFile(templates, "templates/"+filePath)
	if err != nil {
		return nil, err
	}

	return data, nil
}
