package servers

import (
	"context"

	"github.com/carlosCACB333/cb-grpc/libs"
	model "github.com/carlosCACB333/cb-grpc/models"
	"github.com/carlosCACB333/cb-grpc/pb"
	"github.com/carlosCACB333/cb-grpc/utils"
	"github.com/carlosCACB333/cb-grpc/worker"
	"github.com/hibiken/asynq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type AuthServer struct {
	pb.UnimplementedAuthServiceServer
	Config          *utils.Config
	DB              *gorm.DB
	TaskDistributor worker.TaskDistributor
}

func NewAuthServer(
	cfg *utils.Config,
	db *gorm.DB,
	distributor worker.TaskDistributor,
) *AuthServer {

	return &AuthServer{
		Config:          cfg,
		DB:              db,
		TaskDistributor: distributor,
	}
}

func (s *AuthServer) Signup(c context.Context, req *pb.SignupReq) (*pb.SignupRes, error) {

	user := model.User{
		Email:    utils.NormalizeEmail(req.GetEmail()),
		Password: libs.HashPassword(req.GetPassword()),
		Username: req.GetUsername(),
		Model: model.Model{
			ID: utils.NewID(),
		},
		FirstName: req.GetFirstName(),
		LastName:  req.GetLastName(),
		Gender:    req.GetGender(),
		Photo:     req.GetPhoto(),
		Phone:     req.GetPhone(),
		Status:    "active",
	}

	err := utils.ValidateFields(user)

	if err != nil {
		return nil, err
	}

	err = s.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		payload := &worker.PayloadVerifyEmail{
			Email: user.Email,
		}

		taskOpts := []asynq.Option{
			asynq.MaxRetry(10),
			// asynq.ProcessIn(10 * time.Second),
			asynq.Queue(worker.QUEUE_SEND_EMAILS),
		}
		err = s.TaskDistributor.DistributeTaskSendVerifyEmail(c, payload, taskOpts...)
		if err != nil {
			return err
		}

		return nil
	},
	)

	if err != nil {
		return nil, status.Errorf(codes.DataLoss, "No se pudo crear la cuenta")
	}

	return &pb.SignupRes{
		User: userToPB(&user),
	}, nil

}

func (s *AuthServer) Login(c context.Context, req *pb.LoginReq) (*pb.LoginRes, error) {

	var user model.User
	if err := s.DB.Where("email = ?", utils.NormalizeEmail(req.GetEmail())).First(&user).Error; err != nil {
		return nil, status.Errorf(codes.NotFound, "No se encontro el usuario")
	}

	if !libs.CheckPassword(req.GetPassword(), user.Password) {
		return nil, status.Errorf(codes.Unauthenticated, "Credenciales incorrectas")
	}

	token, exp, err := libs.GenerateToken(s.Config.JWTSecret, user.ID, s.Config.AccessTokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "No se pudo generar el token")
	}

	metada := utils.ExtractMetadata(c)

	session := model.Session{
		UserID:    user.ID,
		Token:     token,
		UserAgent: metada.UserAgent,
		IP:        metada.ClientIP,
		ExpiresAt: exp,
		IsBlocked: false,
		Model: model.Model{
			ID: utils.NewID(),
		},
	}

	if err := s.DB.Create(&session).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "No se pudo crear la sesion")
	}

	return &pb.LoginRes{
		User: userToPB(&user),
		Session: &pb.Session{
			Id:        session.ID,
			UserId:    session.UserID,
			Token:     session.Token,
			UserAgent: session.UserAgent,
			Ip:        session.IP,
			ExpiresAt: timestamppb.New(session.ExpiresAt),
			IsBlocked: session.IsBlocked,
			CreatedAt: timestamppb.New(session.CreatedAt),
			UpdatedAt: timestamppb.New(session.UpdatedAt),
		},
	}, nil
}

func userToPB(user *model.User) *pb.User {
	return &pb.User{
		Id:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Gender:    user.Gender,
		Photo:     user.Photo,
		Phone:     user.Phone,
		Status:    user.Status,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}
