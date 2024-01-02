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

type RuleContext struct {
}

type Rule struct {
	Name string
	Desc string
}

func Get(logger log.Logger) (string, error) {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Dir(filename)

	files, _ := os.ReadDir(dir)

	rulesString := ""
	for _, file := range files {
		if file.Name() == "rule.go" {
			continue
		}
		ruleName := strings.Split(file.Name(), ".")[0]
		r := getRule(ruleName)
		rulesString += r.String()
	}

	return rulesString, nil
}

func getRule(ruleName string) Rule {
	return Rule{
		Name: ruleName,
		Desc: "desc",
	}
}

func (rule *Rule) String() string {
	return fmt.Sprintf(`
rule %s "%s" salience 10 {
    when
	  RuleContext.When_%s(sta)
    then
      RuleContext.PopulateFinding(findings, RuleContext.Then_%s(sta));
      Retract("%s");
}
`, rule.Name, rule.Desc, rule.Name, rule.Name, rule.Name)
}

func (r *RuleContext) PopulateFinding(findings *[]entity.Finding, finding *entity.Finding) {
	*findings = append(*findings, *finding)
	fmt.Println("finding populated")
}
