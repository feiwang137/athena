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
	Interval int    `yaml:"interval" json:"interval"`
	Rules    []Rule `yaml:"rules" json:"rules"`
}

type Rule struct {
	Type        string            `yaml:"type" json:"type"`
	Alert       string            `yaml:"name,omitempty" json:"name"`
	Expr        string            `yaml:"expr" json:"query"`
	For         int               `yaml:"for,omitempty" json:"duration,omitempty"`
	GroupName   string            `yaml:"name" json:"group_name"`
	Labels      map[string]string `yaml:"labels,omitempty" json:"labels"`
	Annotations map[string]string `yaml:"annotations,omitempty" json:"annotations"`
}

type apiFuncResult struct {
	Data RuleGroups
}



