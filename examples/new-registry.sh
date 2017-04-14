#!/usr/bin/env bash

curl -X POST \
     -H "Content-Type: application/json" \
     -d @sample-new-registry.json \
     -i \
     http://localhost:18080/api/registries