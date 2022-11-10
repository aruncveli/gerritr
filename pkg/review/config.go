package review

import (
	"fmt"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
)

/*
A [koanf] config with dot as delimiter.

Loads $XDG_CONFIG_HOME/gerritr/config.yml expecting a structure like below:

	aliases:
		team1:
	  	- b1@org.com
	  	- b2@org.com
		team2:
	  	- f1@org.com
	  	- f2@org.com

[koanf]: https://pkg.go.dev/github.com/knadh/koanf
*/
var Config = koanf.New(".")

func init() {
	fPath := filepath.Join(xdg.ConfigHome, "gerritr", "config.yml")
	fProvider := file.Provider(fPath)
	err := Config.Load(fProvider, yaml.Parser())
	if err != nil {
		fmt.Println("Reading config failed")
		fmt.Println(err)
	}
}
