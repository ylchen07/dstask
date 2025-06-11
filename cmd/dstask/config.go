package main

import (
	"os"
	"path"
)

// Config models the dstask application's required configuration. All paths
// are absolute.
type Config struct {
	Repo          string // Path to the git repository
	StateFile     string // Path to the dstask local state file. State will differ between machines
	IDsFile       string // Path to the ids file
	CtxFromEnvVar string // An unparsed context string, provided via DSTASK_CONTEXT
}

func (self *Config) InitConfig() {
	self.CtxFromEnvVar = getEnv("DSTASK_CONTEXT", "")
	self.Repo = getEnv("DSTASK_GIT_REPO", os.ExpandEnv("$HOME/.dstask"))
	self.StateFile = path.Join(self.Repo, ".git", "dstask", "state.bin")
	self.IDsFile = path.Join(self.Repo, ".git", "dstask", "ids.bin")
}

// NewConfig generates a new Config struct from the environment.
func NewConfig() Config {
	var conf Config

	conf.CtxFromEnvVar = getEnv("DSTASK_CONTEXT", "")
	conf.Repo = getEnv("DSTASK_GIT_REPO", os.ExpandEnv("$HOME/.dstask"))
	conf.StateFile = path.Join(conf.Repo, ".git", "dstask", "state.bin")
	conf.IDsFile = path.Join(conf.Repo, ".git", "dstask", "ids.bin")

	return conf
}

// getEnv returns an env var's value, or a default.
func getEnv(key string, _default string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}

	return _default
}
