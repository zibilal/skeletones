package cli

import (
	"bytes"
	"fmt"
	"testing"
)

const (
	success = "\u2713"
	failed  = "\u2717"
)

func TestNewCommandLineInterface(t *testing.T) {
	t.Log("Test CommandLineInterface")
	{
		buff := bytes.NewBuffer([]byte{})
		appCli := NewCommandLineInterface("skeletones")
		appCli.AddCli("serve", "serve [service-name]", "Start service",
			"Start service define in the service name", CommandLineFunc(func(args ...string) error {

				fmt.Println("The service!!! ", args)

				return nil

			}))
		appCli.AddCli("consumer", "consumer [consumer-name]", "Start consumer",
			"Start consumer defined in the consumer name", CommandLineFunc(func(args ...string) error {

				fmt.Println("The consumer!!! ", args)

				return nil
			}))
		appCli.SetArgs([]string{"consumer", "default consumer"})
		appCli.ExecuteCli(buff)

		if buff.String() == "" {
			t.Logf("%s expected buff is empty", success)
		} else {
			t.Errorf("%s expected buff is empty, got %s", failed, buff.String())
		}

		appCli.SetArgs([]string{"serve", "default service"})
		appCli.ExecuteCli(buff)

		if buff.String() == "" {
			t.Logf("%s expected buff is empty", success)
		} else {
			t.Errorf("%s expected buff is empty, got %s", failed, buff.String())
		}
	}
}
