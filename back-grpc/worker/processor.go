package worker

import (
	"context"

	"github.com/carlosCACB333/cb-grpc/utils"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

const (
	QUEUE_SEND_EMAILS = "send_emails"
)

type TaskProcessor interface {
	ProcessTaskSendVerifyEmail(ctx context.Context, t *asynq.Task) error
	Start() error
}

type RedisTaskProcessor struct {
	server *asynq.Server
	db     *gorm.DB
	cfg    *utils.Config
}

func NewRedisTaskProcessor(cfg *utils.Config, op asynq.RedisClientOpt, db *gorm.DB) TaskProcessor {
	return &RedisTaskProcessor{
		server: asynq.NewServer(op, asynq.Config{
			Concurrency: 10,
			Queues: map[string]int{
				QUEUE_SEND_EMAILS: 1,
			},
			ErrorHandler: asynq.ErrorHandlerFunc(
				func(ctx context.Context, task *asynq.Task, err error) {
					log.Error().
						Str("task", task.Type()).
						Str("id", task.ResultWriter().TaskID()).
						Bytes("payload", task.Payload()).
						Err(err).
						Msg("task error")

				},
			),
		}),
		db:  db,
		cfg: cfg,
	}
}

func (d *RedisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()
	mux.HandleFunc(TASK_SEND_VERIFY_EMAIL, d.ProcessTaskSendVerifyEmail)
	return d.server.Start(mux)
}
