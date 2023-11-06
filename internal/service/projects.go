package service

import (
	"context"
	"fmt"
	"fold/internal/config"
	"fold/internal/model"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
)

type Project struct {
	Client  *elasticsearch.Client
	ESIndex string
}

func NewProject(esClient *elasticsearch.Client) (*Project, error) {

	return &Project{
		Client:  esClient,
		ESIndex: config.ESProjectsIndex,
	}, nil
}

func (p *Project) SearchProjectsByUser(ctx context.Context, userID string) ([]model.Project, error) {
	query := fmt.Sprintf(config.SearchProjectsByUserQuery, userID)
	return p.getFromES(query)
}

func (p *Project) SearchProjectsByHashtags(ctx context.Context, hashTags []string) ([]model.Project, error) {
	hashtagsStr := "[" + strings.Join(quoteStrings(hashTags), ",") + "]"
	fmt.Println(hashtagsStr)
	query := fmt.Sprintf(config.SearchProjectsByHashtagsQuery, hashtagsStr)
	return p.getFromES(query)
}

func quoteStrings(strings []string) []string {
	quotedStrings := make([]string, len(strings))
	for i, str := range strings {
		quotedStrings[i] = fmt.Sprintf(`"%s"`, str)
	}
	return quotedStrings
}

func (p *Project) FullTextSearchProjects(ctx context.Context, keyword string) ([]model.Project, error) {
	query := fmt.Sprintf(config.FullTextSearchProjectsQuery, keyword)
	return p.getFromES(query)
}
