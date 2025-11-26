package config

import (
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func init() {
	style := log.DefaultStyles()
	style.Timestamp = style.Timestamp.Foreground(lipgloss.ANSIColor(8))

	for k, val := range style.Levels {
		style.Levels[k] = val.PaddingLeft(2).PaddingRight(2).Width(10).MaxWidth(10).AlignHorizontal(lipgloss.Center)
	}
	style.Levels[log.ErrorLevel] = style.Levels[log.ErrorLevel].Foreground(lipgloss.ANSIColor(14)).Background(lipgloss.ANSIColor(1))
	style.Levels[log.WarnLevel] = style.Levels[log.WarnLevel].Foreground(lipgloss.ANSIColor(0)).Background(lipgloss.ANSIColor(3))
	style.Levels[log.InfoLevel] = style.Levels[log.InfoLevel].Foreground(lipgloss.ANSIColor(14)).Background(lipgloss.ANSIColor(4))
	style.Levels[log.DebugLevel] = style.Levels[log.DebugLevel].Foreground(lipgloss.ANSIColor(14)).Background(lipgloss.ANSIColor(5))
	log.SetStyles(style)

	log.SetLevel(log.DebugLevel)
	log.SetTimeFormat(time.TimeOnly)
}

func New() *viper.Viper {
	conf := viper.NewWithOptions(
		viper.EnvKeyReplacer(
			strings.NewReplacer("-", "_", ".", "_"),
		),
	)

	conf.AddConfigPath(".")
	conf.AddConfigPath("/etc")

	conf.SetConfigName("aura-print")

	if err := godotenv.Load(); err != nil {
		log.Warn("⚠️  No .env file found, using system environment variables")
	}

	conf.SetEnvPrefix("AP")
	conf.AutomaticEnv()

	if err := conf.ReadInConfig(); err != nil {
		log.Warn("Something wrong with config file", "error", err)
	}

	return conf
}
