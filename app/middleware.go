package app

import (
	"strings"

	"github.com/antonmedv/expr"
	"github.com/antonmedv/expr/vm"
)

type Middleware struct {
	Rule    string
	program *vm.Program
	geo     *Geo
}

func exprEnv(geo *Geo, address string) map[string]interface{} {
	return map[string]interface{}{
		"ip":      address,
		"city":    geo.City,
		"country": geo.Country,
	}
}

func NewMiddleware(rule string, geo *Geo) (*Middleware, error) {
	p, err := expr.Compile(rule, expr.Env(exprEnv(geo, "")))
	if err != nil {
		return nil, err
	}

	return &Middleware{
		Rule:    rule,
		program: p,
		geo:     geo,
	}, nil
}

func (m *Middleware) Passes(address string) bool {
	ip, _, f := strings.Cut(address, ":")
	if !f {
		return false
	}
	v, err := expr.Run(m.program, exprEnv(m.geo, ip))
	if err != nil {
		return false
	}

	switch v.(type) {
	case bool:
		return v.(bool)
	case int:
		return v.(int) > 0
	case float64:
		return v.(float64) > 0
	case float32:
		return v.(float32) > 0
	case string:
		return v.(string) != ""
	default:
		return false
	}
}
