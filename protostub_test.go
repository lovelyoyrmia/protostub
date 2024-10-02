package protostub

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	testCases := []struct {
		name          string
		pbStub        *ProtoStub
		checkResponse func(t *testing.T, err error)
	}{
		{
			name: "SUCCESS",
			pbStub: &ProtoStub{
				ProtoDir:   "./examples",
				DestDir:    "./examples",
				ServiceDir: "./examples/service",
			},
			checkResponse: func(t *testing.T, err error) {
				require.NoError(t, err)
				require.DirExists(t, "./examples/service")
				require.DirExists(t, "./examples/pb")
				require.FileExists(t, "./examples/service/user_service_impl.go")
				require.FileExists(t, "./examples/service/customer_service_impl.go")
			},
		},
		{
			name:   "FAILED",
			pbStub: &ProtoStub{},
			checkResponse: func(t *testing.T, err error) {
				require.Error(t, err)
			},
		},
	}

	for _, v := range testCases {
		t.Run(v.name, func(tt *testing.T) {
			err := v.pbStub.Generate()
			v.checkResponse(tt, err)
		})
	}
}

func TestGenerateServices(t *testing.T) {
	testCases := []struct {
		name          string
		pbStub        *ProtoStub
		checkResponse func(t *testing.T, err error, services []*ServiceStub)
	}{
		{
			name: "SUCCESS",
			pbStub: &ProtoStub{
				ProtoDir:   "./examples",
				DestDir:    "./examples",
				ServiceDir: "./examples/services",
			},
			checkResponse: func(t *testing.T, err error, services []*ServiceStub) {
				require.NoError(t, err)
				res := &ServiceStub{
					ServiceName: "UserService",
					Package:     "pb",
					GoPackage:   "github.com/lovelyoyrmia/protostub/examples/pb",
					Method:      "GetUser",
					InputType:   "GetUserRequest",
					OutputType:  "GetUserResponse",
				}
				require.Equal(t, res, services[1])
			},
		},
		{
			name:   "NO_SERVICE",
			pbStub: &ProtoStub{},
			checkResponse: func(t *testing.T, err error, services []*ServiceStub) {
				require.Error(t, err)
				require.Empty(t, services)
			},
		},
	}

	for _, v := range testCases {
		t.Run(v.name, func(tt *testing.T) {
			services, err := v.pbStub.GenerateServices()
			v.checkResponse(tt, err, services)
		})
	}

}
