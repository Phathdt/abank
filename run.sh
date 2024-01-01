#!/usr/bin/env bash

echo "Migrate"
./abank migrate up

echo "Start server..."
./abank
