# pp-go

Pretty print command written in GO.

## Usage

JSON Parse

```
pp '{"key1": "val1", "key2": "val2"}'
```

```
less test.json | pp
```

```
pp -file=test.json
```

## Build

To make `pp` command, run following command.

```
go build -o ./bin/pp
```
