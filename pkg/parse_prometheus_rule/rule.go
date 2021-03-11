/*
@Description:
@Author: fei.wang
@Date: 2021/03/xx
*/

package parse_prometheus_rule

type RuleGroups struct {
	Groups []RuleGroup `yaml:"groups"`
}

type RuleGroup struct {
	Name     string `yaml:"name" json:"name"`
	Interval int    `yaml:"interval,omitempty" json:"interval"`
	Rules    []Rule `yaml:"rules" json:"rules"`
}

type Rule struct {
	ID          int               `yaml:"-" json:"id"`
	Record      string            `yaml:"record,omitempty" json:"record,omitempty"`
	Alert       string            `yaml:"alert,omitempty" json:"name"`
	Expr        string            `yaml:"expr" json:"query"`
	For         int               `yaml:"for,omitempty" json:"duration"`
	GroupName   string            `yaml:"-" json:"group_name"`
	Labels      map[string]string `yaml:"labels,omitempty" json:"labels"`
	Annotations map[string]string `yaml:"annotations,omitempty" json:"annotations"`
}

type apiFuncResult struct {
	Data RuleGroups
}