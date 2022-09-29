#!/bin/bash

# Update
sudo apt-get update
echo Y | sudo apt-get upgrade

# Install MySQL Server
echo Y | sudo apt install mysql-server

# Create user
sudo mysql --execute="grant all privileges on *.* to 'myuser2'@'%' identified by 'password';"
sudo mysql --execute="flush privileges;"

# Allow remote access
sed -i s/"bind-address            = 127.0.0.1"/"bind-address            = 0.0.0.0"/g /etc/mysql/mysql.conf.d/mysqld.cnf

# Restart MySQL
sudo systemctl restart mysql
