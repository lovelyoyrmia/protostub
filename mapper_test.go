package protostub

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConvertSnakeCase(t *testing.T) {

	tests := []struct {
		req      string
		expected string
	}{
		{
			req:      "UserService",
			expected: "user_service_impl.go",
		},
		{
			req:      "userService",
			expected: "user_service_impl.go",
		},
	}

	for _, v := range tests {
		t.Run("TO_SNAKE_CASE", func(tt *testing.T) {
			expected := toSnakeCase(v.req)
			require.Equal(tt, expected, v.expected)
		})
	}
}
