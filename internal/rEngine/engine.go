package rEngine

import (
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/qiangxue/go-rest-api/internal/rEngine/knowledge"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

type RuleEngine struct {
	RulesEngine *engine.GruleEngine
	Knowledge   *knowledge.Knowledge
}

func Initialize(logger log.Logger) (*RuleEngine, error) {
	knowledge, err := knowledge.Initialize(logger)
	if err != nil {
		return nil, err
	}

	return &RuleEngine{
		RulesEngine: engine.NewGruleEngine(),
		Knowledge:   knowledge,
	}, nil
}
