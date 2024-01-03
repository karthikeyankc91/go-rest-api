package rules

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/qiangxue/go-rest-api/internal/entity"
	"github.com/qiangxue/go-rest-api/pkg/log"
)

type RulesMap map[string]*Rule

type RuleContext struct {
	RulesMap *RulesMap
	Analysis *entity.Analysis
}

type Rule struct {
	Name    string
	Desc    string
	RuleStr string
}

func Initialize(logger log.Logger) (*RulesMap, error) {
	// if RulesMap != nil {
	// 	return fmt.Errorf("RulesMap already created")
	// }

	rulesMap := make(RulesMap)

	_, filename, _, _ := runtime.Caller(0)
	dir := path.Dir(filename)

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), "_rule.go") {
			continue
		}
		ruleName := strings.TrimSuffix(file.Name(), "_rule.go")
		rule := &Rule{
			Name: ruleName,
			Desc: "desc",
		}
		rule.RuleStr = rule.toString()
		rulesMap[ruleName] = rule
	}
	return &rulesMap, nil
}

func (rulesMap *RulesMap) GetAll() string {
	var rulesStr string
	for ruleName := range *rulesMap {
		rulesStr += (*rulesMap)[ruleName].RuleStr
	}
	return rulesStr
}

func (rulesMap *RulesMap) Filter(ruleNames []string) (string, error) {
	var rulesStr string
	for _, ruleName := range ruleNames {
		if rule, ok := (*rulesMap)[ruleName]; ok {
			rulesStr += rule.RuleStr
		} else {
			return "", fmt.Errorf("rule %s not found", ruleName)
		}
	}
	return rulesStr, nil
}

func (rule *Rule) toString() string {
	return fmt.Sprintf(`
rule %s "%s" salience 10 {
    when
      RuleContext.ExecWhen("%s", RuleContext.When_%s(sta))
    then
      RuleContext.ExecThen("%s", RuleContext.Then_%s(sta));
      Retract("%s");
}
`, rule.Name, rule.Desc, rule.Name, rule.Name, rule.Name, rule.Name, rule.Name)
}

func (RuleContext *RuleContext) ExecWhen(ruleName string, ruleWhenResult bool) bool {
	// can remove this as when status is already populated in finding
	// but other metadata from sta also can be populated here
	RuleContext.InitFinding(ruleName, ruleWhenResult)
	return ruleWhenResult
}

func (RuleContext *RuleContext) ExecThen(ruleName string, rulefinding *entity.RuleFinding) {
	RuleContext.AddRuleFinding(ruleName, rulefinding)
}

func (RuleContext *RuleContext) InitAnalysis(sta *entity.STA) {
	RuleContext.Analysis = &entity.Analysis{
		ID:          entity.GenerateID(),
		StaId:       sta.StaId,
		FindingsMap: make(map[string]entity.Finding),
	}
}

func (RuleContext *RuleContext) InitFinding(ruleName string, ruleWhenResult bool) bool {
	finding := RuleContext.Analysis.FindingsMap[ruleName]
	if finding.RuleName == "" {
		finding = entity.Finding{
			RuleName:     ruleName,
			WhenStatus:   ruleWhenResult,
			RuleFindings: []entity.RuleFinding{},
		}
		RuleContext.Analysis.FindingsMap[ruleName] = finding

		// Initialized the finding
		return true
	}
	return false
}

func (RuleContext *RuleContext) AddRuleFinding(ruleName string, rulefinding *entity.RuleFinding) {
	finding := RuleContext.Analysis.FindingsMap[ruleName]
	finding.RuleFindings = append(finding.RuleFindings, *rulefinding)
	RuleContext.Analysis.FindingsMap[ruleName] = finding
}
