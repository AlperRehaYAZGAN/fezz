### Fezz (!under maintainance!)
Fezz is a project that aiming to bring web3 contract (Ethereum-EVM-, Solana-solana programs-) philosophy to web2. Fezz is a HTTP JSON RPC web service application that supports deploy web2 services like web3 contracts (via golang shared objects -".so extension"-).

#### Motivation  
Long story short, Fezz is a pluginable HTTP JSON RPC server for only serving your services via executing your only deployed byte code (lightweight, only ~1MB per service) and managing their lifecycle, versioning, and upgradability (like Docker lifecycle management but lightweight one -really-).  

Web3 contracts are immutable, upgradable, and can be deployed by anyone. Fezz aims to bring these features to web2 services. In web3 contract development, developers write their contract code in a high-level language like Solidity, Rust, etc. and compile it to EVM or Solana bytecode. Then they deploy it to the blockchain. Fezz aims to bring this philosophy to web2 services. Developers can write their service code in any language (at this point golang) that can be compiled to a shared object file (".so extension") and deploy it to running json rpc Fezz while runtime. Fezz will manage the service lifecycle, versioning, and upgradability. This will ensure that no downtime and no service interruption will occur during the upgrade process. And its lightweighest solution considering the docker daemon. Cause it's not a container runtime, it's a only bytecode runtime.

#### Guide  
Repository is under development and it is in PoC stage. You could test it by following the steps below.  

1. Build sample extension in file [hello_world.go](./plugins/hello_world.go)
```bash
# first extension type is fully stateless app.
go build \
    -buildmode=plugin \
    -ldflags="-s -w" \
    -o plugins/hello_world.so plugins/hello_world.go

# second extension is sqlite select row by find query application.
go build \
    -buildmode=plugin \
    -ldflags="-s -w" \
    -o plugins/get_user_by_name.so plugins/get_user_by_name.go

```  

2. Run application
```bash
go run main.go
```

3. Test application via [VSCode REST Client extension](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) or [Postman](https://www.postman.com/) in files [get_user_by_name.http](./test/get_user_by_name.http) and [hello_world.http](./test/hello_world.http)
```http
POST http://localhost:7071/rpc
Content-Type: application/json

{
  "jsonrpc": "2.0",
  "method": "hello_world.so",
  "params": [],
  "id": 1
}
```

```http
POST http://localhost:7071/rpc
Content-Type: application/json

{
  "jsonrpc": "2.0",
  "method": "get_user_by_name.so",
  "params": [
    {
        "name": "john"
    }
  ],
  "id": 1
}
```
