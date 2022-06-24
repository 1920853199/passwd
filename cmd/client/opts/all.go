package opts

import (
	"encoding/json"
	"fmt"

	"github.com/1920853199/passwd/service"
	"github.com/1920853199/passwd/util"
	"github.com/spf13/cobra"
)

// Cmd
func NewAllCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "all",
		Short: "get all account list",
		Long:  `get all account list`,
		Run:   RunAll,
	}

	return cmd
}

func RunAll(cmd *cobra.Command, args []string) {

	config, err := util.GetConfigByClient()
	if err != nil {
		cmd.Println(err.Error())
		return
	}

	data, err := util.Post(fmt.Sprintf("%s%s", config.Addr, config.Api), service.ExecuteParams{
		Param: map[string]service.Args{
			"all": {
				Key: "",
			},
		},
	}, config.Token)

	if err != nil {
		cmd.Println(err.Error())
		return
	}

	var ret service.Result
	json.Unmarshal(data, &ret)
	util.PrintlnData(ret.Data, "")

}
