package app

import (
	mock_app "github.com/benrowe/demo-web-services/src/app/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConfigFromEnv(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	// no env
	rm := mock_app.NewMockEnvironment(ctrl)

	rm.
		EXPECT().
		LookupEnv(gomock.Any()).
		Return("rubbish", false).
		AnyTimes()

	defaultCfg := LoadConfigFromEnv(rm)

	assert.Equal(t, "127.0.0.1", defaultCfg.DB.Host)
	assert.Equal(t, 80, defaultCfg.DB.Port)
	assert.Equal(t, "root", defaultCfg.DB.Username)
	assert.Equal(t, "password", defaultCfg.DB.Password)

	// ensure env variables are picked up
	m := mock_app.NewMockEnvironment(ctrl)

	m.EXPECT().LookupEnv(gomock.Eq("DB_HOST")).Return("localhost", true).Times(1)
	m.EXPECT().LookupEnv(gomock.Eq("DB_DIALECT")).Return("rubbish", false).Times(1)
	m.EXPECT().LookupEnv(gomock.Eq("DB_PORT")).Return("3333", true).Times(1)
	m.EXPECT().LookupEnv(gomock.Eq("DB_USERNAME")).Return("rubbish", false).Times(1)
	m.EXPECT().LookupEnv(gomock.Eq("DB_PASSWORD")).Return("rubbish", false).Times(1)
	m.EXPECT().LookupEnv(gomock.Eq("DB_DATABASE")).Return("rubbish", false).Times(1)
	m.EXPECT().LookupEnv(gomock.Eq("DB_CHARSET")).Return("rubbish", false).Times(1)

	cfg := LoadConfigFromEnv(m)

	assert.Equal(t, "mysql", cfg.DB.Dialect)
	assert.Equal(t, "localhost", cfg.DB.Host)
	assert.Equal(t, 3333, cfg.DB.Port)

}

func Test_env(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := mock_app.NewMockEnvironment(ctrl)

	// value exists
	m.EXPECT().LookupEnv(gomock.Eq("EXISTS")).Return("yes", true)

	assert.Equal(t, "yes", env(m, "EXISTS", "na"))

	// value does not exist
	m.EXPECT().LookupEnv(gomock.Eq("FOO")).Return("junk", false)

	assert.Equal(t, "default", env(m, "FOO", "default"))

}
