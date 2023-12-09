#!/usr/bin/env bash
./abigen --abi="./abis/PayDB.abi" --type=PayDB --pkg=contracts --out=PayDB.go
./abigen --abi="./abis/VWManager.abi" --type=VWManager --pkg=contracts --out=VWManager.go