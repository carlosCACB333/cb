package utils

import (
	"context"
	"mime/multipart"

	"github.com/tmc/langchaingo/documentloaders"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/textsplitter"
)

func GetContentPDF(fileHeader *multipart.FileHeader) ([]schema.Document, error) {

	file, err := fileHeader.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()
	pdf := documentloaders.NewPDF(file, fileHeader.Size)
	doc, err := pdf.LoadAndSplit(context.TODO(), textsplitter.RecursiveCharacter{
		Separators: []string{"\n\n", "\n", " ", ""},
		ChunkSize:  100,
	})
	if err != nil {
		return nil, err
	}

	return doc, nil
}
