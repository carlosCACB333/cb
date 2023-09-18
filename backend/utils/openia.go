package utils

import (
	"context"

	"github.com/tmc/langchaingo/llms/openai"
)

func getEmbeddings(text []string) [][]float64 {
	llm, err := openai.New()
	if err != nil {
		panic(err)
	}

	embeddings, err := llm.CreateEmbedding(context.Background(), text)
	if err != nil {
		panic(err)
	}
	return embeddings

}
