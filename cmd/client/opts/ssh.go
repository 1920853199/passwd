package opts

import (
	"encoding/json"
	"fmt"

	"github.com/1920853199/passwd/service"
	"github.com/1920853199/passwd/util"
	"github.com/spf13/cobra"
)

// Cmd
func NewSshCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ssh",
		Short: "get account ssh string for: `ssh root@127.0.0.1`",
		Long:  "get account ssh string for: `ssh root@127.0.0.1`",
		Run:   RunSsh,
	}

	return cmd
}

func RunSsh(cmd *cobra.Command, args []string) {

	config, err := util.GetConfigByClient()
	if err != nil {
		cmd.Println(err.Error())
		return
	}

	data, err := util.Post(fmt.Sprintf("%s%s", config.Addr, config.Api), service.ExecuteParams{
		Param: map[string]service.Args{
			"get": {
				Key: args[0],
			},
		},
	}, config.Token)

	if err != nil {
		cmd.Println(err.Error())
		return
	}

	var ret service.Result
	json.Unmarshal(data, &ret)

	util.PrintlnSSH(ret.Data, args[0])

}
