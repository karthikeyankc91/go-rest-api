package ruleengine_test

import (
	"testing"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/internal/ruleengine"
	testcontext "github.com/qiangxue/go-rest-api/internal/test/knowledge"
	"github.com/qiangxue/go-rest-api/pkg/log"
	"github.com/stretchr/testify/assert"
)

var (
	logger              = log.New()
	testKnowledgeFolder = "../test/knowledge/"
)

func TestEngineStartup(t *testing.T) {
	rulesEngine, err := ruleengine.InitializeEngine(logger, testKnowledgeFolder)
	assert.NoError(t, err)
	assert.NotNil(t, rulesEngine.RulesEngine)
	assert.NotNil(t, rulesEngine.AllKnowledge)

	assert.NotNil(t, rulesEngine.KnowledgeMap)
	assert.Equal(t, 2, len(*rulesEngine.KnowledgeMap))
	assert.Equal(t, "knowledge1", (*rulesEngine.KnowledgeMap)["knowledge1"].KnowledgeName)
	assert.Equal(t, "knowledge2", (*rulesEngine.KnowledgeMap)["knowledge2"].KnowledgeName)

	assert.Equal(t, 2, len(*((*rulesEngine.KnowledgeMap)["knowledge1"]).RulesMap))
	assert.Equal(t, 2, len(*((*rulesEngine.KnowledgeMap)["knowledge2"]).RulesMap))
	assert.Equal(t, "rule1", (*((*rulesEngine.KnowledgeMap)["knowledge1"]).RulesMap)["rule1"].RuleName)
	assert.Equal(t, "rule2", (*((*rulesEngine.KnowledgeMap)["knowledge1"]).RulesMap)["rule2"].RuleName)
	assert.Equal(t, "rule1", (*((*rulesEngine.KnowledgeMap)["knowledge2"]).RulesMap)["rule1"].RuleName)
	assert.Equal(t, "rule2", (*((*rulesEngine.KnowledgeMap)["knowledge2"]).RulesMap)["rule2"].RuleName)
}

func TestEngineExecution(t *testing.T) {
	rulesEngine, err := ruleengine.InitializeEngine(logger, testKnowledgeFolder)
	assert.NoError(t, err)
	assert.NotNil(t, rulesEngine.RulesEngine)
	assert.NotNil(t, rulesEngine.AllKnowledge)
	assert.NotNil(t, rulesEngine.KnowledgeMap)

	sta := &entity.STA{
		StaId: "sta1",
	}

	dataContext := ast.NewDataContext()
	err = dataContext.Add("sta", sta)
	assert.NoError(t, err)

	err = dataContext.Add("ctx", &testcontext.RuleContext{})
	assert.NoError(t, err)

	analysis := &entity.Analysis{}
	analysis.Init(sta)
	err = dataContext.Add("analysis", analysis)
	assert.NoError(t, err)

	err = rulesEngine.RulesEngine.Execute(dataContext, rulesEngine.AllKnowledge)
	assert.NoError(t, err)

	kMap := *analysis.KnowledgeMap
	assert.Equal(t, 2, len(kMap))
	assert.Equal(t, "knowledge1", (kMap)["knowledge1"].KnowledgeName)
	assert.Equal(t, "knowledge2", kMap["knowledge2"].KnowledgeName)

	k1RuleFindingsMap := *(kMap["knowledge1"]).RuleFindingsMap
	k2RuleFindingsMap := *(kMap["knowledge2"]).RuleFindingsMap

	assert.Equal(t, 2, len(k1RuleFindingsMap))
	assert.Equal(t, 2, len(k2RuleFindingsMap))

	k1RFRule1 := k1RuleFindingsMap["rule1"]
	k1RFRule2 := k1RuleFindingsMap["rule2"]
	k2RFRule1 := k2RuleFindingsMap["rule1"]
	k2RFRule2 := k2RuleFindingsMap["rule2"]

	assert.Equal(t, "rule1", k1RFRule1.RuleName)
	assert.Equal(t, "rule2", k1RFRule2.RuleName)
	assert.Equal(t, "rule1", k2RFRule1.RuleName)
	assert.Equal(t, "rule2", k2RFRule2.RuleName)

	assert.Equal(t, 1, len(*k1RFRule1.Findings))
	assert.Equal(t, 1, len(*k1RFRule2.Findings))
	assert.Equal(t, 1, len(*k2RFRule1.Findings))
	assert.Equal(t, 1, len(*k2RFRule2.Findings))

	assert.Equal(t, "executed then for knowledge1 rule1", (*k1RFRule1.Findings)[0].Desc)
	assert.Equal(t, "executed then for knowledge1 rule2", (*k1RFRule2.Findings)[0].Desc)
	assert.Equal(t, "executed then for knowledge2 rule1", (*k2RFRule1.Findings)[0].Desc)
	assert.Equal(t, "executed then for knowledge2 rule2", (*k2RFRule2.Findings)[0].Desc)

}
