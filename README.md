# pp-go

`ppgo` is pretty print command written in Go.

## Usage

### JSON Parse

```shell script
ppgo '{"key1":"val1","key2":"val2","arr":[101,202,303,404]}'
```

Output:

```json
{
  "arr": [
    101,
    202,
    303,
    404
  ],
  "key1": "val1",
  "key2": "val2"
}
```

---
As other usage,

Pipeline:

```shell script
less test.json | ppgo
```

`-file` option:

```shell script
ppgo -file=test.json
```

### Options

The `ppgo` command has some options.

|Option|Explanation|default|
|:-:|:-:|:-:|
|file|Input JSON file.||
|indent|The number of space to indent.|2|

For example:

```shell script
ppgo -indent=4 '{"key1":"val1","key2":"val2"}'
```

```shell script
less test.json | ppgo -indent=4
```

```shell script
ppgo -file=test.json -indent=4
```

To show help, run following command.

```shell script
ppgo -h
```

## Build

To make the `ppgo` command, run following command. Then add `path/to/bin` to your $PATH.

```
go build -o ./bin/ppgo
```

## Test

To test, run `go test`.

```
go test -v
```
