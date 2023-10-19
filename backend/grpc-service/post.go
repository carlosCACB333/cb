package grpcServer

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/carlosCACB333/cb-back/model"
	pb "github.com/carlosCACB333/cb-back/proto"
	"github.com/carlosCACB333/cb-back/server"
	"github.com/carlosCACB333/cb-back/util"
	"github.com/google/uuid"
)

type PostServiceServer struct {
	*server.Server
	pb.UnimplementedPostServiceServer
}

func NewPostServiceServer(s *server.Server) *PostServiceServer {
	return &PostServiceServer{
		Server: s,
	}
}

func (p *PostServiceServer) Read(ctx context.Context, req *pb.PostReq) (*pb.PostRes, error) {
	var post model.Post
	if err := p.DB().Preload("Tags").Preload("Category").Where("id = ?", req.GetId()).First(&post).Error; err != nil {
		return nil, err
	}

	tags := make([]*pb.Tag, len(post.Tags))
	for i, tag := range post.Tags {
		tags[i] = &pb.Tag{
			Id:   tag.ID,
			Name: tag.Name,
			Icon: tag.Icon,
		}
	}
	category := &pb.Category{
		Id:        post.Category.ID,
		Slug:      post.Category.Slug,
		Name:      post.Category.Name,
		Detail:    post.Category.Detail,
		Icon:      post.Category.Icon,
		Img:       post.Category.Img,
		CreatedAt: post.Category.CreatedAt.Unix(),
		UpdatedAt: post.Category.UpdatedAt.Unix(),
	}

	return &pb.PostRes{
		Post: &pb.Post{
			Id:         post.ID,
			Title:      post.Title,
			Content:    post.Content,
			CategoryId: post.CategoryId,
			Slug:       post.Slug,
			Summary:    post.Summary,
			Banner:     post.Banner,
			AuthorId:   post.AuthorId,
			Tags:       tags,
			CreatedAt:  post.CreatedAt.Unix(),
			UpdatedAt:  post.UpdatedAt.Unix(),
			Category:   category,
		},
	}, nil

}

func (ps *PostServiceServer) Create(ctx context.Context, p *pb.Post) (*pb.PostRes, error) {

	tags := make([]*model.Tag, len(p.Tags))
	for i, tag := range p.Tags {
		t := model.Tag{
			Name: tag.Name,
			Icon: tag.Icon,
		}
		if t.ID == "" {
			t.ID = uuid.New().String()
			err := util.ValidateFields(t)
			if err != nil {
				return nil, util.NewError(400, "Los datos de las etiquetas son erróneas", err)
			}

		}
		tags[i] = &t
	}

	category := model.Category{
		Model:  model.Model{ID: p.Category.Id},
		Slug:   p.Category.Slug,
		Name:   p.Category.Name,
		Icon:   p.Category.Icon,
		Detail: p.Category.Detail,
		Img:    p.Category.Img,
	}
	if category.ID == "" {
		category.ID = uuid.New().String()
		err := util.ValidateFields(category)
		fmt.Println(err)
		if err != nil {
			return nil, util.NewError(400, "Los datos de la categoría son erróneas", err)
		}
	}

	post := model.Post{
		Model:      model.Model{ID: uuid.New().String()},
		Title:      p.Title,
		Content:    p.Content,
		CategoryId: p.CategoryId,
		Slug:       p.Slug,
		Summary:    p.Summary,
		Banner:     p.Banner,
		AuthorId:   p.AuthorId,
		Tags:       tags,
		Category:   &category,
	}

	err := util.ValidateFields(post)
	if err != nil {
		return nil, util.NewError(400, "Los datos del post son erróneas", err)
	}

	if err := ps.DB().Create(&post).Error; err != nil {
		return nil, util.NewError(400, "Error al crear el post", err)
	}

	return &pb.PostRes{
		Post: &pb.Post{
			Id:         post.ID,
			Title:      post.Title,
			Content:    post.Content,
			CategoryId: post.CategoryId,
			Slug:       post.Slug,
			Summary:    post.Summary,
			Banner:     post.Banner,
			AuthorId:   post.AuthorId,
			Tags:       p.Tags,
		},
	}, nil

}

func (p *PostServiceServer) Update(context.Context, *pb.Post) (*pb.PostRes, error) {
	return nil, nil
}
func (p *PostServiceServer) Delete(context.Context, *pb.PostReq) (*pb.PostRes, error) {
	return nil, nil
}

func (p *PostServiceServer) AddTag(stream pb.PostService_AddTagServer) error {
	var success int32 = 0
	var failed int32 = 0
	for {
		tagReq, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AddTagRes{
				Success: success,
				Failed:  failed,
				Status:  "COMPLETED",
			})
		}
		if err != nil {
			log.Fatalf("Error reading stream: %v", err)
			return err
		}

		postId := tagReq.GetPostId()
		tagID := tagReq.GetTagId()

		if postId == "" || tagID == "" {
			failed++
			continue
		}

		post := model.Post{}
		if err := p.DB().Preload("Tags").Where("id = ?", postId).First(&post).Error; err != nil {
			return err
		}

		post.Tags = append(post.Tags, &model.Tag{Model: model.Model{ID: tagID}})

		if err := p.DB().Model(&post).Association("Tags").Replace(post.Tags); err != nil {
			failed++
			continue
		}

		success++

	}

}

func (p *PostServiceServer) List(_ *pb.PostReq, stream pb.PostService_ListServer) error {

	var posts []*model.Post
	if err := p.DB().Preload("Tags").Preload("Category").Find(&posts).Error; err != nil {
		return err
	}

	for _, post := range posts {
		tags := make([]*pb.Tag, len(post.Tags))
		for i, tag := range post.Tags {
			tags[i] = &pb.Tag{
				Id:   tag.ID,
				Name: tag.Name,
				Icon: tag.Icon,
			}
		}
		category := &pb.Category{
			Id:        post.Category.ID,
			Slug:      post.Category.Slug,
			Name:      post.Category.Name,
			Detail:    post.Category.Detail,
			Icon:      post.Category.Icon,
			Img:       post.Category.Img,
			CreatedAt: post.Category.CreatedAt.Unix(),
			UpdatedAt: post.Category.UpdatedAt.Unix(),
		}

		pbPost := &pb.PostRes{
			Post: &pb.Post{
				Id:         post.ID,
				Title:      post.Title,
				Content:    post.Content,
				CategoryId: post.CategoryId,
				Slug:       post.Slug,
				Summary:    post.Summary,
				Banner:     post.Banner,
				AuthorId:   post.AuthorId,
				Tags:       tags,
				CreatedAt:  post.CreatedAt.Unix(),
				UpdatedAt:  post.UpdatedAt.Unix(),
				Category:   category,
			},
		}

		if err := stream.Send(pbPost); err != nil {
			return err
		}
	}

	return nil

}
