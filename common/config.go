package common

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

var (
	// Conf - The global configure.
	Conf config
)

type config struct {
	Service service
}

type service struct {
	ListenAddr         string
	TCPServerAddr      string
	TCPServerAddrDebug string
}

// InitConfig - Parse configure file by path.
func InitConfig(filepath string) (err error) {
	var (
		content []byte
	)
	if content, err = ioutil.ReadFile(filepath); err != nil {
		return err
	}
	if _, err = toml.Decode(string(content), &Conf); err != nil {
		return err
	}
	return nil
}
