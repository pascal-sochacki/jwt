package jwt

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var InspectCmd = &cobra.Command{
	Use:     "inspect",
	Aliases: []string{"ins"},
	Short:   "inspect a jwt",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		jwt := args[0]
		parts := strings.Split(jwt, ".")

		println(len(parts))
		if len(parts) != 3 {
			fmt.Println("jwt does not contain three parts")
			os.Exit(1)
		}

		for _, part := range parts {
			if l := len(part) % 4; l > 0 {
				part += strings.Repeat("=", 4-l)
			}

			decodedData, err := base64.URLEncoding.DecodeString(part)
			if err != nil {
				fmt.Printf("could not decode: %s\n", part)
				os.Exit(1)
			}

			var pretty bytes.Buffer
			err = json.Indent(&pretty, []byte(decodedData), "", "  ")
			if err != nil {
				fmt.Printf("could not pretty part: %s", part)
				os.Exit(1)
			}
			println(pretty.String())

		}

	},
}

func init() {
	rootCmd.AddCommand(InspectCmd)
}
