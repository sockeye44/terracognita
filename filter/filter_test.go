package filter_test

import (
	"testing"

	"github.com/sockeye44/terracognita/errcode"
	"github.com/sockeye44/terracognita/filter"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsExcluded(t *testing.T) {
	t.Run("True", func(t *testing.T) {
		f := filter.Filter{Exclude: []string{"a", "b"}}
		assert.True(t, f.IsExcluded("a"))
	})
	t.Run("True", func(t *testing.T) {
		f := filter.Filter{Exclude: []string{"a", "b"}}
		assert.True(t, f.IsExcluded("a", "b"))
	})
	t.Run("False", func(t *testing.T) {
		f := filter.Filter{Exclude: []string{"a", "b"}}
		assert.False(t, f.IsExcluded("c"))
	})
	t.Run("False", func(t *testing.T) {
		f := filter.Filter{Exclude: []string{"a", "b"}}
		assert.False(t, f.IsExcluded("a", "c"))
	})
}

func TestIsIncluded(t *testing.T) {
	t.Run("True", func(t *testing.T) {
		f := filter.Filter{Include: []string{"a", "b"}}
		assert.True(t, f.IsIncluded("a"))
	})
	t.Run("True", func(t *testing.T) {
		f := filter.Filter{Include: []string{"a", "b"}}
		assert.True(t, f.IsIncluded("a", "b"))
	})
	t.Run("False", func(t *testing.T) {
		f := filter.Filter{Include: []string{"a", "b"}}
		assert.False(t, f.IsIncluded("c"))
	})
	t.Run("False", func(t *testing.T) {
		f := filter.Filter{Include: []string{"a", "b"}}
		assert.False(t, f.IsIncluded("a", "c"))
	})
}

func TestTargetsTypesWithIDs(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		f := filter.Filter{Targets: []string{"aws_instance.2", "aws_instance.3", "aws_iam_user.2", "aws_instance.2"}}
		assert.Equal(t, map[string][]string{
			"aws_instance": []string{"2", "3"},
			"aws_iam_user": []string{"2"},
		}, f.TargetsTypesWithIDs())
	})
}

func TestValidate(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		f := filter.Filter{Targets: []string{"aws_instance.2"}}
		err := f.Validate()
		require.NoError(t, err)
	})
	t.Run("SuccessWithMultipleDots", func(t *testing.T) {
		f := filter.Filter{Targets: []string{"aws_instance.pepito.grillo.extra"}}
		err := f.Validate()
		require.NoError(t, err)
	})
	t.Run("ErrorNoID", func(t *testing.T) {
		f := filter.Filter{Targets: []string{"aws_instance"}}
		err := f.Validate()
		assert.Error(t, errors.Cause(err), errcode.ErrFilterTargetsInvalid)
	})
}
