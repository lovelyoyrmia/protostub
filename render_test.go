package protostub

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRenderTemplate(t *testing.T) {

	testCases := []struct {
		name          string
		stub          *ServiceStub
		checkResponse func(t *testing.T, err error, data []byte)
	}{
		{
			name: "SUCCESS",
			stub: &ServiceStub{
				ServiceName: "UserService",
				Package:     "pb",
				GoPackage:   "github.com/lovelyoyrmia/protostub/examples/pb",
				Method:      "GetUser",
				InputType:   "GetUserRequest",
				OutputType:  "GetUserResponse",
			},
			checkResponse: func(t *testing.T, err error, data []byte) {
				require.NoError(t, err)
				require.NotEmpty(t, data)
			},
		},
		{
			name: "FAILED",
			stub: nil,
			checkResponse: func(t *testing.T, err error, data []byte) {
				require.Error(t, err)
				require.Empty(t, data)
			},
		},
	}

	for _, v := range testCases {
		t.Run(v.name, func(t *testing.T) {
			data, err := RenderTemplate(v.stub)
			v.checkResponse(t, err, data)
		})
	}
}
