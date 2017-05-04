# jitjson
Faster JSON marshaling in Go by code gening encoders that don't rely on reflection

## Usage
```
usage: jitjson [<flags>] <command> [<args> ...]

Finds structs with json tags and generates efficient (non-reflection) based JSON
encoders

Flags:
  --help  Show context-sensitive help (also try --help-long and --help-man).

Commands:
  help [<command>...]
    Show help.

  gen [<root-dir>]
    Generate json encoders

  list [<flags>] [<root-dir>]
    List eligible structs that were found

  clean [<root-dir>]
    Delete all auto-generated source files

  dump [<flags>] [<root-dir>]
    Dump auto-generated source code

  files [<root-dir>]
    List json encoder files that have been generated
```

