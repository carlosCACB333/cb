package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"
	"github.com/tmc/langchaingo/schema"
)

type Metadata struct {
	Page       int    `json:"page"`
	TotalPages int    `json:"totalPages"`
	Text       string `json:"text"`
}

type Vector struct {
	ID       string    `json:"id"`
	Values   []float32 `json:"values"`
	Metadata Metadata  `json:"metadata"`
}

type VectorUpsert struct {
	Vectors   []Vector `json:"vectors"`
	Namespace string   `json:"namespace"`
}

type VectorQuery struct {
	Vector          []float32 `json:"vector"`
	Namespace       string    `json:"namespace"`
	TopK            int       `json:"topK"`
	IncludeValues   bool      `json:"includeValues"`
	IncludeMetadata bool      `json:"includeMetadata"`
}

type Maches struct {
	ID       string    `json:"id"`
	Score    float32   `json:"score"`
	Values   []float32 `json:"values"`
	Metadata Metadata  `json:"metadata"`
}
type QueryResponse struct {
	Matches   []Maches `json:"matches"`
	Namespace string   `json:"namespace"`
}

func getPineconeCredentials(url string) (string, string) {
	return os.Getenv("PINECONE_URL") + url, os.Getenv("PINECONE_API_KEY")
}
func SaveVectorOnPinacone(docs []schema.Document, embeds []openai.Embedding, namespace string) error {
	url, apiKey := getPineconeCredentials("/vectors/upsert")
	var vectors []Vector
	for i, doc := range docs {
		e := embeds[i]
		vectors = append(vectors, Vector{
			ID:     fmt.Sprintf("%d", e.Index),
			Values: e.Embedding,
			Metadata: Metadata{
				Page:       doc.Metadata["page"].(int),
				TotalPages: doc.Metadata["total_pages"].(int),
				Text:       doc.PageContent,
			},
		})
	}

	vu := VectorUpsert{
		Vectors:   vectors,
		Namespace: namespace,
	}

	jsonBody, _ := json.Marshal(vu)
	body := strings.NewReader(string(jsonBody))

	req, _ := http.NewRequest("POST", url, body)
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Api-Key", apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()
	io.ReadAll(res.Body)

	return nil
}

func GetTopkMachesFromPinecone(embedding []float32, namespace string, topK int) (QueryResponse, error) {
	url, apiKey := getPineconeCredentials("/query")

	vq := VectorQuery{
		Vector:          embedding,
		Namespace:       namespace,
		TopK:            topK,
		IncludeValues:   true,
		IncludeMetadata: true,
	}

	jsonBody, _ := json.Marshal(vq)
	body := strings.NewReader(string(jsonBody))

	req, _ := http.NewRequest("POST", url, body)
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Api-Key", apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return QueryResponse{}, err
	}
	defer res.Body.Close()
	var qr QueryResponse
	resBody, _ := io.ReadAll(res.Body)
	json.Unmarshal(resBody, &qr)
	return qr, nil
}

func DeleteVectorsByNamespace(namespace string) error {
	url, apiKey := getPineconeCredentials("/vectors/delete")

	jsonBody, _ := json.Marshal(map[string]any{
		"namespace": namespace,
		"deleteAll": true,
	})
	body := strings.NewReader(string(jsonBody))

	req, _ := http.NewRequest("DELETE", url, body)
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("Api-Key", apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()
	resBody, _ := io.ReadAll(res.Body)
	fmt.Println("line 149", string(resBody))

	return nil
}
