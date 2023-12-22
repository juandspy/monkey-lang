# Monkey lang

An interpreter written in Go, based on the https://interpreterbook.com.

## Running the REPL locally

You just need to have go installed. Then run:

```
git clone https://github.com/juandspy/monkey-lang.git
go run main.go
```

and start writing Monkey lang code.

## Language specs

### Operators

- Bang (!): it takes any input and returns the opposite. For example `!true = false` and `!5 = false`, as `5` acts as "truthy". However, `!!5` would be `true` as it's the same as `!false`.
