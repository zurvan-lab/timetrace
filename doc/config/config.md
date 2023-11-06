# Config

In this part of documentation we are going to explain config in ttrace.

### config file

In ttrace we have config in `yaml` format, you can see default config example [here](../../config/config.yaml).
Now we explain each part of config one by one.

#### name

First of all we have a name, which is name of instance you are running (you can use it on application layer and get it by commands).
This is name field in `config.yaml`:

```yml
name: time_trace
```

#### server

In server part you can config this two: which ip? which port? for listening and serving the TCP server.

This is how it looks in `config.yaml`:

```yml
server:
  listen: localhost
  port: "7070"
```

#### log

This part will help you to config log stuff (levels, saving path and...).

How it looks in `config.yaml`:

```yml
log:
  write_to_file: true
  path: log.ttrace
```

#### users

In users part, you define who can access the database and set permission for them. `name` and `password` field is user pass for connecting (required in `CON` command).

In `cmds` you provide which command in [TQL](../TQL/) they have access to. It can be a list of commands or just a `'*'` which means all.
Also the `CON` command is open for everyone.

example:

```yml
users:
- name: root
  password: super_secret_password
  cmds:
  - '*'

# Also you can use this for multiple users and limited commands
# users:
#   - name: root
#     password: super_secret_password
#     cmd:
#       - '*' # all commands.
#   - name: developer
#     password: secret_password
#     cmd:
#       - 'GET'
#       - 'PUSH'
#       - 'DEL'
```
