package main

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	completionLong = `Output shell completion code for the specified shell.
The shell code must be evaluated to provide interactive completions
of kni-install commands.

For examples of loading/evaluating the completions see:
  kni-install completion bash --help`

	completionExampleBash = `  # Installing bash completion on macOS using homebrew
  ## If running Bash 3.2 included with macOS
      brew install bash-completion
  ## or, if running Bash 4.1+
      brew install bash-completion@2
  ## If you've installed via other means, you may need add the completion to your completion directory
      kni-install completion bash > $(brew --prefix)/etc/bash_completion.d/openshift-install

  # Installing bash completion on Linux
  ## Load the kni-install completion code for bash into the current shell
      source <(kni-install completion bash)
  ## Write bash completion code to a file and source it from .bash_profile
      kni-install completion bash > ~/.openshift-install/completion.bash.inc
      printf "
        # Kubectl shell completion
        source '$HOME/.kni-install/completion.bash.inc'
        " >> $HOME/.bash_profile
      source $HOME/.bash_profile`

	completionExampleZsh = `# Load the kni-install completion code for zsh[1] into the current shell
      source <(kni-install completion zsh)
  # Set the kni-install completion code for zsh[1] to autoload on startup
      kni-install completion zsh > "${fpath[1]}/_openshift-install"`
)

func newCompletionCmd() *cobra.Command {
	completionCmd := &cobra.Command{
		Use:   "completion",
		Short: "Outputs shell completions for the kni-install command",
		Long:  completionLong,
	}

	bashCompletionCmd := &cobra.Command{
		Use:     "bash",
		Short:   "Outputs the bash shell completions",
		Example: completionExampleBash,
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, _ []string) error {
			return cmd.Root().GenBashCompletion(os.Stdout)
		},
	}
	completionCmd.AddCommand(bashCompletionCmd)

	// The zsh completions didn't work for crawford, so commenting them out
	//zshCompletionCmd := &cobra.Command{
	//Use:     "zsh",
	//Short:   "Outputs the zsh shell completions",
	//Example: completionExampleZsh,
	//RunE: func(cmd *cobra.Command, _ []string) error {
	//return cmd.Root().GenZshCompletion(os.Stdout)
	//},
	//}
	//completionCmd.AddCommand(zshCompletionCmd)

	return completionCmd
}
