package tunnel

import (
	"github.com/elliotchance/sshtunnel"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSSHTunnel(t *testing.T) {
	tunnel := sshtunnel.NewSSHTunnel(
		"pengganyu@10.20.34.27",
		sshtunnel.PrivateKeyFile("/Users/pgy/.ssh/id_rsa"),
		"10.20.87.32:3306",
		"3306",
	)
	ast := assert.New(t)
	err := tunnel.Start()
	ast.Nil(err)
}
