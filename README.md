# Go JavaScript Engine Benchmarks

Performance benchmarks for three JavaScript engines in Go.

## Engines Tested

- **[Goja](https://github.com/dop251/goja)**: Pure Go implementation of ECMAScript 5.1
- **[ModerncQuickJS](https://gitlab.com/modernc.org/quickjs)**: QuickJS using [ccgo](https://pkg.go.dev/modernc.org/ccgo) (C-to-Go translator) with [mmap memory](https://pkg.go.dev/modernc.org/memory)
- **[QJS](https://github.com/fastschema/qjs)**: QuickJS compiled to WebAssembly


## Factorial Benchmark Results

| Iteration | GOJA | ModerncQuickJS | QJS |
| --- | --- | --- | --- |
| 1 | 1.128s | 1.897s | 737.635ms |
| 2 | 1.134s | 1.936s | 742.670ms |
| 3 | 1.123s | 1.898s | 738.737ms |
| 4 | 1.120s | 1.900s | 754.692ms |
| 5 | 1.132s | 1.918s | 756.924ms |
| Average | 1.127s | 1.910s | **746.132ms** |
| Total | 5.637s | 9.549s | **3.731s** |
| Speed | 1.51x | 2.56x | 1.00x |

*Benchmarks run on AMD Ryzen 7 7840HS, 32GB RAM, Linux*

## V8v7 Benchmark Results

| Metric | GOJA | ModerncQuickJS | QJS |
| --- | --- | --- | --- |
| Richards | 345 | 189 | **434** |
| DeltaBlue | 411 | 205 | **451** |
| Crypto | 203 | 305 | **393** |
| RayTrace | 404 | 347 | **488** |
| EarleyBoyer | 779 | 531 | **852** |
| RegExp | **381** | 145 | 142 |
| Splay | 1289 | 856 | **1408** |
| NavierStokes | 324 | 436 | **588** |
| Score (version 7) | 442 | 323 | **498** |
| Duration (seconds) | 78.349s | 97.240s | **72.004s** |

*Benchmarks run on AMD Ryzen 7 7840HS, 32GB RAM, Linux*

## What Gets Tested

### Factorial Benchmark

Calculates `factorial(10)` one million times using recursion. This tests how fast each engine handles computation and function calls. Each engine runs 5 times with alternating order to reduce bias.

### V8v7 Benchmark Suite

A standard JavaScript benchmark from the V8 project. Includes these tests:

- **Richards**: OS kernel simulation
- **DeltaBlue**: Constraint solving
- **Crypto**: Cryptographic operations
- **RayTrace**: 3D rendering
- **EarleyBoyer**: Parser and logic
- **RegExp**: Regular expressions
- **Splay**: Tree operations
- **NavierStokes**: Fluid dynamics

Each engine gets a score for each test and an overall score.

## Why Only Time Is Measured

Memory usage cannot be compared fairly between these engines. Here's why:

| Engine | Memory Type | Visible to Go |
|--------|-------------|---------------|
| Goja | Go heap and stack | Yes |
| QJS | WASM linear memory | No |
| ModerncQuickJS | mmap allocations | No |

Goja uses normal Go memory that shows up in `runtime.MemStats`. The other two use memory that Go cannot see. This makes memory comparisons meaningless.

Only execution time can be compared fairly across all three engines.

## How to Run

Clone the repository:

```bash
git clone http://github.com/ngocphuongnb/go-js-engines-benchmark.git
cd go-js-engines-benchmark
```

Run the factorial benchmark:

```bash
cd factorial
go run .
```

Run the V8v7 benchmark:

```bash
cd arewefastyet-v8v7
go run .
```

## How Each Engine Works

### Goja
```
┌─────────────────────────┐
│   Goja JS Engine        │
│  ┌──────────────────┐   │
│  │  Go Memory       │   │  <-  Go can see this
│  │  (Heap, Stack)   │   │
│  └──────────────────┘   │
└─────────────────────────┘
```

Written entirely in Go. Memory is managed by Go's garbage collector.

### ModerncQuickJS
```
┌─────────────────────────┐
│  QuickJS (ccgo)         │  <-  C code translated to Go
│  ┌──────────────────┐   │
│  │  mmap Memory     │   │  <-  Go cannot see this
│  │  (modernc.org/   │   │
│  │   memory)        │   │
│  └──────────────────┘   │
└─────────────────────────┘
```

Uses ccgo to translate C code to Go. Memory is allocated via mmap, which bypasses Go's memory tracking.

### QJS
```
┌─────────────────────────┐
│   QJS Wrapper (Go)      │
└─────────────────────────┘
           |
┌─────────────────────────┐
│   Wazero Runtime        │
│  ┌──────────────────┐   │
│  │ WASM Memory      │   │  <-  Go cannot see this
│  └──────────────────┘   │
└─────────────────────────┘
```

Runs QuickJS compiled to WebAssembly. Uses Wazero as the WebAssembly runtime. Memory is inside the WASM module.

## Keeping Tests Fair

The factorial benchmark uses several techniques to ensure fair comparison:

1. Runs garbage collection before each test
2. Waits 10ms before measuring, 100ms between tests
3. Alternates engine order every other iteration
4. Runs each engine 5 times and averages results
5. Measures only execution time, not startup (WASM compilation happens before timing starts)

## Contributing

Contributions are welcome. When reporting issues, include:

- Go version (`go version`)
- OS and CPU (`uname -a`)
- Full benchmark output
- System load and background processes

To add a benchmark:

1. Implement the `JSEngine` interface
2. Add it to `engines.Engines()`
3. Follow the existing structure

## License

Provided as-is for testing purposes. Engine licenses:
- Goja: MIT
- QJS: MIT
- ModerncQuickJS: BSD-3-Clause
