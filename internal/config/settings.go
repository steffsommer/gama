package config

import "time"

const githubAPIURL = "https://api.github.com"

func fillDefaultSettings(cfg *Config) *Config {
	if cfg.Settings.LiveMode.Interval == time.Duration(0) {
		cfg.Settings.LiveMode.Interval = 15 * time.Second
	}
	if cfg.Github.URL == "" {
		cfg.Github.URL = githubAPIURL
	}
	return cfg
}
