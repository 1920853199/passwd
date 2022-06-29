package util

import (
	"fmt"
)

type STable struct {
	Key      string
	Username string
	Password string
	Ip       string
	Remark   string
}

type SSLTable struct {
	Key      string
	SSH      string
	Password string
}

const versionString string = `
---------------------------------------------------------------------------------------------------------------------------------------------------------
|        ____                               __                                                                                                          |
|       / __ \____ ____________      ______/ /                                                                                                          |
|      / /_/ / __/ / ___/ ___/ | /| / / __  /                                                                                                           |
|     / ____/ /_/ (__  |__  )| |/ |/ / /_/ /                                                                                                            |
|    /_/    \__,_/____/____/ |__/|__/\__,_/ v0.0.1                                                                                                      |
|                                                                                                                                                       |
|    %-147s|
|                                                                                                                                                       |
|    Token: %-140s|
|                                                                                                                                                       |
---------------------------------------------------------------------------------------------------------------------------------------------------------
`

func Println(args []string) {
	fmt.Println(fmt.Sprintf(versionString, args[0], args[1]))
}

func PrintlnData(v interface{}, key string) {

	dataTable := make([]STable, 0)
	// key 不为空为GET结果
	if key != "" {
		var item = v.(map[string]interface{})
		dataTable = append(dataTable, STable{
			key,
			item["username"].(string),
			item["password"].(string),
			item["ip"].(string),
			item["remark"].(string),
		})

	} else {
		var item = v.([]interface{})
		for _, val := range item {
			tmpVal := val.(map[string]interface{})
			key := tmpVal["key"]
			it := tmpVal["value"].(map[string]interface{})

			dataTable = append(dataTable, STable{
				key.(string),
				it["username"].(string),
				it["password"].(string),
				it["ip"].(string),
				it["remark"].(string),
			})
		}
	}

	OutputLine(dataTable)

	//fmt.Println(str)
}

func PrintlnSSH(v interface{}, key string) {
	var item = v.(map[string]interface{})
	dataTable := make([]SSLTable, 0)
	dataTable = append(dataTable, SSLTable{
		key,
		fmt.Sprintf("ssh %s@%s", item["username"].(string), item["ip"].(string)),
		item["password"].(string),
	})

	Output(dataTable)
}
