package http

import "github.com/spf13/cobra"

func GetCommand() *cobra.Command {
	return &cobra.Command{
		Use: "http",
		Run: func(_ *cobra.Command, _ []string) {
			main()
		},
	}
}

func main() {

}
