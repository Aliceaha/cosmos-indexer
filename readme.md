# Cosmos indexer

WIP indexer for Cosmos hub, using [SurrealDB](https://surrealdb.com/ "SurrealDB")

### realtime indexing

`go run indexer.go realtime cosmos_mainnet`

### basic ingester

`go run indexer.go ingester cosmos_mainnet`

### basic ingester + realtime

`go run indexer.go full cosmos_mainnet`



## SurrealDB

Install [tiup](https://github.com/pingcap/tiup)

`curl -sSf https://tiup-mirrors.pingcap.com/install.sh | sh`

run tikv instance

`tiup playground --tag surrealdb --mode tikv-slim --pd 1 --kv 1`

then, run surrealDB

`surreal start tikv://127.0.0.1:2379 --bind 0.0.0.0:8000 --user root --pass 123456`
