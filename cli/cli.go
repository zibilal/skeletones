package cli

import (
	"github.com/spf13/cobra"
	"github.com/zibilal/skeletones/logger"
	"io"
	"sync"
)

type CommandLineInterface struct {
	cmds    map[string]*cobra.Command
	rootCmd *cobra.Command
	locker  sync.Mutex
}

func (c *CommandLineInterface) GetRootCommand() *cobra.Command {
	return c.rootCmd
}

func (c *CommandLineInterface) AddCli(name, iuse, ishort, ilong string, handler CommandLineFunc) {
	cobraCmd := &cobra.Command{
		Use:   iuse,
		Short: ishort,
		Long:  ilong,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := handler(args...)
			if err != nil {
				logger.Fatal("[CLI][Error]Unable to run " + name + " command, due to: " + err.Error())
			}
		},
	}

	c.locker.Lock()
	c.cmds[name] = cobraCmd
	c.locker.Unlock()
}

func (c *CommandLineInterface) SetArgs(args []string) {
	c.rootCmd.SetArgs(args)
}

func (c *CommandLineInterface) ExecuteCli(w io.Writer) error {

	for _, cmd := range c.cmds {
		c.rootCmd.AddCommand(cmd)
	}
	c.rootCmd.SetOutput(w)
	return c.rootCmd.Execute()
}

func NewCommandLineInterface(appName string) *CommandLineInterface {
	theCli := new(CommandLineInterface)
	theCli.cmds = make(map[string]*cobra.Command)
	theCli.rootCmd = &cobra.Command{Use: appName}
	return theCli
}

type CommandLineFunc func(args ...string) error
