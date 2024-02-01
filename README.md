# envset GitHub Action

This action sets an environment variable in GitHub Actions workflow.
It is useful when you want to set an environment variable that may contain double or single quotes, and you don't know the value in advance.

Action writes into a `GITHUB_ENV` file, so the environment variable will be available in the next steps of the same job.

This is a Docker-based action, so it may not work on Windows and MacOS runners.

## Inputs

| Name  | Description                | Required |
|-------|----------------------------|----------|
| name  | Environment variable name  | true     |
| value | Environment variable value | true     |

## Usage

```yaml
name: dispatch

on:
  workflow_dispatch:
    inputs:
      some_value:
        description: Some value
        required: true

dispatch:
  pull:
    runs-on: ubuntu-latest
    steps:
      - name: Set environment variable
        uses: alsosee/envset@v1
        with:
          name: SOME_VARIABLE
          value: ${{ github.event.inputs.some_value }}

      ...
```
