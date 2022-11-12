package migrate

import "github.com/spf13/cobra"

func GetCommand() *cobra.Command {
	return &cobra.Command{
		Use: "migrate",
		Run: func(_ *cobra.Command, _ []string) {
			main()
		},
	}
}

func main() {

}
