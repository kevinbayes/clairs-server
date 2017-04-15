#!/usr/bin/env bash

curl -X PUT \
     -H "Content-Type: application/json" \
     -d @sample-new-registry.json \
     -i \
     http://localhost:18080/api/registries/4