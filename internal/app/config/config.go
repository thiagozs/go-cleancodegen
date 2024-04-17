package config

var Folders = []string{
	"api",
	"build",
	"config",
	"cmd",
	"docs",
	"infra/adapter/postgres",
	"infra/adapter/mysql",
	"infra/adapter/cache",
	"internal/domain",
	"internal/services",
	"internal/initiate",
	"pkg/utils",
	"tmp",
	".github",
}

var Files = map[string]string{
	"LICENSE.md":         "license.tmpl",
	"README.md":          "readme.tmpl",
	"Makefile":           "makefile.tmpl",
	"main.go":            "main.tmpl",
	".gitignore":         "gitignore.tmpl",
	"dockerfile":         "dockerfile.tmpl",
	"docker-compose.yml": "docker-compose.tmpl",
	".envrc":             "envrc.tmpl",
	"config/config.go":   "config.tmpl",
	"tmp/.gitkeep":       "gitkeep.tmpl",
	"build/.gitkeep":     "gitkeep.tmpl",
}

var Cfg *CfgParams

type ConfigOpts func(c *CfgParams) error

type CfgParams struct {
	Path    string
	Project string
}

func NewCfgParams(opts ...ConfigOpts) *CfgParams {
	c := &CfgParams{}
	for _, opt := range opts {
		if err := opt(c); err != nil {
			panic(err)
		}
	}
	Cfg = c
	return c
}

func CfgProject(name string) ConfigOpts {
	return func(c *CfgParams) error {
		c.Project = name
		return nil
	}
}

func CfgPath(path string) ConfigOpts {
	return func(c *CfgParams) error {
		c.Path = path
		return nil
	}
}
