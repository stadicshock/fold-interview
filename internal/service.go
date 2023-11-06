package internal

import (
	"context"
	"fold/internal/model"
)

type Service interface {
	SearchProjectsByUser(ctx context.Context, userID string) ([]model.Project, error)
	SearchProjectsByHashtags(ctx context.Context, hashTags []string) ([]model.Project, error)
	FullTextSearchProjects(ctx context.Context, keyword string) ([]model.Project, error)
}
