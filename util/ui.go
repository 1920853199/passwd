package util

import (
	"fmt"
)

const versionString string = `%s%s
---------------------------------------------------------------------------------------------------------------------------------------------------------
|        ____                               __                                                                                                          |
|       / __ \____ ____________      ______/ /                                                                                                          |
|      / /_/ / __/ / ___/ ___/ | /| / / __  /                                                                                                           |
|     / ____/ /_/ (__  |__  )| |/ |/ / /_/ /                                                                                                            |
|    /_/    \__,_/____/____/ |__/|__/\__,_/ v0.0.1                                                                                                      |
|                                                                                                                                                       |
|    %-147s|
|                                                                                                                                                       |
|    Token: %s%-140s%s%s|
|                                                                                                                                                       |
---------------------------------------------------------------------------------------------------------------------------------------------------------
`

func Println(args []string) {
	fmt.Println(fmt.Sprintf(versionString, "\x1b[35m", "\x1b[1m", args[0], "\x1b[32m", args[1], "\x1b[35m", "\x1b[1m"))
}

const tableHead string = `%s%s
---------------------------------------------------------------------------------------------------------------------------------------------
|     KEY       |          Username            |           Password           |              Ip              |            Remark            |
---------------------------------------------------------------------------------------------------------------------------------------------`

const tableLine string = `
| %s%-14s%s%s| %s%-29s%s%s| %s%-29s%s%s| %s%-29s%s%s| %s%-25s%s%s|`

const tableFooter string = `
---------------------------------------------------------------------------------------------------------------------------------------------`

func PrintlnData(v interface{}, key string) {

	str := fmt.Sprintf(tableHead, "\x1b[35m", "\x1b[1m")

	// key 不为空为GET结果
	if key != "" {
		var item = v.(map[string]interface{})
		str += fmt.Sprintf(tableLine, "\x1b[32m", key, "\x1b[35m", "\x1b[1m", "\x1b[32m", item["username"], "\x1b[35m", "\x1b[1m", "\x1b[32m", item["password"], "\x1b[35m", "\x1b[1m", "\x1b[32m", item["ip"], "\x1b[35m", "\x1b[1m", "\x1b[32m", item["remark"], "\x1b[35m", "\x1b[1m")
		str += tableFooter
	} else {
		var item = v.([]interface{})

		for _, val := range item {
			tmpVal := val.(map[string]interface{})
			key := tmpVal["key"]
			it := tmpVal["value"].(map[string]interface{})
			str += fmt.Sprintf(tableLine, "\x1b[32m", key, "\x1b[35m", "\x1b[1m", "\x1b[32m", it["username"], "\x1b[35m", "\x1b[1m", "\x1b[32m", it["password"], "\x1b[35m", "\x1b[1m", "\x1b[32m", it["ip"], "\x1b[35m", "\x1b[1m", "\x1b[32m", it["remark"], "\x1b[35m", "\x1b[1m")
			str += tableFooter
		}
	}

	fmt.Println(str)
}

func PrintlnSSH(v interface{}, key string) {
	const sshTableHead string = `%s%s
-----------------------------------------------------------------------------------------------
|     KEY       |                     SSH                      |           Password           |
-----------------------------------------------------------------------------------------------`

	const sshTableLine string = `
| %s%-14s%s%s| %s%-45s%s%s| %s%-29s%s%s|`

	const sshTableFooter string = `
-----------------------------------------------------------------------------------------------`

	str := fmt.Sprintf(sshTableHead, "\x1b[35m", "\x1b[1m")
	var item = v.(map[string]interface{})
	str += fmt.Sprintf(sshTableLine, "\x1b[32m", key, "\x1b[35m", "\x1b[1m", "\x1b[32m", "ssh "+item["username"].(string)+"@"+item["ip"].(string), "\x1b[35m", "\x1b[1m", "\x1b[32m", item["password"], "\x1b[35m", "\x1b[1m")
	str += sshTableFooter
	fmt.Println(str)
}
