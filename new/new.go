package new

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/lovego/fs"
	"github.com/spf13/cobra"
)

func Cmd() *cobra.Command {
	var typ string
	var force bool
	cmd := &cobra.Command{
		Use: `new <module path> <registry> [<domain>] [flags]

module path: Module path for the project, required. The last element of the module path is used as project name and directory.
   registry: Docker registry url for images of the project, required.
     domain: Parent domain for the project, required for app/web project. Used for config.yml, access.conf.tmpl, readme.md, .gitlab-ci.yml.`,
		Short:   `Create a new project.`,
		Example: `  xiaomei new accounts registry.abc.com/go abc.com`,
		// DisableFlagsInUseLine: true,
		RunE: func(c *cobra.Command, args []string) error {
			var expect = 3
			if typ == `logc` {
				expect = 2
			}
			if len(args) != expect {
				return fmt.Errorf(`exactly %d arguments required for %s project.`, expect, typ)
			}
			var domain string
			if len(args) == 3 {
				domain = args[2]
			}
			return New(typ, args[0], args[1], domain, force)
		},
	}
	cmd.Flags().StringVarP(&typ, `type`, `t`, `app`, `project type.
 app: only service that provides Golang API.
 web: only service that provides frontend UI.
logc: only service that collect logs to ElasticSearch.
`)
	cmd.Flags().BoolVarP(&force, `force`, `f`, false, `force overwrite existing files.`)
	return cmd
}

func New(typ, dir, registry, domain string, force bool) error {
	if dir == `` {
		return errors.New(`dir can't be empty.`)
	}
	var err error
	if !filepath.IsAbs(dir) {
		if dir, err = filepath.Abs(dir); err != nil {
			return err
		}
	}
	if registry != "" && registry[len(registry)-1] != '/' {
		registry += "/"
	}

	config, err := getConfig(typ, dir, registry, domain)
	if err != nil {
		return err
	}
	tmplsDir := filepath.Join(fs.GetGoSrcPath(), `github.com/lovego/xiaomei/new/templates/`+typ)
	return walk(tmplsDir, dir, config, force)
}
