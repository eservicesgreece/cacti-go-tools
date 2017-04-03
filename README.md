cacti-go-tools
==============
Tools for statistic acquisition from services for cacti


Install
-------

Configuration
-------------
All configuration is made in cacti-go-tools.json. The file needs to be copied in /etc/cacti-go-tools/ in order for the tool to function properly from any path.

SNMP Configuration
------------------
In your snmpd.conf you will need to add the following:
```bash
extend nginx /root/cacti-go-tools engine nginx
```

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