package review

import (
	"fmt"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

var Config = koanf.New(".")

func InitConfig() {
	fPath := filepath.Join(xdg.ConfigHome, "gerritr", "config.yml")
	fProvider := file.Provider(fPath)
	err := Config.Load(fProvider, yaml.Parser())
	if err != nil {
		fmt.Println("Reading config failed", err)
		fmt.Println("Reverting to defaults wherever possible")
	}

	SetAllowedEmailDomains()
}
