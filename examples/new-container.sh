#!/usr/bin/env bash

curl -X POST \
     -H "Content-Type: application/json" \
     -d @sample-new-container.json \
     -i \
     http://localhost:18080/api/registries/4/containers