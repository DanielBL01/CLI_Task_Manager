# CLI Task Manager

```bash
This is your Task Manager

Usage:
  Task [command]

Available Commands:
  add         Add a task
  done        Finish a task
  help        Help about any command
  list        List tasks

Flags:
  -h, --help   help for Task

Use "Task [command] --help" for more information about a command.
```

This project uses **Bolt** which is a pure Go key/value store. Bolt will store your tasks in a local database called "tasks.db" in your home directory

To find the home directory for Mac OS/Linux or Windows (whatever you use to run this binary) we use the **go-homedir** package to avoid any issues


