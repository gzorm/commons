package service

import (
	"testing"

	"github.com/gzorm/common/core/logx"
	"github.com/gzorm/common/internal/devserver"
	"github.com/stretchr/testify/assert"
)

func TestServiceConf(t *testing.T) {
	c := ServiceConf{
		Name: "foo",
		Log: logx.LogConf{
			Mode: "console",
		},
		Mode: "dev",
		DevServer: devserver.Config{
			Port:       6470,
			HealthPath: "/healthz",
		},
	}
	c.MustSetUp()
}

func TestServiceConfWithMetricsUrl(t *testing.T) {
	c := ServiceConf{
		Name: "foo",
		Log: logx.LogConf{
			Mode: "volume",
		},
		Mode:       "dev",
		MetricsUrl: "http://localhost:8080",
	}
	assert.NoError(t, c.SetUp())
}
