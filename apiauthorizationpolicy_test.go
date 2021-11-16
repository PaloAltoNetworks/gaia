package gaia

import (
	"testing"

	"go.aporeto.io/elemental"
)

func TestAPIAuthorizationPolicy_Validate_custom(t *testing.T) {

	cases := map[string]struct {
		p   *APIAuthorizationPolicy
		err elemental.Error
	}{
		"valid": {
			p: &APIAuthorizationPolicy{
				Name:                 t.Name(),
				AuthorizedIdentities: []string{"service,post"},
				AuthorizedNamespace:  "/test1",
				AuthorizedNamespaces: []string{"/test2"},
			},
		},
		"valid/only-AuthorizedNamespace": {
			p: &APIAuthorizationPolicy{
				Name:                 t.Name(),
				AuthorizedIdentities: []string{"service,post"},
				AuthorizedNamespace:  "/test",
			},
		},
		"valid/only-AuthorizedNamespaces": {
			p: &APIAuthorizationPolicy{
				Name:                 t.Name(),
				AuthorizedIdentities: []string{"service,post"},
				AuthorizedNamespaces: []string{"/test"},
			},
		},
		"invalid": {
			p: &APIAuthorizationPolicy{
				Name:                 t.Name(),
				AuthorizedIdentities: []string{"service,post"},
				AuthorizedNamespace:  "",
				AuthorizedNamespaces: nil,
			},
			err: makeValidationError("authorizedNamespaces", "A minimum of one authorized namespace must be given"),
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			checkValidationReturnsExpectedErr(c.p, c.err, t)
		})
	}
}
