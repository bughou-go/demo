package new

import (
	"errors"
	"path/filepath"

	"github.com/lovego/xiaomei/utils/fs"
)

func getTmplDir(isInfra bool) (string, error) {
	srcPath, err := fs.GetGoSrcPath()
	if err != nil {
		return ``, err
	}
	tmplDir := filepath.Join(srcPath, `github.com/lovego/xiaomei/xiaomei/new`)
	if isInfra {
		return filepath.Join(tmplDir, `infra`), nil
	} else {
		return filepath.Join(tmplDir, `webapp`), nil
	}
}

func getProjectPath(proDir string) (string, error) {
	if proDir == `` {
		return ``, errors.New(`project path can't be empty.`)
	}

	if !filepath.IsAbs(proDir) {
		var err error
		if proDir, err = filepath.Abs(proDir); err != nil {
			return ``, err
		}
	}

	srcPath, err := fs.GetGoSrcPath()
	if err != nil {
		return ``, err
	}

	proPath, err := filepath.Rel(srcPath, proDir)
	if err != nil {
		return ``, err
	}
	if proPath[0] == '.' {
		return ``, errors.New(`project dir must be under ` + srcPath + "\n")
	}
	return proPath, nil
}
