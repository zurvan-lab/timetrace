# Usage of time trace

Currently you can run your timetrace instance using [ttrace CLI](../../cmd/main.go). 
You can download latest version based on your OS and CPU arch [here](https://github.com/zurvan-lab/TimeTrace/releases).

## checking installation

You can check if you are installed ttrace properly:
```sh
ttrace --version
```

And:
```sh
ttrace --help
```

## run an instance

To run an new instance you can simply make a `config.yaml` file as config first.

> NOTE: see config details [here](../config/config.md).

Then run your instance:

```sh
ttrace run -c {path-to-your-config.yaml}
```

## connecting with REPL

You can use ttrace CLI to connect to your instance and execute TQL queries:

```sh
ttrace connect -u username -p password -a remote-address
```

## other commands

Check ttrace `--help` to find-out more commands.

## other usage ways

You can also implement or use a timetrace client or drive which have timetrace an TQL protocol and language implemented to run or interact with a timetrace instance.
