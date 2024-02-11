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
- implement algebra, conds, boolean alg
- implement function call
- implement control flow
    - make sure to handle early returns
- standard functions on wrapped value
