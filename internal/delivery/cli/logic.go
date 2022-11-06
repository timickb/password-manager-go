package cli

import (
	"fmt"
)

func (c *PasswordManagerCLI) get(secretName string) {
	sec, err := c.pm.Read(c.ctx, secretName)

	if err != nil {
		fmt.Printf("error occurred: %s", err.Error())
		return
	}

	fmt.Printf("secret: %s\n", sec)
}

func (c *PasswordManagerCLI) set(secretName string, secret string) {
	err := c.pm.Set(c.ctx, secretName, secret)

	if err != nil {
		fmt.Printf("error occured: %s\n", err.Error())
		return
	}

	fmt.Printf("secret was set.\n")
}

func (c *PasswordManagerCLI) delete(secretName string) {
	err := c.pm.Delete(c.ctx, secretName)

	if err != nil {
		fmt.Printf("error occured: %s\n", err.Error())
		return
	}

	fmt.Printf("secret was deleted.\n")
}

func (c *PasswordManagerCLI) setup() {
	err := c.pm.Setup(c.ctx, c.ins)

	if err != nil {
		fmt.Printf("error occured: %s\n", err.Error())
		return
	}

	fmt.Println("\nSetup completed. You can use the password manager.")
}

func (c *PasswordManagerCLI) help() {
	fmt.Println("CLI password manager")
	fmt.Println()

	fmt.Println("Usage:")
	fmt.Println("\t gpm <command> [arguments]")
	fmt.Println()

	fmt.Println("The commands are:")
	fmt.Println()

	fmt.Println("\t setup \t Setup password manager store")
	fmt.Println("\t set \t Create new or update an existing password")
	fmt.Println("\t get \t Show stored password")
	fmt.Println("\t delete \t Delete a password")
}
