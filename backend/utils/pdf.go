package utils

import (
	"context"
	"fmt"

	"io"
	"os"
	"time"

	"github.com/tmc/langchaingo/documentloaders"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/textsplitter"
)

func GetContentPDF(s3File io.ReadCloser) ([]schema.Document, error) {

	fileBytes, err := io.ReadAll(s3File)
	if err != nil {
		return nil, err
	}
	fileName := "temp/" + time.Now().Format(time.RFC3339) + ".pdf"
	file, e := os.Create(fileName)
	if e != nil {
		return nil, e
	}

	defer file.Close()
	defer os.Remove(fileName)

	_, e2 := file.Write(fileBytes)

	if e2 != nil {
		return nil, e2
	}
	stats, _ := file.Stat()
	pdf := documentloaders.NewPDF(file, stats.Size())

	doc, err := pdf.LoadAndSplit(context.TODO(), textsplitter.RecursiveCharacter{
		Separators: []string{"\n\n", "\n", " ", ""},
		ChunkSize:  100,
	})
	if err != nil {
		return nil, err
	}
	return doc, nil
}

func GetEmbddingsPDF(texts []string) {
	em := embeddings.BatchTexts(texts, 5)
	fmt.Println(em)
}
