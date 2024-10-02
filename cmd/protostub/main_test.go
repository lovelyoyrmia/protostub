package main

import (
	"log"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCommand(t *testing.T) {

	testCases := []struct {
		name          string
		bCommand      []string
		filePath      string
		checkResponse func(t *testing.T, err error, filePath string)
	}{
		{
			name:     "SERVER_TYPE",
			bCommand: []string{"--service_dir=../../examples/service", "--type=server"},
			filePath: "../../examples/service/user_service_impl.go",
			checkResponse: func(t *testing.T, err error, filePath string) {
				require.NoError(t, err)
				require.FileExists(t, filePath)
			},
		},
		{
			name:     "CLIENT_TYPE",
			bCommand: []string{"--client_dir=../../examples/client", "--type=client"},
			filePath: "../../examples/client/user_service_client.go",
			checkResponse: func(t *testing.T, err error, filePath string) {
				require.NoError(t, err)
				require.FileExists(t, filePath)
			},
		},
		{
			name:     "REQUIRED_CLIENT_DIR",
			bCommand: []string{"--service_dir=../../examples/service", "--type=client"},
			filePath: "../../examples/client/user_service_client.go",
			checkResponse: func(t *testing.T, err error, filePath string) {
				require.Error(t, err)
			},
		},
		{
			name:     "REQUIRED_SERVER_DIR",
			bCommand: []string{"--client_dir=../../examples/client", "--type=server"},
			filePath: "../../examples/client/user_service_impl.go",
			checkResponse: func(t *testing.T, err error, filePath string) {
				require.Error(t, err)
			},
		},
	}

	for _, v := range testCases {
		t.Run(v.name, func(tt *testing.T) {
			cmdArgs := append([]string{"run", ".", "--proto_dir=../../examples", "--dest_dir=../../examples/pb"}, v.bCommand...)
			cmd := exec.Command("go", cmdArgs...)
			log.Println(cmd.String())
			err := cmd.Run()
			v.checkResponse(tt, err, v.filePath)
		})
	}
}
