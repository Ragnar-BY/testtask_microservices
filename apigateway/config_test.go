package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_settings_LoadOptionsFromFile(t *testing.T) {
	tt := []struct {
		name             string
		filename         string
		expectError      bool
		expectedSettings settings
	}{
		{
			name:     "Success",
			filename: "testdata/settings.yaml",
			expectedSettings: settings{
				ServerAddress: ":8080",
				PlayerAddress: ":8081",
				ConfigFile:    "testdata/settings.yaml",
			},
			expectError: false,
		},
		{
			name:        "FileError",
			filename:    "testdata/filenotexist.yaml",
			expectError: true,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s := settings{ConfigFile: tc.filename}
			err := s.LoadOptionsFromFile()
			if tc.expectError {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tc.expectedSettings, s)
			}
		})
	}
}

func TestSettings_Parse(t *testing.T) {
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	tt := []struct {
		name                  string
		args                  []string
		expectedPlayers       string
		expectedServerAddress string
		expectedConfig        string
	}{
		{
			name:                  "success",
			args:                  []string{"cmd", "--dbname=db", "-a=127.0.0.1", "--players=:8082", "--server=:1234"},
			expectedPlayers:       ":8082",
			expectedServerAddress: ":1234",
			expectedConfig:        "",
		},
		{
			name:                  "SuccessPartialSettings",
			args:                  []string{"cmd", "--dbname=db", "-a=127.0.0.1"},
			expectedPlayers:       ":8081",
			expectedServerAddress: ":8080",
			expectedConfig:        "",
		},
		{
			name:                  "SuccessLostConfigFile",
			args:                  []string{"cmd", "--dbname=db", "-a=127.0.0.1", "--players=:8082", "--server=:1234", "--configfile=testdata/filenotexist"},
			expectedPlayers:       ":8082",
			expectedServerAddress: ":1234",
			expectedConfig:        "testdata/filenotexist",
		},
		{
			name:                  "SuccessConfigFile",
			args:                  []string{"cmd", "--dbname=testDB", "-a=127.0.0.1", "--players=coll", "--server=:1234", "--configfile=testdata/config.yaml"},
			expectedPlayers:       ":8081",
			expectedServerAddress: ":1234",
			expectedConfig:        "testdata/config.yaml",
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s := settings{}
			os.Args = tc.args
			err := s.Parse()
			require.NoError(t, err)
			assert.Equal(t, tc.expectedServerAddress, s.ServerAddress)
			assert.Equal(t, tc.expectedConfig, s.ConfigFile)
		})
	}
}
