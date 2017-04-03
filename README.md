cacti-go-tools
==============

# cacti-go-tools
Tools for statistic acquisition from services for cacti


Install
-------

nginx
-----

You will need to enable the /status end point in nginx, you can include the configuration in conf/nginx/nginx_status.conf

```bash
Server {
...
    include /etc/nginx/nginx_status.conf
...
}
```