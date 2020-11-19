# pp-go

`pp` is pretty print command written in Go.

## Usage

### JSON Parse

```shell script
pp '{"key1": "val1", "key2": "val2"}'
```

```shell script
less test.json | pp
```

```shell script
pp -file=test.json
```

### Options

The `pp` command has some options.

|Option|Explanation|default|
|:-:|:-:|:-:|
|file|Input JSON file.||
|indent|The number of space to indent.|2|


```shell script
pp -indent=4 '{"key1": "val1", "key2": "val2"}'
```

```shell script
less test.json | pp -indent=4
```

```shell script
pp -file=test.json -indent=4
```

To show help, run following command.

```shell script
pp -h
```

## Build

To make the `pp` command, run following command. Then add `path/to/bin` to your $PATH.

```
go build -o ./bin/pp
```
