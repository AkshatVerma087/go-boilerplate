package job

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hibiken/asynq"
)

const (
	TaskWelcome = "email:welcome"
)

type WelcomeEmailPayload struct {
	To        string `json:"to"`
	FirstName string `json:"first_name"`
}

func NewWelcomeEmailTask(to string, firstName string) (*asynq.Task, error) {
	payload, err := json.Marshal(WelcomeEmailPayload{
		To:        to,
		FirstName: firstName,
	})

	if err != nil {
		return nil, err
	}

	return asynq.NewTask(TaskWelcome, payload,
		asynq.MaxRetry(3),
		asynq.Queue("default"),
		asynq.Timeout(30*time.Second)), nil
}

func (j *JobService) HandleWelcomeEmailTask(ctx context.Context, task *asynq.Task) error {
	var payload WelcomeEmailPayload
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("decode welcome email task payload: %w", err)
	}

	j.logger.Info().
		Str("to", payload.To).
		Str("first_name", payload.FirstName).
		Msg("processed welcome email task")

	return nil
}

