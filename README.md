# Monkey lang

An interpreter written in Go, based on the https://interpreterbook.com.

## Running the REPL locally

You just need to have go installed. Then run:

```
git clone https://github.com/juandspy/monkey-lang.git
go run main.go
```

and start writing Monkey lang code. The results of the statements will be printed in the stdout. For example:

```
>> 3 * 7
21
>> let x = 3 * 7
>> x
21
```

Note that some statements like variable bindings don't print anything in the stdout.

## Language specs

### Types

There are 5 types supported:

- Booleans: `true` or `false`
- Integers: `1`, `-1`, `12345`...
- Strings: `"Hello World"`
- Arrays: `[1, 2, 3]`. You can access a given position of an array by using indexes: `[1, 2, 3][1]` or `myArray[1]`.
- Hashes: `{"a": 1, 5: "test", true: "bool"}`

### Operators

#### Prefix expressions

- Bang (`!`): it takes any input and returns the opposite. For example `!true = false` and `!5 = false`, as `5` acts as "truthy". However, `!!5` would be `true` as it's the same as `!false`.
- Minus (`-`): changes the sign of an integer e.g. `-5`.

#### Infix expressions

- Arithmetic expressions: `(1 + 2) * 3 / 4`
- Comparisons: `3 != 2`
- Conditionals: `if (3 == 3) {"equals"} else {"not equals"}`

### Variables

You can define a variable by using `let` statements, e.g. `let x = 3`. You can also bind expressions: `let x = 3 * 7`.

#### Functions

You can bind functions to variables using the `let` statement:

```
let sum = fn(x, y) {return x + y}; sum(1, 2)
```

You can also build recursive functions:
```
>> let fib = fn(n) { if (n < 2) { return n; } else {return fib(n-1) + fib(n-2); } }
>> fib(20)
6765
```

#### Builtin functions

There is a set of builtin functions available which are defined in [builtins.go](evaluator/builtins.go):
- `len`: returns the length of a string or array.
- `first`: returns the first element of an array.
- `last`: returns the last element of an array.
- `rest`: returns all the elements except the first one.
- `push`: appends an item to an array.
- `puts`: output to stdout.
