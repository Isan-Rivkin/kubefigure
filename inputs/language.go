package inputs

import (
	"github.com/isan-rivkin/kubefigure/sources"
)

type ValueToken struct {
	Type               sources.SourceType
	Value              string
	ValueIndicatorPath string
}

func NewValueToken(val, valIndicator string, t sources.SourceType) *ValueToken {
	return &ValueToken{
		Type:               t,
		Value:              val,
		ValueIndicatorPath: valIndicator,
	}
}
