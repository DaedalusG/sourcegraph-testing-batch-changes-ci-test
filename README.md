# Batch Changes CI Test

Tests CI within a batch spec.

## CI in batch spec

There are two applications in this repo - a [Go App](./go-app) and a [Rust App](./rust-app). Each app has a failing
test. To test out batch changes running CI, there are two batch specs.

* [Spec to fix tests](./.batchchanges/fix-tests.yml)
* [Spec that fails CI](./.batchchanges/broken-tests.yml)

Each spec runs CI after doing the "main" logic of the batch change.

```yaml
  # Run CI
  - run: cd go-app && go test ./...
    container: golang:latest
  - run: cargo test --manifest-path rust-app/Cargo.toml
    container: rust:latest
```

### Example CI Failure

In the case of [broken-tests.yml spec](./.batchchanges/broken-tests.yml), the spec will fail with the following message,

#### Local

```text
✅ Executing... (1/1, 1 errored)  ████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████  100%
│                                                                                                                                                                                                               
└── github.com/sourcegraph-testing/bc-ci-test  go: downloading github.com/stretchr/testify v1.8.0                                                                                                                            4s

❌ Error:
   github.com/sourcegraph-testing/bc-ci-test:
   run: cd go-app && go test ./...
   container: golang:latest
   
   standard out:
        --- FAIL: Test1 (0.00s)
            main_test.go:10: 
                        Error Trace:    /work/go-app/main_test.go:10
                        Error:          Not equal: 
                                        expected: "foo"
                                        actual  : "faz"
                                        
                                        Diff:
                                        --- Expected
                                        +++ Actual
                                        @@ -1 +1 @@
                                        -foo
                                        +faz
                        Test:           Test1
        FAIL
        FAIL    github.com/sourcegraph-testing/bc-ci-test      0.003s
        FAIL
   
   standard error:
        go: downloading github.com/stretchr/testify v1.8.0
        go: downloading gopkg.in/yaml.v3 v3.0.1
        go: downloading github.com/pmezard/go-difflib v1.0.0
        go: downloading github.com/davecgh/go-spew v1.1.1
   
   Command failed with exit code 1.
```

#### Server-side

![img](./img/ssbc-failed-ci.png)


## CI in batch change

Running the [basic.yml Spec](./.batchchanges/basic.yml) does not run CI at batch spec time. Instead, it depends on GitHub
to run the CI checks.

### Example Failure

![img](./img/change-failed.png)
Hello World
