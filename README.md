cacti-go-tools
==============
Tools for statistic acquisition from services for cacti

You can find information on setup in the wiki
* [NGINX](https://github.com/eservicesgreece/cacti-go-tools/wiki)

## Usage
```bash
usage: cacti-go-tools.exe [<flags>] <command> [<args> ...]

Flags:
  -h, --help     Show context-sensitive help (also try --help-long and
                 --help-man).
      --version  Show application version.

Commands:
  help [<command>...]
    Show help.

  url <Status URL>
    Acquisition URL

  config
    Show Configuration

  engine <engine type>
    Acquisition Engine
```

## Configuration
The tool retrieves its configuration from the cacti-go-tools.json file, in normal operation you will not need to pass any other parameters other than which engine you want it to use.

### Testing the Setup
You can always test if your setup is working by invocking the tool directly from the command line
```bash
cacti-go-tools engine nginx
```
This should show something similar to this:
```
1
181
181
172
0
1
0
```
In case there is an error the tool will pass that error directly to the result, for http errors you will receive http/{http status code}
Error Example:
`http/403 Forbidden`