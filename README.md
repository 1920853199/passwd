## Passwd

A service-based password management tool.

### Usage

`make install` install the server and client for your pc.

#### server

`passwd-server` running the server:

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
You must write down the token information for the client to use.

#### You can specify the open port and the directory where the data is saved and your IP address.

```
flag.StringVar(&path, "path", util.GetConfigPath(), "Directory storing data.")
flag.StringVar(&host, "host", "0.0.0.0", "The ip on which to serve")
flag.IntVar(&port, "port", 22622, "The port on which to serve")
```

#### client

```passwd-cli``` running the client:

Before starting the client, you must modify the configuration file of the client.

```~/.passwd/comfig.json```

Example：
```
{
    "Addr":"http://127.0.0.1:22622",
    "Api":"/v1/execute", // Cannot be modified
    "token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IpXVCJ9.eyJqdGkiOiIxZTphZjpMDpmMDo1ZTozMyIsInN1YiI6InBhc3N3ZCJ9.dPkQ-ButAVlANDq89sPJ_UbyYJ80v6NQi68JBoE-ug"
}
```

```bash
account management client

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

#### Example

<img src="./static/all.png" width=800>

<img src="./static/get.png" width=800>

<img src="./static/set.png" width=800>

<img src="./static/ssh.png" width=800>

