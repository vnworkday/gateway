package utils

import "go.uber.org/fx"

func FxNamedRegister(constructor interface{}, name string, params ...string) any {
	paramTags := make([]string, 0, len(params))

	for _, p := range params {
		paramTags = append(paramTags, `name:"`+p+`"`)
	}

	return fx.Annotate(
		constructor,
		fx.ParamTags(paramTags...),
		fx.ResultTags(`name:"`+name+`"`),
	)
}

func FxGroupedRegister(constructor interface{}, group string, g interface{}) any {
	return fx.Annotate(
		constructor,
		fx.As(g),
		fx.ResultTags(`group:"`+group+`"`),
	)
}
