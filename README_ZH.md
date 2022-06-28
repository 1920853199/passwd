[中文](./README_ZH.md) 

[English](./README.md)

## Passwd

一个基于服务的密码管理工具，他的数据存在于你的服务端，你的任何客户端可以通过Token来获取你保存的信息。

### 安装使用

可以直接通过 `make install` 安装客户端以及服务端（需要有 Golang 环境）

#### 启动服务

直接执行 `./passwd-server`:


``` bash
badger 2022/06/24 09:56:42 INFO: All 0 tables opened in 0s
badger 2022/06/24 09:56:42 INFO: Discard stats nextEmptySlot: 0
badger 2022/06/24 09:56:42 INFO: Set nextTxnTs to 7

---------------------------------------------------------------------------------------------------------------------------------------------------------
|        ____                               __                                                                                                          |
|       / __ \____ ____________      ______/ /                                                                                                          |
|      / /_/ / __/ / ___/ ___/ | /| / / __  /                                                                                                           |
|     / ____/ /_/ (__  |__  )| |/ |/ / /_/ /                                                                                                            |
|    /_/    \__,_/____/____/ |__/|__/\__,_/ v0.0.1                                                                                                      |
|                                                                                                                                                       |
|    Listening and serving HTTP on 0.0.0.0:22622                                                                                                        |
|                                                                                                                                                       |
|    Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IpXVCJ9.eyJqdGkiOiIxZTphZjpMDpmMDo1ZTozMyIsInN1YiI6InBhc3N3ZCJ9.dPkQ-ButAVlANDq89sPJ_UbyYJ80v6NQi68JBoE-ug      |
|                                                                                                                                                       |
---------------------------------------------------------------------------------------------------------------------------------------------------------

```
你必须保存好 `Token` 以提供给客户端使用 

服务端的启动参数有三个：
1. `path` 数据保存的位置，默认是用户目录下的 `.passwd` 文件夹
2. `ip` 访问IP范围，默认 `0.0.0.0`
3. `port` 服务的端口，默认 `22622` 

#### 客户端

启动客户端之前，必须修改客户端的配置文件。文件默认读取 `~/.passwd/comfig.json`

如：

```
{
    "Addr":"http://127.0.0.1:22622",
    "Api":"/v1/execute", // Cannot be modified
    "token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IpXVCJ9.eyJqdGkiOiIxZTphZjpMDpmMDo1ZTozMyIsInN1YiI6InBhc3N3ZCJ9.dPkQ-ButAVlANDq89sPJ_UbyYJ80v6NQi68JBoE-ug"
}
```

您需要把`Addr`修改成你服务端的地址以及端口，`Token`为上面服务启动时候获取到的。

你可以看到： 

```
./passwd-cli
```

```bash

A service-based password management tool.

Usage:
  passwd-cli [command]

Available Commands:
  all         get all account list
  clear       clear all account
  del         del account by key
  get         get account by key
  help        Help about any command
  set         set account info. `passwd-cli set [key] [username] [password] [ip] [remark]`
  ssh         get account ssh string for: `ssh root@127.0.0.1`

Flags:
  -h, --help   help for passwd-cli

Use "passwd-cli [command] --help" for more information about a command.
```

#### 例子

<img src="./static/all.png" width=800>

<img src="./static/get.png" width=800>

<img src="./static/set.png" width=800>

<img src="./static/ssh.png" width=800>

