package servers

import (
	"context"

	"github.com/carlosCACB333/cb-grpc/libs"
	model "github.com/carlosCACB333/cb-grpc/models"
	"github.com/carlosCACB333/cb-grpc/pb"
	"github.com/carlosCACB333/cb-grpc/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type PostServer struct {
	pb.UnimplementedPostServiceServer
	Config *utils.Config
	DB     *gorm.DB
}

func NewPostServer(cfg *utils.Config, db *gorm.DB) *PostServer {
	return &PostServer{
		Config: cfg,
		DB:     db,
	}
}

func (s *PostServer) Create(ctx context.Context, req *pb.Post) (*pb.PostRes, error) {
	token := utils.GetBererAuth(ctx)
	if token == "" {
		return nil, status.Error(codes.Unauthenticated, "No se ha enviado el token de autenticación")
	}

	userId, err := libs.ValidateToken(s.Config.JWTSecret, token)
	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "El token de autenticación es inválido")
	}

	var user model.User
	if err := s.DB.Where("id = ?", userId).First(&user).Error; err != nil {
		return nil, status.Error(codes.NotFound, "No se ha encontrado el usuario")
	}

	post := pbToPost(req)
	post.ID = utils.NewID()
	post.AuthorId = userId

	if err := utils.ValidateFields(*post); err != nil {
		return nil, err
	}

	if err := s.DB.Create(&post).Error; err != nil {
		return nil, err
	}

	return &pb.PostRes{
		Post: postToPB(post),
	}, nil

}
func (s *PostServer) Read(ctx context.Context, req *pb.PostReq) (*pb.PostRes, error) {
	var post model.Post
	if err := s.DB.Preload("Category").Preload("Tags").Where("id = ?", req.Id).First(&post).Error; err != nil {
		return nil, status.Error(codes.NotFound, "No se ha encontrado el post")
	}

	
	return &pb.PostRes{
		Post: postToPB(&post),
	}, nil

}
func (s *PostServer) Update(ctx context.Context, req *pb.Post) (*pb.PostRes, error) {
	return nil, nil
}
func (s *PostServer) Delete(ctx context.Context, req *pb.PostReq) (*pb.PostRes, error) {
	return nil, nil
}
func (s *PostServer) List(*emptypb.Empty, pb.PostService_ListServer) error {
	return nil
}
func (s *PostServer) AddTag(pb.PostService_AddTagServer) error {
	return nil
}

func postToPB(post *model.Post) *pb.Post {
	pbTags := []*pb.Tag{}
	for _, tag := range post.Tags {
		t := &pb.Tag{
			Id:        tag.ID,
			Name:      tag.Name,
			Icon:      tag.Icon,
			CreatedAt: timestamppb.New(tag.CreatedAt),
			UpdatedAt: timestamppb.New(tag.UpdatedAt),
		}
		pbTags = append(pbTags, t)
	}

	return &pb.Post{
		Id:         post.ID,
		Title:      post.Title,
		Content:    post.Content,
		Banner:     post.Banner,
		Summary:    post.Summary,
		Slug:       post.Slug,
		AuthorId:   post.AuthorId,
		CategoryId: post.CategoryId,
		CreatedAt:  timestamppb.New(post.CreatedAt),
		UpdatedAt:  timestamppb.New(post.UpdatedAt),
		Category: &pb.Category{
			Id:        post.Category.ID,
			Name:      post.Category.Name,
			Icon:      post.Category.Icon,
			Slug:      post.Category.Slug,
			Detail:    post.Category.Detail,
			Img:       post.Category.Img,
			CreatedAt: timestamppb.New(post.Category.CreatedAt),
			UpdatedAt: timestamppb.New(post.Category.UpdatedAt),
		},
		Tags: pbTags,
	}
}

func pbToPost(post *pb.Post) *model.Post {
	postTags := []*model.Tag{}
	for _, tag := range post.Tags {
		t := &model.Tag{
			Model: model.Model{
				ID:        tag.GetId(),
				CreatedAt: tag.GetCreatedAt().AsTime(),
				UpdatedAt: tag.GetUpdatedAt().AsTime(),
			},
			Name: tag.GetName(),
			Icon: tag.GetIcon(),
		}
		postTags = append(postTags, t)
	}

	category := &model.Category{
		Model: model.Model{
			ID:        post.GetCategory().GetId(),
			CreatedAt: post.GetCategory().GetCreatedAt().AsTime(),
			UpdatedAt: post.GetCategory().GetUpdatedAt().AsTime(),
		},
		Name:   post.GetCategory().GetName(),
		Icon:   post.GetCategory().GetIcon(),
		Slug:   utils.Slug(post.GetCategory().GetName()),
		Detail: post.GetCategory().GetDetail(),
		Img:    post.GetCategory().GetImg(),
	}

	if category.ID == "" {
		if post.GetCategoryId() != "" {
			category.ID = post.GetCategoryId()
		} else {
			category.ID = utils.NewID()
		}
	}

	return &model.Post{
		Model: model.Model{
			ID:        post.GetId(),
			CreatedAt: post.GetCreatedAt().AsTime(),
			UpdatedAt: post.GetUpdatedAt().AsTime(),
		},
		Title:      post.GetTitle(),
		Content:    post.GetContent(),
		Category:   category,
		Slug:       utils.Slug(post.GetTitle()),
		Banner:     post.GetBanner(),
		Summary:    post.GetSummary(),
		Tags:       postTags,
		AuthorId:   post.GetAuthorId(),
		CategoryId: post.GetCategoryId(),
	}
}
