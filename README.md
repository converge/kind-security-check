# Kind security check

This is a simple tool to check if your Kubernetes ([Kind](http://kind.sigs.k8s.io)) cluster is secure.

## Current validations

- [x] [Default namespace](./docs/reference.md#default-namespace-100)
- [x] [Expose control plane](./docs/reference.md#exposed-control-plane-101)

## Usage

```bash
cd cmd/cli && go run .
```

### work in progress

**next:**  
&nbsp;&nbsp;&nbsp; -> improve how warnings are displayed  
&nbsp;&nbsp;&nbsp; -> add basic terminal UI  
&nbsp;&nbsp;&nbsp; -> extend the amount of validations  

&nbsp;

[![codecov](https://codecov.io/gh/converge/kind-security-check/branch/main/graph/badge.svg?token=k9yhzXWY9S)](https://codecov.io/gh/converge/kind-security-check)
