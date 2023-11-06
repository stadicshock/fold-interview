package service

import (
	"encoding/json"
	"errors"
	"fold/internal/model"
	"log"
	"strings"
)

func (p *Project) getFromES(query string) ([]model.Project, error) {

	res, err := p.Client.Search(
		p.Client.Search.WithIndex(p.ESIndex),
		p.Client.Search.WithBody(strings.NewReader(query)),
	)

	if err != nil {
		log.Println("Error from ES Search:", err)
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Println("Error in ES response", res)
		return nil, errors.New("error in ES response")
	}

	var ESData model.ESData
	if err := json.NewDecoder(res.Body).Decode(&ESData); err != nil {
		log.Println("Error while decoding:", err)
		return nil, err
	}

	hits := ESData.Hits.Hits

	var projects []model.Project
	for _, hit := range hits {
		projects = append(projects, hit.Source)
	}
	return projects, nil
}
