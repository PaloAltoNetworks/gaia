package gaia

import (
	"errors"
	"reflect"
	"testing"

	"go.aporeto.io/elemental"
)

func checkValidationReturnsExpectedErr(obj elemental.Validatable, expected elemental.Error, t *testing.T) {
	t.Helper()

	err := obj.Validate()
	if err == nil && expected == (elemental.Error{}) {
		return
	}

	var errs elemental.Errors
	if ok := errors.As(err, &errs); !ok {
		t.Fatalf("unexpected error type\nwant: %T\n got: %T", errs, err)
	}

	// at this point, we should have at least one "legit" elemental error in the slice
	actualErr := errs[0]

	// we do this specific check so the error message is clear
	if expected == (elemental.Error{}) {
		t.Fatalf("expected no error, got: %#v", actualErr)
	}

	if !reflect.DeepEqual(actualErr, expected) {
		t.Fatalf("unexpected error\nwant: %#v\n got: %#v", expected, actualErr)
	}
}
