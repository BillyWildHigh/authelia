package schema_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/authelia/authelia/v4/internal/configuration/schema"
)

type TestNestedStruct struct {
	MustBe5 int
}

func (tns *TestNestedStruct) Validate(validator *schema.StructValidator) {
	if tns.MustBe5 != 5 {
		validator.Push(fmt.Errorf("MustBe5 must be 5"))
	}
}

type TestStruct struct {
	MustBe10       int
	ShouldBeAbove5 int
	NotEmpty       string
	SetDefault     string
	Nested         TestNestedStruct
	Nested2        TestNestedStruct
	NilPtr         *int
	NestedPtr      *TestNestedStruct
}

func (ts *TestStruct) Validate(validator *schema.StructValidator) {
	if ts.MustBe10 != 10 {
		validator.Push(fmt.Errorf("MustBe10 must be 10"))
	}

	if ts.NotEmpty == "" {
		validator.Push(fmt.Errorf("NotEmpty must not be empty"))
	}

	if ts.ShouldBeAbove5 <= 5 {
		validator.PushWarning(fmt.Errorf("ShouldBeAbove5 should be above 5"))
	}

	if ts.SetDefault == "" {
		ts.SetDefault = "xyz"
	}
}

func TestValidator(t *testing.T) {
	validator := schema.NewValidator()

	s := TestStruct{
		MustBe10:  5,
		NotEmpty:  "",
		NestedPtr: &TestNestedStruct{},
	}

	err := validator.Validate(&s)
	if err != nil {
		panic(err)
	}

	errs := validator.Errors()
	assert.Equal(t, 4, len(errs))

	assert.Equal(t, 2, len(errs["root"]))
	assert.ElementsMatch(t, []error{
		fmt.Errorf("MustBe10 must be 10"),
		fmt.Errorf("NotEmpty must not be empty")}, errs["root"])

	assert.Equal(t, 1, len(errs["root.Nested"]))
	assert.ElementsMatch(t, []error{
		fmt.Errorf("MustBe5 must be 5")}, errs["root.Nested"])

	assert.Equal(t, 1, len(errs["root.Nested2"]))
	assert.ElementsMatch(t, []error{
		fmt.Errorf("MustBe5 must be 5")}, errs["root.Nested2"])

	assert.Equal(t, 1, len(errs["root.NestedPtr"]))
	assert.ElementsMatch(t, []error{
		fmt.Errorf("MustBe5 must be 5")}, errs["root.NestedPtr"])

	assert.Equal(t, "xyz", s.SetDefault)
}

func TestStructValidator(t *testing.T) {
	validator := schema.NewStructValidator()
	s := TestStruct{
		MustBe10:       5,
		ShouldBeAbove5: 2,
		NotEmpty:       "",
		NestedPtr:      &TestNestedStruct{},
	}
	s.Validate(validator)

	assert.True(t, validator.HasWarnings())
	assert.True(t, validator.HasErrors())

	require.Len(t, validator.Warnings(), 1)
	require.Len(t, validator.Errors(), 2)

	assert.EqualError(t, validator.Warnings()[0], "ShouldBeAbove5 should be above 5")
	assert.EqualError(t, validator.Errors()[0], "MustBe10 must be 10")
	assert.EqualError(t, validator.Errors()[1], "NotEmpty must not be empty")

	validator.Clear()

	assert.False(t, validator.HasWarnings())
	assert.False(t, validator.HasErrors())

	assert.Len(t, validator.Warnings(), 0)
	assert.Len(t, validator.Errors(), 0)
}
