#!/bin/sh
if [ "$DEBUG" = "true" ]; then
    echo "Starting in DEBUG mode (Delve)"
    exec /dlv exec /aletheia-api --headless --listen=:40000 --api-version=2 --accept-multiclient --continue
else
    echo "Starting in PRODUCTION mode"
    exec /aletheia-api
fi