package opts

import (
	"encoding/json"
	"fmt"

	"github.com/1920853199/passwd/service"
	"github.com/1920853199/passwd/util"
	"github.com/spf13/cobra"
)

// Cmd
func NewGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "get account by key",
		Long:  `get account by key`,
		Run:   RunGet,
	}

	return cmd
}

func RunGet(cmd *cobra.Command, args []string) {
	// 读取环境配置

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
	util.PrintlnData(ret.Data, args[0])

}
