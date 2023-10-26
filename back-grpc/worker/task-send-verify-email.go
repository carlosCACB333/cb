package worker

import (
	"context"
	"encoding/json"
	"fmt"

	model "github.com/carlosCACB333/cb-grpc/models"
	"github.com/carlosCACB333/cb-grpc/utils"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const (
	TASK_SEND_VERIFY_EMAIL = "[TASK] SEND_VERIFY_EMAIL"
)

type PayloadVerifyEmail struct {
	Email string `json:"email"`
}

func (d *RedisTaskDistributor) DistributeTaskSendVerifyEmail(
	ctx context.Context,
	payload *PayloadVerifyEmail,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	task := asynq.NewTask(TASK_SEND_VERIFY_EMAIL, jsonPayload, opts...)
	info, err := d.client.EnqueueContext(ctx, task)
	if err != nil {
		return err
	}
	log.Info().
		Str("task", TASK_SEND_VERIFY_EMAIL).
		Str("id", info.ID).
		Str("type", info.Type).
		Bytes("payload", info.Payload).
		Str("queue", info.Queue).
		Int("max retry", info.MaxRetry).
		Msg("task enqueued")
	return nil
}

func (d *RedisTaskProcessor) ProcessTaskSendVerifyEmail(ctx context.Context, t *asynq.Task) error {
	var payload PayloadVerifyEmail
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return fmt.Errorf("could not unmarshal payload: %w", asynq.SkipRetry)
	}

	var user model.User

	if err := d.db.Where("email = ?", payload.Email).First(&user).Error; err != nil {
		// if err == gorm.ErrRecordNotFound {
		// 	return fmt.Errorf("could not find user: %w", asynq.SkipRetry)
		// }
		return err
	}

	v := model.Verification{
		Model: model.Model{
			ID: utils.NewID(),
		},
		UserId: user.ID,
		Code:   utils.NewOtp(),
	}

	if err := d.db.Create(&v).Error; err != nil {
		return err
	}

	// TODO: send email
	subject := "Verify your email"
	body := "Dear <strong>" + user.FirstName + "</strong>,\n\n" +
		"Please verify your email by clicking the link below:\n\n" +
		d.cfg.ServerUrl + "/verify-email?token=" + v.Code + "\n\n" +
		"If you have any questions, just reply to this email—we're always happy to help out.\n\n" +
		"Thanks,\n" +
		"carloscb"

	if err := utils.SendMail(d.cfg, []string{user.Email}, subject, body); err != nil {
		return err
	}

	log.Info().
		Str("task", TASK_SEND_VERIFY_EMAIL).
		Str("type", t.Type()).
		Bytes("payload", t.Payload()).
		Str("email", payload.Email).
		Msg("task processed")

	return nil

}
