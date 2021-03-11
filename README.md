# bump
version bump tool compatible with govvv

## Installation
```go get github.com/thisisdevelopment/bump```

## Getting Help
```bump -h```

## Notes
Standard operation is to bump a version section indicated with the -b option. The exception to this rule is the -s option, it will set the version as indicated.

## Examples:


| Command                                                   | Output |
| --------------------------------------------------------- |:------:|
| bump -f -b patch                                          | 0.0.1 |
| bump -b M                                                 | 1.0.1 |
| bump -b minor -c 4ea0adbe4dcf631f99c311ce0dfefedfa53b391f | 1.1.1-4ea0adbe4dcf631f99c311ce0dfefedfa53b391f |
| bump -b p                                                 | 1.1.2-4ea0adbe4dcf631f99c311ce0dfefedfa53b391f |
| bump -s 1                                                 | 1.0.0 |
| bump -s 1.2                                               | 1.2.0 |
| bump -s 1.2 -c 4ea0adbe4dcf631f99c311ce0dfefedfa53b391f   | 1.2.0-4ea0adbe4dcf631f99c311ce0dfefedfa53b391f |
| bump                                                      | 1.2.1-4ea0adbe4dcf631f99c311ce0dfefedfa53b391f |
