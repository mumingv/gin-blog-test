package settings

import (
	"fmt"
	"testing"
)

func TestConfig(t *testing.T) {
	if err := Init(); err != nil {
		fmt.Printf("Loading config failed, err: %v", err)
	}
	fmt.Println(Conf.Mode)
	fmt.Println(Conf.MySQLConfig.Host)
}
