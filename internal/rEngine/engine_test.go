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
	err := rEngine.Start(logger)
	assert.NoError(t, err)

	err = rEngine.Start(logger)
	assert.Equal(t, "RulesEngine already started", err.Error())

	rEngine.Stop()
}

func TestRuleExecution(t *testing.T) {
	err := rEngine.Start(logger)
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

	err = dataContext.Add("RuleContext", &rules.RuleContext{})
	if err != nil {
		t.Fatal(err)
	}

	err = dataContext.Add("executionMap", &map[string]bool{})
	if err != nil {
		t.Fatal(err)
	}

	findings := &[]entity.Finding{}
	err = dataContext.Add("findings", findings)
	if err != nil {
		t.Fatal(err)
	}

	err = rEngine.RulesEngine.Execute(dataContext, rEngine.KnowledgeBase)
	assert.NoError(t, err)

	assert.NotEmpty(t, findings)
	assert.Equal(t, 2, len(*findings))
}
