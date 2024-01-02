package rEngine

import (
	"fmt"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/qiangxue/go-rest-api/internal/rEngine/rules"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

var (
	RulesEngine   *engine.GruleEngine
	KnowledgeBase *ast.KnowledgeBase
)

func Stop() {
	RulesEngine = nil
	KnowledgeBase = nil
}

func Start(logger log.Logger) error {
	if RulesEngine != nil {
		return fmt.Errorf("RulesEngine already started")
	}
	if KnowledgeBase != nil {
		return fmt.Errorf("KnowledgeBase already created")
	}

	rulesString, err := rules.Get(logger)
	if err != nil {
		return err
	}

	lib := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(lib)
	err = ruleBuilder.BuildRuleFromResource("Test", "0.1.1", pkg.NewBytesResource([]byte(rulesString)))
	if err != nil {
		return err
	}

	KnowledgeBase, err = lib.NewKnowledgeBaseInstance("Test", "0.1.1")
	if err != nil {
		return err
	}

	RulesEngine = &engine.GruleEngine{MaxCycle: 2}

	return nil
}
