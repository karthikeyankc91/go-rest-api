package knowledge

import (
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
	"github.com/qiangxue/go-rest-api/pkg/log"

	"github.com/qiangxue/go-rest-api/internal/rEngine/rules"
)

type Knowledge struct {
	KnowlodgeBase *ast.KnowledgeBase
	RulesMap      *rules.RulesMap
}

func Initialize(logger log.Logger) (*Knowledge, error) {
	rulesMap, err := rules.Initialize(logger)
	if err != nil {
		return nil, err
	}

	rulesString := rulesMap.GetAll()

	lib := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(lib)

	err = ruleBuilder.BuildRuleFromResource("Test", "0.1.1", pkg.NewBytesResource([]byte(rulesString)))
	if err != nil {
		return nil, err
	}

	allRulesKnowledge, err := lib.NewKnowledgeBaseInstance("Test", "0.1.1")
	if err != nil {
		return nil, err
	}

	return &Knowledge{
		KnowlodgeBase: allRulesKnowledge,
		RulesMap:      rulesMap,
	}, nil
}

func (k *Knowledge) BuildKnowledge(ruleNames []string) (*ast.KnowledgeBase, error) {
	lib := ast.NewKnowledgeLibrary()
	ruleBuilder := builder.NewRuleBuilder(lib)

	ruleString, err := k.RulesMap.Filter(ruleNames)
	if err != nil {
		return nil, err
	}

	err = ruleBuilder.BuildRuleFromResource("CustomRules", "", pkg.NewBytesResource([]byte(ruleString)))
	if err != nil {
		return nil, err
	}

	knowledge, err := lib.NewKnowledgeBaseInstance("CustomKnowledge", "")
	if err != nil {
		return nil, err
	}

	return knowledge, nil
}
