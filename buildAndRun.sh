#!/bin/bash

#
# BUILD
#
echo ""
echo -e "\033[1;32m"
echo "***************************** BUILD RESULT *****************************"
echo -e "\033[0m"

go build -o bin/fforce src/*.go


#
# CLEAN FILES CREATED BY APP
#
echo ""
echo -e "\033[1;33m"
echo "***************************** CLEAN FILES ******************************"
echo -e "\033[0m"

rm -rf bin/src2
rm -rf bin/executeAnonymous
rm -rf bin/query
rm bin/build.properties
rm bin/build.xml
rm bin/openUrl
rm bin/login


#
# RUN APP
#
echo ""
echo -e "\033[1;36m"
echo "***************************** RUNNING APP ******************************"
echo -e "\033[0m"

cd ./bin
./fforce init -h -hello
# ./fforce