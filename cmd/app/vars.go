package app

import "github.com/spf13/cobra"

func VarsCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "variables",
		Aliases: []string{"vars"},
		Short:   "Show variables for active environment",
		RunE:    Contextualize(handler.Variables, handler.Panic),
	}

	cmd.Flags().StringP("service", "s", "", "Fetch variables accessible to a specific service")

	variablesAddCmd := &cobra.Command{
		Use:     "get key",
		Short:   "Get the value of a variable",
		RunE:    Contextualize(handler.VariablesGet, handler.Panic),
		Args:    cobra.MinimumNArgs(1),
		Example: "  botway variables get DISCORD_TOKEN",
	}

	cmd.AddCommand(variablesAddCmd)
	variablesAddCmd.Flags().StringP("service", "s", "", "Fetch variables accessible to a specific service")

	variablesSetCmd := &cobra.Command{
		Use:     "set key=value",
		Short:   "Create or update the value of a variable",
		RunE:    Contextualize(handler.VariablesSet, handler.Panic),
		Args:    cobra.MinimumNArgs(1),
		Example: "  botway variables set TELEGRAM_TOKEN=TOKEN",
	}

	cmd.AddCommand(variablesSetCmd)
	variablesSetCmd.Flags().StringP("service", "s", "", "Fetch variables accessible to a specific service")

	variablesDeleteCmd := &cobra.Command{
		Use:     "delete key",
		Short:   "Delete a variable",
		RunE:    Contextualize(handler.VariablesDelete, handler.Panic),
		Example: "  botway variables delete MY_KEY",
	}

	cmd.AddCommand(variablesDeleteCmd)
	variablesDeleteCmd.Flags().StringP("service", "s", "", "Fetch variables accessible to a specific service")

	return cmd
}
