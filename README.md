# ilang

Toy programming language to experiment with writing an interpreter.

Next maybe in this repo, maybe in new ones:
- LSP server
- treesitter integration
- Static types
- Compiler? maybe
- Garbage collection

## Setup

Build generated code. Note: must have antlr installed and aliased to `antlr4`

```bash
./build.sh
```

Run tests

```bash
go test -v ./...
```

Run interpreter

```bash
go run .

# or
./build/ilang

# debug mode
./build/ilang -d
```

## Interpreter TODO

- ~~WrappedValue type~~
    - has map children with child bindings
    - implemented by all primitives
- ~~update current to use wrapped value~~
- ~~parse and impl array and map~~
- ~~parse and impl reassignments~~
- ~~parse and impl return~~
- ~~implement algebra, conds, boolean alg~~
- ~~implement control flow~~
- ~~refactor value types~~
- ~~function definition type~~
- ~~implement function call~~
- ~~command line with file and interactive mode~~
- ~~map/array indexing~~
- ~~handle early returns~~
- try/catch/throw
- custom error stat https://stackoverflow.com/questions/18132078/handling-errors-in-antlr4
- imports/exports
- standard functions on wrapped value
- fix bug where (i < nums.length()) tries to do i < nums before nums.length()
- simple package manager (un-versioned git repo curls? maybe)
