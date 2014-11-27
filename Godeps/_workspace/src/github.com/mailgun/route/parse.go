package route

import (
	"fmt"

	"github.com/mailgun/vulcand/Godeps/_workspace/src/github.com/mailgun/predicate"
)

func parse(expression string, result *match) (matcher, error) {
	p, err := predicate.NewParser(predicate.Def{
		Functions: map[string]interface{}{
			"Host":       hostTrieMatcher,
			"HostRegexp": hostRegexpMatcher,

			"Path":       pathTrieMatcher,
			"PathRegexp": pathRegexpMatcher,

			"Method": methodTrieMatcher,

			"Header":       headerTrieMatcher,
			"HeaderRegexp": headerRegexpMatcher,
		},
		Operators: predicate.Operators{
			AND: newAndMatcher,
		},
	})
	if err != nil {
		return nil, err
	}
	out, err := p.Parse(expression)
	if err != nil {
		return nil, err
	}
	m, ok := out.(matcher)
	if !ok {
		return nil, fmt.Errorf("unknown result type: %T", out)
	}
	m.setMatch(result)
	return m, nil
}
