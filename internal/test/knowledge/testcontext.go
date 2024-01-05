package testcontext

import (
	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/internal/test/knowledge/knowledge1"
	"github.com/qiangxue/go-rest-api/internal/test/knowledge/knowledge2"
)

type RuleContext struct {
}

func (b *RuleContext) ExecWhen(knowledgeName string, ruleName string, ruleWhenResult bool, analysis *entity.Analysis) bool {
	knowledge, _ := analysis.AddKnowledge(knowledgeName)
	knowledge.AddRuleFinding(ruleName, ruleWhenResult)
	return ruleWhenResult
}

func (b *RuleContext) ExecThen(knowledgeName string, ruleName string, finding *entity.Finding, analysis *entity.Analysis) {
	knowledge, _ := analysis.GetKnowledge(knowledgeName)
	ruleFinding, _ := knowledge.GetRuleFinding(ruleName)
	findings := ruleFinding.Findings
	*findings = append(*findings, *finding)
}

func (ctx *RuleContext) When_knowledge1_rule1(sta *entity.STA) bool {
	return knowledge1.When_rule1(sta)
}

func (ctx *RuleContext) When_knowledge1_rule2(sta *entity.STA) bool {
	return knowledge1.When_rule2(sta)
}

func (ctx *RuleContext) When_knowledge2_rule1(sta *entity.STA) bool {
	return knowledge2.When_rule1(sta)
}

func (ctx *RuleContext) When_knowledge2_rule2(sta *entity.STA) bool {
	return knowledge2.When_rule2(sta)
}

func (ctx *RuleContext) Then_knowledge1_rule1(sta *entity.STA) *entity.Finding {
	return knowledge1.Then_rule1(sta)
}

func (ctx *RuleContext) Then_knowledge1_rule2(sta *entity.STA) *entity.Finding {
	return knowledge1.Then_rule2(sta)
}

func (ctx *RuleContext) Then_knowledge2_rule1(sta *entity.STA) *entity.Finding {
	return knowledge2.Then_rule1(sta)
}

func (ctx *RuleContext) Then_knowledge2_rule2(sta *entity.STA) *entity.Finding {
	return knowledge2.Then_rule2(sta)
}
