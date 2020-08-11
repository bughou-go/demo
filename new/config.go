package new

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"path/filepath"
	"strings"

	"github.com/lovego/fs"
)

type Config struct {
	ProName        string
	ProPath        string
	ProNameUrlSafe string
	Domain         string
	Registry       string
	RepoPrefix     string
}

func getConfig(typ, dir, registry, domain string) (*Config, error) {
	var proName = filepath.Base(dir)
	var proPath string
	if typ == `app` {
		var err error
		if proPath, err = getProjectPath(dir); err != nil {
			return nil, err
		}
	}

	config := &Config{
		ProName:        proName,
		ProNameUrlSafe: strings.Replace(proName, `_`, `-`, -1),
		ProPath:        proPath,
		Domain:         domain,
		Registry:       registry,
	}
	if i := strings.IndexByte(registry, '/'); i > 0 {
		config.RepoPrefix = strings.TrimSuffix(registry[i:], "/")
	}

	return config, nil
}

func getProjectPath(dir string) (string, error) {
	srcPath := fs.GetGoSrcPath()
	proPath, err := filepath.Rel(srcPath, dir)
	if err != nil {
		return ``, err
	}
	if proPath[0] == '.' {
		return ``, errors.New(`project dir must be under ` + srcPath + "\n")
	}
	return proPath, nil
}

// 32 byte hex string
func genSecret() string {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}
