# jerseyctf-registration-site
The website for the JerseyCTF Site

## Technical Details
* A static site 
* The back-end (web server) is written in **golang-go**
* The front-end is written using **Bootstrapv5**

## Purpose
* Made to be deployed onto a Droplet (cloud server on DigitalOcean) easily
* Check '[How To Deploy a Go Web Application Using Nginx on Ubuntu 18.04](https://www.digitalocean.com/community/tutorials/how-to-deploy-a-go-web-application-using-nginx-on-ubuntu-18-04)' for further instructions on deploying the site to web.

## Notes
* **[setup.sh](setup.sh)** 
    * Installs :
        *  _nginx_
        * _certbot_ 
        * _golang-go_
        * _python3-certbot-nginx_
    * Firewall :
        * _Enables Firewall if not active_
        * _Deletes a potential redundant Nginx 'HTTP' profile_
        * _Allows the Firewall to open for Nginx 'Full' Profile_
        * _Allows the Firewall to open for 'OpenSSH' connections_

