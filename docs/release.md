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

1. Deploy storage stack
    ```shell script
    make deploy-storage ENV=...
    ```
1. Deploy stack first (to get Bot URL)
    ```shell script
    make deploy-with-ready-secret ENV=... VERSION=...
    ```
1. Publish the bot following the instructions on [publishing bots](https://developers.google.com/hangouts/chat/how-tos/bots-publish?authuser=1)
    and using BOT URL from above.

    Icon URL can be taken from https://image.flaticon.com/icons/svg/248/248654.svg
1. Store credentials (from the step above) in the
GChatSecretKey via AWS Console (created in the deploy-storage step)
1. Deploy the application
```shell script
make deploy ENV=... VERSION=...
```