# ilang

Isaac's toy programming language

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

TODO:
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
- map/array indexing
- standard functions on wrapped value
- handle early returns
- custom error stat https://stackoverflow.com/questions/18132078/handling-errors-in-antlr4
- imports
