package deploy

import (
	"github.com/lovego/xiaomei/release"
	//	"github.com/lovego/xiaomei/registry"
)

func GetCommonArgs(svcName, env, tag string) []string {
	args := []string{`-e`, release.EnvironmentEnvVar + `=` + env}

	service := release.GetService(svcName, env)
	args = append(args, service.Options...)
	args = append(args, service.ImageName(tag))
	args = append(args, service.Command...)
	return args
}
