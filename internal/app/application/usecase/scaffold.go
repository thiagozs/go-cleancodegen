package usecase

import (
	"fmt"
	"time"

	"github.com/lucasepe/spinner"

	"github.com/thiagozs/go-cleancodegen/internal/app/config"
	"github.com/thiagozs/go-cleancodegen/pkg/utils"
	"github.com/thiagozs/go-utils/files"
)

func CreateScaffold() error {

	if config.Cfg.Path == "" {
		config.Cfg.Path = "/tmp"
	}

	if config.Cfg.Project == "" {
		config.Cfg.Project = "project"
	}

	config.Cfg.Project = utils.RemovePrefixSuffix(config.Cfg.Project, "/")

	projectPath := config.Cfg.Path + "/" + config.Cfg.Project

	if files.FileExists(projectPath) {
		return fmt.Errorf("path %s not exists", config.Cfg.Path)
	}

	s := spinner.StartNew("Task 1: Processing...")
	for _, folder := range config.Folders {
		s.SetText(projectPath + "/" + folder)
		if err := files.MkdirAll(projectPath + "/" + folder); err != nil {
			return err
		}
		time.Sleep(time.Duration(100) * time.Millisecond)
	}
	s.Stop()
	fmt.Println("✓ Task 1: Completed")

	s = spinner.StartNew("Task 2: Processing...")

	for _, file := range config.Files {
		s.SetText(projectPath + "/" + file)
		if err := files.WriteFile(projectPath+"/"+file, []byte{}); err != nil {
			return err
		}
		time.Sleep(time.Duration(100) * time.Millisecond)
	}
	s.Stop()
	fmt.Println("✓ Task 2: Completed")

	return nil
}
