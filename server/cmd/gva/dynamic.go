package main

import "github.com/spf13/cobra"

func attachDynamicCommands(root *cobra.Command, manifest Manifest, cfg *CliConfig) {
	for _, command := range manifest.Commands {
		cmdDef := command
		cmd := &cobra.Command{
			Use:   cmdDef.Name,
			Short: cmdDef.Summary,
			Long:  cmdDef.Description,
			RunE: func(cmd *cobra.Command, args []string) error {
				return runManifestCommand(cmdDef, cfg, cmd)
			},
		}
		addManifestFlags(cmd, cmdDef.Parameters)
		root.AddCommand(cmd)
	}
}
