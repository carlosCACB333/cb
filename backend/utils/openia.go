package utils

import (
	"cb/libs"
	"context"

	openai "github.com/sashabaranov/go-openai"
)

func GetEmbddingsPDF(texts []string) ([]openai.Embedding, error) {
	oia := libs.OpenIA()
	emb, err := oia.CreateEmbeddings(context.Background(), openai.EmbeddingRequest{
		Input: texts,
		Model: openai.AdaEmbeddingV2,
	})
	if err != nil {
		return nil, err
	}
	return emb.Data, nil
}
