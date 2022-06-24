package opts

import (
	"github.com/spf13/cobra"
	"encoding/json"
	"fmt"

	"github.com/1920853199/passwd/service"
	"github.com/1920853199/passwd/util"

)

// Cmd
func NewDelCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "del",
		Short: "del account by key",
		Long:  `del account by key`,
		Run:   RunDel,
	}

	return cmd
}

func RunDel(cmd *cobra.Command, args []string) {
	config, err := util.GetConfigByClient()
	if err != nil {
		cmd.Println(err.Error())
		return
	}

	data, err := util.Post(fmt.Sprintf("%s%s", config.Addr, config.Api), service.ExecuteParams{
		Param: map[string]service.Args{
			"del": {
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
	if ret.Code == 200 {
		fmt.Printf("%s%s\n","\x1b[32m",ret.Msg)
	}else{
		fmt.Printf("%s%s\n","\x1b[31m",ret.Msg)
	}

}
