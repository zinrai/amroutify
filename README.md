# amroutify

`amroutify` is a testing tool for [Alertmanager](https://prometheus.io/docs/alerting/latest/alertmanager/) routing configurations. It helps verify that alerts with specific labels are routed to the expected receivers.

## Overview

Maintaining Alertmanager routing configurations can become challenging as they grow more complex. Manual testing of routes is time-consuming and error-prone. amroutify helps by:

1. **Automating verification**: Test all routing paths with a single command
2. **Preventing regressions**: Catch unintended routing changes before deployment
3. **Documentation as code**: Test cases serve as documentation for expected routing behavior
4. **Supporting complex scenarios**: Test routes with `continue: true` and multiple receivers

## Features

- **Table-based testing**: Define multiple test cases in a YAML format
- **Support for complex routing**: Test scenarios with `continue: true` resulting in multiple receivers
- **Order-independent verification**: Receiver order doesn't matter for test success
- **CI/CD Integration**: Integrate with GitHub Actions or other CI systems
- **Detailed output**: Verbose mode shows matching routes and receivers for debugging

## Installation

```bash
$ go build
```

## Usage

Basic usage:

```bash
$ amroutify -config alertmanager.yml -tests routing_tests.yml
```

Detailed output:

```bash
$ amroutify -config alertmanager.yml -tests routing_tests.yml --verbose
```

## Example

The `example/` directory contains sample configuration files demonstrating complex routing scenarios:

- `alertmanager.yml` - Alertmanager configuration with nested routes, `continue` flags, and multiple receivers
- `routing_tests.yml` - Corresponding test cases covering all routing paths

Run the example:
```bash
$ amroutify -config example/alertmnager.yml -tests example/routing_tests.yml
```

## Exit Codes

- `0`: All tests passed
- `1`: One or more tests failed

## Test Case Format

The test cases are defined in YAML format:

```yaml
tests:
  - name: "Critical database alert"
    labels:
      service: database
      severity: critical
    expected_receivers:
      - "team-DB-pager"
    description: "Critical database alerts should go to DB team pager"

  - name: "Multi-receiver test with continue flag"
    labels:
      service: database
      owner: team-X
    expected_receivers:
      - "team-DB-pager"
      - "team-X-email"
    description: "Alert matches multiple receivers due to continue flag"
```

## License

This project is licensed under the [MIT License](./LICENSE).
