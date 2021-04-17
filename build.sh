#!/bin/bash
set -x

rm -f athena*
rm -f /Users/feiwang/prom-data/athena*
rm -f /Users/feiwang/prom-data/rules.yml
go build -o bin/athena
