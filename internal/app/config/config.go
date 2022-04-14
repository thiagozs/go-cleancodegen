package config

var Folders = []string{
	"build",
	"cmd",
	"internal/app/adapter/postgres/model",
	"internal/app/adapter/mysql/model",
	"internal/app/adapter/repository",
	"internal/app/adapter/service",
	"internal/app/application",
	"internal/app/application/usecase",
	"internal/app/application/service",
	"internal/app/domain/model",
	"internal/app/domain/repository",
	"internal/app/domain/valueobject",
	"pkg/utils",
	"tests",
}

var Files = []string{
	"LICENSE",
	"README.md",
	"Makefile",
	"main.go",
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
