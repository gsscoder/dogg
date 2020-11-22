# dogg

Simple program to poll processes for resources consumption and existence. On request it can revive a process that dies. Built with [gopsutil](https://github.com/shirou/gopsutil) and [go-yaml](https://github.com/go-yaml/yaml).

## Build

**NOTE**: Go 1.13 or higher is required.

```sh
# clone the repository
$ git clone https://github.com/gsscoder/dogg.git

# change the working directory
$ cd dogg

# build the executable
$ sh ./build.sh

# test if it works
$ ./artifacts/dogg -version
```

## Configuration

**dogg.yaml**:

```yaml
constraints:
  processGroups:
    - process: chrome
      match: chrome.exe|Chrome\.app
      cpu: 0.5
      mem: 1
    - process: skype
      match: Skype.exe|Skype\.app
      cpu: 0.5
      mem: 1
```

Each process group is identified by a name (like `chrome`) and all processes bound to it are selected using one or more regular expression. The match is done using the executable path of the process (you can easly discover it with command `ps -A`). Defined constraints are expressed in percentage and are checked for all processes of a group.

## Usage

```sh
$ ./artifacts/dogg -help
dogg: Polls processes for resources consumption and existence
Version: v0.4.0
usage:
  -poll duration
    	defines polling interval (default 2s)
  -restart
    	restart process if terminated
  -version
    	displays version information
  -wait duration
    	time to wait before polling again after a restart (default 5s)

$ ./artifacts/dogg -restart 2>>dogg.log &
[1] 14593
$ tail -f dogg.log
2020/11/22 11:48:06 skype: CPU constraint of 0.50% violated by +6.40%
2020/11/22 11:48:06 skype: Memory constraint of 1.00% violated by +0.59%
2020/11/22 11:48:08 chrome: CPU constraint of 0.50% violated by +0.02%
2020/11/22 11:48:08 chrome: CPU constraint of 0.50% violated by +6.60%
...
$ sudo kill 14593
[1]  + terminated  ./artifacts/dogg -restart 2>> dogg.log 
```

### Notes

- Restarted processes will die, when `dogg` terminates in case it wasn't started in background.
- Tesred on **macOS** and **Windows**.
- For now it's a [Go language](https://golang.org/) learning project.
