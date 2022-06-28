package util

import (
	"fmt"
)

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

const tableHead string = `
---------------------------------------------------------------------------------------------------------------------------------------------
|     KEY       |          Username            |           Password           |              Ip              |            Remark            |
---------------------------------------------------------------------------------------------------------------------------------------------`

const tableLine string = `
| %-14s| %-29s| %-29s| %-29s| %-25s|`

const tableFooter string = `
---------------------------------------------------------------------------------------------------------------------------------------------`

func PrintlnData(v interface{}, key string) {

	str := fmt.Sprintf(tableHead)

	// key 不为空为GET结果
	if key != "" {
		var item = v.(map[string]interface{})
		str += fmt.Sprintf(tableLine, key, item["username"], item["password"], item["ip"], item["remark"])
		str += tableFooter
	} else {
		var item = v.([]interface{})

		for _, val := range item {
			tmpVal := val.(map[string]interface{})
			key := tmpVal["key"]
			it := tmpVal["value"].(map[string]interface{})
			str += fmt.Sprintf(tableLine, key, it["username"], it["password"], it["ip"], it["remark"])
			str += tableFooter
		}
	}

	fmt.Println(str)
}

func PrintlnSSH(v interface{}, key string) {
	const sshTableHead string = `
-----------------------------------------------------------------------------------------------
|     KEY       |                     SSH                      |           Password           |
-----------------------------------------------------------------------------------------------`

	const sshTableLine string = `
| %-14s| %-45s| %-29s|`

	const sshTableFooter string = `
-----------------------------------------------------------------------------------------------`

	str := fmt.Sprintf(sshTableHead)
	var item = v.(map[string]interface{})
	str += fmt.Sprintf(sshTableLine, key, "ssh "+item["username"].(string)+"@"+item["ip"].(string), item["password"])
	str += sshTableFooter
	fmt.Println(str)
}
