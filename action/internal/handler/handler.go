package handler

import (
	"context"
	"embed"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/urfave/cli/v3"
)

//go:embed templates/*
var templateFS embed.FS

func Action(_ context.Context, command *cli.Command) error {
	t, parseFSErr := template.ParseFS(templateFS, "**/*")
	if parseFSErr != nil {
		return parseFSErr
	}

	packagePath := command.String("path")
	_, packageName := path.Split(packagePath)

	packageName = strings.ReplaceAll(strings.ToLower(packageName), "_", "")
	packageName = strings.ReplaceAll(strings.ToLower(packageName), "-", "")

	handlerFileName := path.Join(packagePath, packageName+".go")
	handlerOptionFileName := path.Join(packagePath, packageName+"option.go")

	const FilePerm = 0644

	fileOpenFlag := os.O_CREATE | os.O_EXCL | os.O_WRONLY
	if command.Bool("force") {
		fileOpenFlag |= os.O_TRUNC
	}

	handlerFile, handlerFileOpenErr := os.OpenFile(handlerFileName, fileOpenFlag, FilePerm)
	if handlerFileOpenErr != nil {
		return handlerFileOpenErr
	}

	defer func() { _ = handlerFile.Close() }()

	handlerOptionFile, handlerOptionFileErr := os.OpenFile(handlerOptionFileName, fileOpenFlag, FilePerm)
	if handlerOptionFileErr != nil {
		return handlerOptionFileErr
	}

	defer func() { _ = handlerOptionFile.Close() }()

	var templateData struct {
		Package string
	}

	templateData.Package = packageName

	if err := t.ExecuteTemplate(handlerFile, "handler.go.tmpl", templateData); err != nil {
		return err
	}

	if err := t.ExecuteTemplate(handlerOptionFile, "handleroption.go.tmpl", templateData); err != nil {
		return err
	}

	return nil
}
