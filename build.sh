#!/bin/bash

#
# BUILD
#
echo ""
echo -e "\033[1;32m"
echo "***************************** BUILD RESULT *****************************"
echo -e "\033[0m"

go build -o bin/fforce src/*.go
