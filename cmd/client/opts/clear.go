package opts

import (
	"github.com/spf13/cobra"

	"encoding/json"
	"fmt"

	"github.com/1920853199/passwd/service"
	"github.com/1920853199/passwd/util"
)

// Cmd
func NewClearCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clear",
		Short: "clear all account",
		Long:  `clear all account`,
		Run:   RunClear,
	}

	return cmd
}

func RunClear(cmd *cobra.Command, args []string) {
	config, err := util.GetConfigByClient()
	if err != nil {
		cmd.Println(err.Error())
		return
	}

	data, err := util.Post(fmt.Sprintf("%s%s", config.Addr, config.Api), service.ExecuteParams{
		Param: map[string]service.Args{
			"clear": {
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
	if ret.Code == 200 {
		fmt.Printf("%s\n", ret.Msg)
	} else {
		fmt.Printf("%s\n", ret.Msg)
	}
}
