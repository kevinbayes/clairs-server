#!/usr/bin/env bash

curl -X PUT \
     -H "Content-Type: application/json" \
     -d @sample-new-container.json \
     -i \
     "http://localhost:18080/api/containers/$1/_evaluate"