package sta

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/internal/ruleengine"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

// Service encapsulates usecase logic for stas.
type Service interface {
	Analyze(ctx context.Context, input CreateSTARequest) (*entity.Analysis, error)
}

// STA represents the data about an sta.
type STA struct {
	entity.STA
}

// CreateSTARequest represents an sta creation request.
type CreateSTARequest struct {
	StaId string `json:"staId"`

	Property1 string `json:"property1"`
	Property2 string `json:"property2"`
	Property3 string `json:"property3"`
}

// Validate validates the CreateSTARequest fields.
func (m CreateSTARequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.StaId, validation.Required, validation.Length(0, 128)),
	)
}

type service struct {
	repo        Repository
	logger      log.Logger
	rulesEngine *ruleengine.RuleEngine
}

// NewService creates a new sta service.
func NewService(repo Repository, rulesEngine *ruleengine.RuleEngine, logger log.Logger) Service {
	return service{repo, logger, rulesEngine}
}

func (s service) Analyze(context context.Context, request CreateSTARequest) (*entity.Analysis, error) {
	var analysis = &entity.Analysis{}

	if err := request.Validate(); err != nil {
		return analysis, err
	}

	sta := &entity.STA{
		StaId:     request.StaId,
		Property1: request.Property1,
		Property2: request.Property2,
		Property3: request.Property3,
	}

	dataContext := ast.NewDataContext()
	err := dataContext.Add("sta", sta)
	if err != nil {
		return analysis, err
	}

	err = dataContext.Add("ctx", &ruleengine.RuleContext{})
	if err != nil {
		return analysis, err
	}

	analysis.Init(sta)
	err = dataContext.Add("analysis", analysis)
	if err != nil {
		return analysis, err
	}

	err = s.rulesEngine.RulesEngine.Execute(dataContext, s.rulesEngine.AllKnowledge)
	if err != nil {
		return analysis, err
	}

	// err = s.repo.Create(ctx, *ruleCtx.Analysis)
	// if err != nil {
	// 	return analysis, err
	// }

	// return s.Get(ctx, id)
	// return *RuleContext.Analysis, err
	return analysis, nil
}
