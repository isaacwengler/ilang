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
