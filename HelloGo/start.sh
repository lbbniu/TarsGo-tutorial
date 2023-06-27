#!/bin/bash
set -ex
make
./HelloGo --config=config/config.conf
