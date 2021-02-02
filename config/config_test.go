package config_test

import (
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"code.vegaprotocol.io/priceproxy/config"
)

func TestCheckConfig(t *testing.T) {
	err := config.CheckConfig(nil)
	assert.Equal(t, config.ErrNil, err)

	var cfg config.Config
	err = config.CheckConfig(&cfg)
	assert.True(t, strings.HasPrefix(err.Error(), config.ErrMissingConfigSection.Error()))

	cfg.Sources = []*config.SourceConfig{}
	err = config.CheckConfig(&cfg)
	assert.True(t, strings.HasPrefix(err.Error(), config.ErrEmptyConfigSection.Error()))

	cfg.Sources = append(cfg.Sources, &config.SourceConfig{})
	err = config.CheckConfig(&cfg)
	assert.True(t, strings.HasPrefix(err.Error(), config.ErrInvalidValue.Error()))

	cfg.Sources[0].SleepReal = 1
	err = config.CheckConfig(&cfg)
	assert.True(t, strings.HasPrefix(err.Error(), config.ErrInvalidValue.Error()))

	cfg.Sources[0].SleepWander = 1
	err = config.CheckConfig(&cfg)
	assert.True(t, strings.HasPrefix(err.Error(), config.ErrMissingConfigSection.Error()))

	cfg.Prices = []*config.PriceConfig{}
	err = config.CheckConfig(&cfg)
	assert.True(t, strings.HasPrefix(err.Error(), config.ErrEmptyConfigSection.Error()))

	cfg.Prices = append(cfg.Prices, &config.PriceConfig{})
	err = config.CheckConfig(&cfg)
	assert.True(t, strings.HasPrefix(err.Error(), config.ErrInvalidValue.Error()))

	cfg.Prices[0].Factor = 1
	err = config.CheckConfig(&cfg)
	assert.NoError(t, err)
}

func TestConfigureLogging(t *testing.T) {
	err := config.ConfigureLogging(nil)
	assert.Equal(t, config.ErrNil, err)

	var servercfg config.ServerConfig
	err = config.ConfigureLogging(&servercfg)
	assert.NoError(t, err)

	servercfg.LogLevel = "info"
	for _, lf := range []string{"json", "textcolour", "textnocolour", "fred"} {
		servercfg.LogFormat = lf
		err = config.ConfigureLogging(&servercfg)
		assert.NoError(t, err)
	}
}

func TestConfigStringFuncs(t *testing.T) {
	pc := config.PriceConfig{
		Base:   "BBB",
		Quote:  "QQQ",
		Source: "SRC",
		Wander: true,
	}
	assert.Equal(t, "{PriceConfig Base:BBB Quote:QQQ Source:SRC Wander:true}", pc.String())

	ps := config.SourceConfig{
		Name: "NNN",
		URL: url.URL{
			Scheme:   "https",
			Host:     "example.com",
			Path:     "/path",
			RawQuery: "a=b&x=y",
		},
		SleepReal:   11,
		SleepWander: 7,
	}
	assert.Equal(t, "{SourceConfig Name:NNN URL:https://example.com/path?a=b&x=y SleepReal:11s SleepWander:7s}", ps.String())
}
