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
- parse array and map
- implement algebra, conds, boolean alg
- implement control flow
- implement function call
- standard functions on wrapped value
