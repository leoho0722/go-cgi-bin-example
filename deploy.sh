#!/bin/sh

# Change current path to project directory path
cd ~/go-cgi-bin-test

# copy path1
cd ./cgi-bin-path1
sudo cp ./path1 /usr/lib/cgi-bin/path1
sudo cp ./path1 /var/www/html/path1
cd ../
echo "path1 copied success."

# copy path2
cd ./cgi-bin-path2
sudo cp ./path2 /usr/lib/cgi-bin/path2
sudo cp ./path2 /var/www/html/path2
cd ../
echo "path2 copied success."

# copy path3
cd ./cgi-bin-path3
sudo cp ./path3 /usr/lib/cgi-bin/path3
sudo cp ./path3 /var/www/html/path3
cd ../
echo "path3 copied success."

# copy path4
cd ./cgi-bin-path4
sudo cp ./path4 /usr/lib/cgi-bin/path4
sudo cp ./path4 /var/www/html/path4
cd ../
echo "path4 copied success."

# restart apache2
sudo systemctl restart apache2
echo "apache2 restart success."