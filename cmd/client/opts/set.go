package opts

import (
	"github.com/spf13/cobra"
		"encoding/json"
	"fmt"

	"github.com/1920853199/passwd/service"
	"github.com/1920853199/passwd/util"

)

// Cmd
func NewSetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set",
		Short: "set account info. `passwd-cli set [key] [username] [password] [ip] [remark]`",
		Long:  "set account info.`passwd-cli set [key] [username] [password] [ip] [remark]`",
		Run:   RunSet,
	}

	return cmd
}

func RunSet(cmd *cobra.Command, args []string) {
	config, err := util.GetConfigByClient()
	if err != nil {
		cmd.Println(err.Error())
		return
	}

	if len(args) != 5 {
		fmt.Printf("%s%s\n","\x1b[31m","Error: args error. example:passwd-cli set [key] [username] [password] [ip] [remark]")
		return
	}

	data, err := util.Post(fmt.Sprintf("%s%s", config.Addr, config.Api), service.ExecuteParams{

		Param: map[string]service.Args{
			"set": {
				Key: args[0],
				Value:service.Item{
					Username:args[1],
					Password:args[2],
					Ip:args[3],
					Remark:args[4],
				},
			},
		},
	}, config.Token)

	if err != nil {
		cmd.Println(err.Error())
		return
	}

	var ret service.Result
	json.Unmarshal(data, &ret)

	if ret.Code == 200 {
		fmt.Printf("%s%s\n","\x1b[32m",ret.Msg)
	}else{
		fmt.Printf("%s%s\n","\x1b[31m",ret.Msg)
	}
}