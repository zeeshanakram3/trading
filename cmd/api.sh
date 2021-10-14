#!/usr/bin/env bash
echo "Exporting api keys"

# Mongo atlas keys
export MONGO_ATLAS_DB_NAME="trading_db"
export MONGO_ATLAS_DB_PASSWORD="mongo123"

# api keys
export MESSARI_API_KEY="74864659-e0cc-4542-af0e-c0407a270f9c"
export CMC_API_KEY="99771ff5-35d9-459d-8ace-af8cf15f5c1c"

# base urls
export CMC_BASE_BASE_URL="https://pro-api.coinmarketcap.com"
export COINGEKO_BASE_URL="https://api.coingecko.com/api/v3"
export MESSARI_BASE_URL="https://data.messari.io/api"
