package jwt

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

type Body struct {
	IssuedAt int `json:"iat"`
}

var InspectCmd = &cobra.Command{
	Use:     "inspect",
	Aliases: []string{"ins"},
	Short:   "inspect a jwt",
	Run: func(cmd *cobra.Command, args []string) {
		var jwt string
		if len(args) == 1 {
			jwt = args[0]
		} else {
			scanner := bufio.NewScanner(os.Stdin)
			for scanner.Scan() {
				jwt = scanner.Text()
			}
		}
		parts := strings.Split(jwt, ".")

		if len(parts) != 3 {
			fmt.Println("jwt does not contain three parts")
			os.Exit(1)
		}

		if l := len(parts[1]) % 4; l > 0 {
			parts[1] += strings.Repeat("=", 4-l)
		}

		decodedData, err := base64.URLEncoding.DecodeString(parts[1])
		if err != nil {
			fmt.Printf("could not decode: %s\n", parts[1])
			os.Exit(1)
		}
		var body Body
		json.Unmarshal(decodedData, &body)

		var pretty bytes.Buffer
		err = json.Indent(&pretty, []byte(decodedData), "", "  ")
		if err != nil {
			fmt.Printf("could not pretty part: %s", parts[1])
			os.Exit(1)
		}
		println(pretty.String())

		iatTime := time.Unix(int64(body.IssuedAt), 0)
		germanFormat := iatTime.Format("02.01.2006 15:04:05 MST")

		fmt.Printf("Issued At (iat): %s\n", germanFormat)

	},
}

func init() {
	rootCmd.AddCommand(InspectCmd)
}
