package libs

import (
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func OpenIA() *openai.Client {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	return client
}
