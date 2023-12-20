# Monkey lang

An interpreter written in Go, based on the https://interpreterbook.com.

## Language specs

### Operators

- Bang (!): it takes any input and returns the opposite. For example `!true = false` and `!5 = false`, as `5` acts as "truthy". However, `!!5` would be `true` as it's the same as `!false`.
