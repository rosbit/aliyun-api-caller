package conf

import (
	"gopkg.in/yaml.v2"
	"fmt"
	"os"
	"time"
	"path"
)

var (
	ServiceConf struct {
		AccessKey struct {
			AccessKeyId  string `yaml:"access-key-id"`
			AccessSecret string `yaml:"access-secret"`
		} `yaml:"access-key"`
	}

	Loc = time.FixedZone("UTC+8", 8*60*60)
)

func getEnv(name string, result *string, must bool) error {
	s := os.Getenv(name)
	if s == "" {
		if must {
			return fmt.Errorf("env \"%s\" not set", name)
		}
	}
	*result = s
	return nil
}

func CheckGlobalConf() error {
	exePath, err := getExecWD()
	if err != nil {
		return err
	}

	var p string
	getEnv("TZ", &p, false)
	if p != "" {
		if loc, err := time.LoadLocation(p); err == nil {
			Loc = loc
		}
	}

	var confFile string
	if err := getEnv("CONF_FILE", &confFile, true); err != nil {
		return err
	}
	confFile = toAbsPath(exePath, confFile)

	fp, err := os.Open(confFile)
	if err != nil {
		return err
	}
	defer fp.Close()
	dec := yaml.NewDecoder(fp)
	if err = dec.Decode(&ServiceConf); err != nil {
		return err
	}

	if err = checkMust(confFile); err != nil {
		return err
	}

	return nil
}

func checkMust(confFile string) error {
	ak := &ServiceConf.AccessKey
	if len(ak.AccessKeyId) == 0 {
		return fmt.Errorf("access-key/access-key-id expected in conf")
	}
	if len(ak.AccessSecret) == 0 {
		return fmt.Errorf("access-key/access-secret expected in conf")
	}

	return nil
}

func getExecWD() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	return path.Dir(exePath), nil
}

func toAbsPath(absRoot, filePath string) string {
	if path.IsAbs(filePath) {
		return filePath
	}
	return path.Join(absRoot, filePath)
}
