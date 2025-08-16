package nilo

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResult(t *testing.T) {
	t.Run("FromResult", func(t *testing.T) {
		function := func(b bool) (*int, error) {
			if b {
				integer := 10
				return &integer, nil
			}
			return nil, errors.New("some error")
		}

		function2 := func(p *int) (*int, error) {
			if *p == 0 {
				return nil, errors.New("some error 2")
			}
			r := *p + 1
			return &r, nil
		}

		t.Run("when value is not nil", func(t *testing.T) {
			opt := FromResult(function(true))
			assert.Equal(t, 10, *opt.Unwrap())
		})

		t.Run("when value is nil", func(t *testing.T) {
			opt := FromResult(function(false))
			assert.True(t, opt.IsNone())
		})

		t.Run("AndThen when value is not nil", func(t *testing.T) {
			opt := FromResult(function(true)).OkAndResult(function2)
			assert.Equal(t, 11, *opt.Unwrap())
		})

		t.Run("AndThen when value is nil", func(t *testing.T) {
			opt := FromResult(function(false)).OkAndResult(function2)
			assert.True(t, opt.IsNone())
		})
	})
}
