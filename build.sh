#!/bin/sh

# Change current path to project directory path
cd ~/go-cgi-bin-test

# build path1
cd ./cgi-bin-path1
go build -o path1
cd ../
echo "path1 build success."

# build path2
cd ./cgi-bin-path2
go build -o path2
cd ../
echo "path2 build success."

# build path3
cd ./cgi-bin-path3
go build -o path3
cd ../
echo "path3 build success."

# build path4
cd ./cgi-bin-path4
go build -o path4
cd ../
echo "path4 build success."

