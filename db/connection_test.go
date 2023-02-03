package db

import (
	"testing"

	"github.com/juliotorresmoreno/SpecialistTalk/configs"
	"github.com/stretchr/testify/require"
)

func TestNewEngigne(t *testing.T) {
	conf := configs.GetConfig()
	require := require.New(t)
	_, err := NewEngigne(conf.Database)
	require.NoError(err)
}
