package cli

import (
	"context"
	"fmt"

	"github.com/timickb/password-manager/internal/common"
	"github.com/timickb/password-manager/internal/installer"
	"github.com/timickb/password-manager/pkg/api"
)

type CLI interface {
	Execute(cmd string, args []string)
}

const (
	commandNotFound string = "Command not found. Type /help to get a list of available commands"
	getCmdUsage     string = "Usage: get <secret_name>"
	setCmdUsage     string = "Usage: set <secret_name> <secret_value>"
	deleteCmdUsage  string = "Usage: delete <secret_name>"
)

type Command struct {
	Name      string
	ArgsCount int
	Usage     string
}

type PasswordManagerCLI struct {
	commands map[string]*Command
	pm       *api.PasswordManager
	ctx      *context.Context
	ins      installer.Installer
}

func New(ctx *context.Context, pm *api.PasswordManager, ins installer.Installer) (*PasswordManagerCLI, error) {
	return &PasswordManagerCLI{
		commands: map[string]*Command{
			"get":    {Name: "get", ArgsCount: 1, Usage: getCmdUsage},
			"set":    {Name: "set", ArgsCount: 2, Usage: setCmdUsage},
			"delete": {Name: "delete", ArgsCount: 1, Usage: deleteCmdUsage},
			"setup":  {Name: "setup", ArgsCount: 0, Usage: ""},
			"help":   {Name: "help", ArgsCount: 0, Usage: ""},
			"exit":   {Name: "exit", ArgsCount: 0, Usage: ""},
		},
		pm:  pm,
		ctx: ctx,
		ins: ins,
	}, nil
}

func (c *PasswordManagerCLI) CmdList() []string {
	keys := make([]string, 0, len(c.commands))
	for k := range c.commands {
		keys = append(keys, k)
	}
	return keys
}

func (c *PasswordManagerCLI) Execute(cmd string, args ...string) error {
	if _, ok := c.commands[cmd]; !ok {
		fmt.Println(commandNotFound)
		return common.ErrCmdNotFound
	}

	if requiredArgs := c.commands[cmd].ArgsCount; requiredArgs > len(args) {
		fmt.Println(c.commands[cmd].Usage)
		return common.ErrCmdWrongUsage
	}

	switch cmd {
	case "get":
		c.get(args[0])
	case "set":
		c.set(args[0], args[1])
	case "delete":
		c.delete(args[0])
	case "setup":
		c.setup()
	case "help":
		c.help()
	}

	return nil
}
