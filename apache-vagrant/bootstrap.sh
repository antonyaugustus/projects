#!/usr/bin/env bash

# ===================================================================
# Functions
# ===================================================================
start_if_stopped () {
  if [ ! -z "$1" ]
  then
    service=$1
    value=$(systemctl status ${service} | grep -c Active)
    if [ $value -eq 0 ]
    then
      systemctl start ${service}
    fi
  else
    printf "[ERROR] no parameter passed to start_if_stopped.\n"
  fi
}

# ===================================================================
# Pretify logging to screen
# ===================================================================
printLog() {
  printf "[${VAGRANT_HOST}-bootstrap] $1\n";
}

installUpdate() {
  yum update
  nmcli connection reload
  systemctl restart network.service
}

installApache() {
  FILE=/usr/sbin/httpd
  if [ ! -f $FILE ]; then
    printLog "Installing httpd";
    yum --quiet -y install httpd
    if ! [ -L /var/www ]; then
      rm -rf /var/www
      ln -fs /vagrant /var/www
    fi
    systemctl start httpd
    printLog "Altering httpd config";
    FILE=/etc/httpd/conf/httpd.conf
    perl -pi -e 's/#ServerName www.example.com:80/ServerName localhost.localdomain:80/g' $FILE
    start_if_stopped httpd
  fi
}

# ===================================================================
# Install software
# ===================================================================
installUpdate;
installApache;
