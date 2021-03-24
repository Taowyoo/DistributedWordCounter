#!/usr/bin/env bash

go build -o client src/client/main.go

go build -o worker src/worker/main.go