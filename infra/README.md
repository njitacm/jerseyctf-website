# Infrastructure Configurations

## File, Description and Location 
| File | Description | Location (Subject to change)
| :--: | :---------: | :-------
| nginx.conf | Default Nginx configuration | `/etc/nginx/sites-available/jerseyctf.com`
| goweb.conf | Goweb Service Configuration | `/lib/systemd/system/goweb.service`


## Other Technical Details
- Ensure that `/etc/nginx/sites-available/<domain>` has a symbolic link to `/etc/nginx/sites-enabled/<domain>`


## Reference 
- Check '[How To Deploy a Go Web Application Using Nginx on Ubuntu 18.04](https://www.digitalocean.com/community/tutorials/how-to-deploy-a-go-web-application-using-nginx-on-ubuntu-18-04)' to ensure steps are done properly with the golang-go program.