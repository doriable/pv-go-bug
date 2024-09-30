# pv-go-bug

## To reproduce:

Uncomment out the code block under `// failing message` (lines 21-28), and run the file:

```
go run main.go
```

This produces a panic:

```
$ go run main.go
panic: interface conversion: *types.Err is not traits.Lister: missing method Add
```

The expectation would be that it valid to set `proto.Int32(5)` on the message and no panic
occurs. Similarly, when using a built-in rule, `not_in` (set on the `control` message), the
validation is successful.