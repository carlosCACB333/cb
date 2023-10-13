package search

import (
	"context"
	"fmt"

	"github.com/carlosCACB333/cb-back/model"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type PostElasticSearch struct {
	client *elasticsearch.TypedClient
}

func NewPostElasticSearch(url string) (*PostElasticSearch, error) {
	client, err := elasticsearch.NewTypedClient(elasticsearch.Config{
		Addresses: []string{url},
	})
	if err != nil {
		return nil, err
	}

	return &PostElasticSearch{
		client: client,
	}, nil
}

func (e *PostElasticSearch) Close() {
	//
}

// func (e *PostElasticSearch) Create(ctx context.Context) error {

// 	_, err := e.client.Indices.Create("post_search").
// 		Request(&create.Request{
// 			Mappings: &types.TypeMapping{
// 				Properties: map[string]types.Property{
// 					"title":   types.NewTextProperty(),
// 					"content": types.NewTextProperty(),
// 					"summary": types.NewTextProperty(),
// 					"category": types.ObjectProperty{
// 						Properties: map[string]types.Property{
// 							"name":   types.NewKeywordProperty(),
// 							"detail": types.NewTextProperty(),
// 						},
// 					},
// 					"tags": types.NestedProperty{
// 						Properties: map[string]types.Property{
// 							"name":   types.NewKeywordProperty(),
// 							"detail": types.NewTextProperty(),
// 						}},
// 				},
// 			},
// 		}).
// 		Do(ctx)

// 	return err
// }

func (e *PostElasticSearch) Index(ctx context.Context, post *model.Post) error {
	_, err := e.client.Index("post_search").
		Id(post.ID).
		Request(post).
		Do(ctx)

	return err
}

func (e *PostElasticSearch) Search(ctx context.Context, query string) ([]*model.Post, error) {
	res, err := e.client.Search().
		Index("post_search").
		Request(&search.Request{
			Query: &types.Query{
				Match: map[string]types.MatchQuery{
					"title":           {Query: query},
					"content":         {Query: query},
					"summary":         {Query: query},
					"category.name":   {Query: query},
					"category.detail": {Query: query},
					"tags.name":       {Query: query},
					"tags.detail":     {Query: query},
				},
			},
		}).
		Do(ctx)
	if err != nil {
		return nil, err
	}

	posts := make([]*model.Post, 0)
	for _, hit := range res.Hits.Hits {
		var post model.Post
		fmt.Println(hit)
		posts = append(posts, &post)
	}
	return posts, nil
}
