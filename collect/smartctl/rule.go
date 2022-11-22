package smartctl

import (
	"reflect"
	"strconv"
	"strings"
)

// 指标规则
type MetricRule struct {
	Key       string      `json:"key"`
	Type      string      `json:"type"`      //contains\gt\lt\gte\lte
	Threshold interface{} `json:"threshold"` //阈值
}

const (
	RuleTypeContains    = "contains"
	RuleTypeGt          = "gt"
	RuleTypeLt          = "lt"
	RuleTypeGte         = "gte"
	RuleTypeLte         = "lte"
	RuleTypeStringEqual = "string_equal"
	RuleTypeNumEqual    = "num_equal"
)

var ruleFunc = map[string]func(actual, expected interface{}) bool{
	RuleTypeContains:    ContainsFunc,
	RuleTypeGt:          GtFunc,
	RuleTypeLt:          LtFunc,
	RuleTypeGte:         GteFunc,
	RuleTypeLte:         LteFunc,
	RuleTypeStringEqual: StringEqualFunc,
	RuleTypeNumEqual:    NumEqualFunc,
}

func ContainsFunc(actual, expected interface{}) bool {
	return strings.Contains(actual.(string), expected.(string))
}

func GtFunc(actual, expected interface{}) bool {
	return convertFloat64(actual)-convertFloat64(expected) > 0
}

func GteFunc(actual, expected interface{}) bool {
	return convertFloat64(actual)-convertFloat64(expected) >= 0
}

func LtFunc(actual, expected interface{}) bool {
	return convertFloat64(actual)-convertFloat64(expected) < 0
}

func LteFunc(actual, expected interface{}) bool {
	return convertFloat64(actual)-convertFloat64(expected) <= 0
}

func StringEqualFunc(actual, expected interface{}) bool {
	return actual.(string) == expected.(string)
}
func NumEqualFunc(actual, expected interface{}) bool {
	return convertFloat64(actual) == convertFloat64(expected)
}

func convertFloat64(v interface{}) float64 {
	switch reflect.TypeOf(v) {
	case reflect.TypeOf(""):
		v, _ = strconv.ParseFloat(v.(string), 10)
	case reflect.TypeOf(0):
		return float64(v.(int))
	case reflect.TypeOf(0.0):
		return v.(float64)
	}
	return 0.0
}

func (rule *MetricRule) IsPass(actual interface{}) bool {
	f := ruleFunc[rule.Type]
	if f == nil {
		return false
	}
	return f(actual, rule.Threshold)
}
