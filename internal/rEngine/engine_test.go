package rEngine_test

import (
	"testing"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/internal/rEngine"
	"github.com/qiangxue/go-rest-api/internal/rEngine/rules"
	"github.com/qiangxue/go-rest-api/pkg/log"
	"github.com/stretchr/testify/assert"
)

var (
	logger = log.New()
)

func TestStart(t *testing.T) {
	rulesEngine, err := rEngine.Initialize(logger)
	assert.NoError(t, err)

	assert.NotNil(t, rulesEngine.RulesEngine)
	assert.NotNil(t, rulesEngine.Knowledge)
}

func TestRuleExecution(t *testing.T) {
	rulesEngine, err := rEngine.Initialize(logger)
	assert.NoError(t, err)

	sta := &entity.STA{
		Property1: "value1",
		Property2: "value1",
	}

	dataContext := ast.NewDataContext()
	err = dataContext.Add("sta", sta)
	if err != nil {
		t.Fatal(err)
	}

	RuleContext := &rules.RuleContext{
		RulesMap: rulesEngine.Knowledge.RulesMap,
	}
	RuleContext.InitAnalysis(sta)
	err = dataContext.Add("RuleContext", RuleContext)
	if err != nil {
		t.Fatal(err)
	}

	err = rulesEngine.RulesEngine.Execute(dataContext, rulesEngine.Knowledge.KnowlodgeBase)
	assert.NoError(t, err)

	assert.NotEmpty(t, RuleContext.Analysis.FindingsMap)
	assert.Equal(t, 2, len(RuleContext.Analysis.FindingsMap))
}
