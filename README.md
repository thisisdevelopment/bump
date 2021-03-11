# bump
version bump tool compatible with govvv

## Install
```go get github.com/thisisdevelopment/bump```

## Help
```bump -h```

## Notes
Standard operation is to bump a version section indicated with the -b option.

## Examples:

| Command                                                   | Output |
| --------------------------------------------------------- |:------:|
| bump -f -b patch                                          | 0.0.1 |
| bump -b major                                               | 1.0.1 |
| bump -b minor -c                                           | 1.1.1-4ea0adbe |
| bump                                                      | 1.2.1-4ea0adbe |
