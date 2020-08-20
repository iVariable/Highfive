# How to release

Application consists of two stacks:
- Storage (DDB table)
- Application itself

Before running any deployment command make sure that the project
was initialized via `make init`

## Usual deployment flow

Available envs: dev | prod

```shell script
make deploy ENV=... VERSION=...
```

## Storage deployment

```shell script
make deploy-storage ENV=...
```

## First time deployment

```shell script
make deploy-storage ENV=...
make deploy ENV=... VERSION=...
```