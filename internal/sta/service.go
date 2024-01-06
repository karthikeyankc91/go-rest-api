package sta

import (
	"context"
	"fmt"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/internal/ruleengine"
	"github.com/qiangxue/go-rest-api/internal/showtech"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

// Service encapsulates usecase logic for stas.
type Service interface {
	Get(ctx context.Context, id string) (AnalysisOutput, error)
	Query(ctx context.Context, offset, limit int) ([]AnalysisOutput, error)
	Count(ctx context.Context) (int, error)
	Analyze(ctx context.Context, id string) (AnalysisOutput, error)
}

// Showtech represents the data about an sta.
type AnalysisOutput struct {
	entity.Analysis
}

type service struct {
	repo         Repository
	showtechRepo showtech.Repository
	logger       log.Logger
	rulesEngine  *ruleengine.RuleEngine
}

// NewService creates a new sta service.
func NewService(repo Repository, showtechRepo showtech.Repository, rulesEngine *ruleengine.RuleEngine, logger log.Logger) Service {
	return service{repo, showtechRepo, logger, rulesEngine}
}

// Get returns the album with the specified the album ID.
func (s service) Get(ctx context.Context, id string) (AnalysisOutput, error) {
	album, err := s.repo.Get(ctx, id)
	if err != nil {
		return AnalysisOutput{}, err
	}
	return AnalysisOutput{album}, nil
}

// Count returns the number of albums.
func (s service) Count(ctx context.Context) (int, error) {
	return s.repo.Count(ctx)
}

// Query returns the albums with the specified offset and limit.
func (s service) Query(ctx context.Context, offset, limit int) ([]AnalysisOutput, error) {
	items, err := s.repo.Query(ctx, offset, limit)
	if err != nil {
		return nil, err
	}
	result := []AnalysisOutput{}
	for _, item := range items {
		result = append(result, AnalysisOutput{item})
	}
	return result, nil
}

func (s service) Analyze(context context.Context, id string) (AnalysisOutput, error) {
	a, err := s.repo.GetByStaid(context, id)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return AnalysisOutput{a}, nil
	}

	if a.Staid != "" {
		return AnalysisOutput{}, fmt.Errorf("analysis already exists")
	}

	showtechs, err := s.showtechRepo.Get(context, id)
	if err != nil {
		return AnalysisOutput{}, err
	}

	sta, err := showtechs.Unmarshal()
	if err != nil {
		return AnalysisOutput{}, err
	}

	if err := sta.Validate(); err != nil {
		return AnalysisOutput{}, err
	}

	dataContext := ast.NewDataContext()
	err = dataContext.Add("sta", sta)
	if err != nil {
		return AnalysisOutput{}, err
	}

	err = dataContext.Add("ctx", &ruleengine.RuleContext{})
	if err != nil {
		return AnalysisOutput{}, err
	}

	analysisData := entity.AnalysisData{}
	analysisData.Init(sta)
	err = dataContext.Add("analysis", &analysisData)
	if err != nil {
		return AnalysisOutput{}, err
	}

	err = s.rulesEngine.RulesEngine.Execute(dataContext, s.rulesEngine.AllKnowledge)
	if err != nil {
		return AnalysisOutput{}, err
	}

	a, err = analysisData.MarshalAnalysis()
	if err != nil {
		return AnalysisOutput{}, err
	}
	err = s.repo.Create(context, a)

	if err != nil {
		return AnalysisOutput{}, err
	}
	return s.Get(context, a.ID)
}
