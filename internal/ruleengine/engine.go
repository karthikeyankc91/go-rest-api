package ruleengine

import (
	"path"
	"path/filepath"
	"runtime"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

var (
	_, currentFile, _, _ = runtime.Caller(0)
)

type RuleEngine struct {
	RulesEngine  *engine.GruleEngine
	AllKnowledge *ast.KnowledgeBase
	Knowledgemap *Knowledgemap
}

func InitializeEngine(logger log.Logger, rulesFolder string) (*RuleEngine, error) {
	allKnowledge, knowledgeMap, err := CreateKnowledge(logger, path.Join(filepath.Dir(currentFile), rulesFolder))
	if err != nil {
		return nil, err
	}

	return &RuleEngine{
		RulesEngine:  engine.NewGruleEngine(),
		AllKnowledge: allKnowledge,
		Knowledgemap: knowledgeMap,
	}, nil
}
