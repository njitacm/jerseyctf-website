#!/bin/bash

# Meant for Linux, and specfically built for Ubuntu Distro 

# check() -> checks for command, if none existent it installs via apt
check () {
    # the `! command -v ()` checks if the commands / program exists
    if ! command -v $1 &> /dev/null
    then
        echo "Installing $1..."
        sudo apt install $1 -y
        exit
    fi
    echo "$1 is already installed"
}

## INSTALLATION of packages
# Checks and downloads packages
check nginx
check certbot
check python3-certbot-nginx
check golang-go

## FIREWALL
# Checks if ufw is enabled or disabled, if disabled, it turns the firewall on
if [[ "$(sudo ufw status | awk '{print $2; exit}')" == "inactive" ]]; then
    sudo ufw enable
fi

# Delete Redundant Nginx HTTP profile, if it exists
sudo ufw delete allow 'Nginx HTTP'

# Allow Nginx Full Profile
sudo ufw allow 'Nginx Full'

# Opens the Firewall to OpenSSH connections 
sudo ufw allow 'OpenSSH'
# Should be commented out if you are running a more advanced configuration
# This was just put in case, someone forgets to allow for their SSH connections
# I know this is bad in terms of programming, because the program
# does more than it is subscribed to do but lol 

