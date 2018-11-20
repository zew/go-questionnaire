# Application setup and administration

# Setup from scratch

## Linux setup instructions

### Install minimum tools

    apt-get install vim
    apt-get install tofrodos   # turn the newlines of shell scripts into proper format

## Create the Linux application user

    sudo adduser gquser
    sudo passwd  gquser    

    # as user gquser
    echo 'PATH="$PATH:/etc/init.d"' >> ~/.profile



## Prepare app dir

    sudo mkdir               /opt/go-questionnaire
    sudo chmod 755           /opt/go-questionnaire
    sudo chown gquser:gquser /opt/go-questionnaire

## Now copy the app files

The below sections `Golang installation` and `Download and compile this application`
can be executed either on your work machine, or on the server.

Whichever you chose, copy the application files from `~/go/src/github.com/zew/go-questionnaire`
to your application directory:

    cp go-questionnaire /opt/go-questionnaire
    cp static           /opt/go-questionnaire
    cp templates        /opt/go-questionnaire
    cp config.json      /opt/go-questionnaire
    cp logins.json      /opt/go-questionnaire
    cp server.key       /opt/go-questionnaire
    cp server.pem       /opt/go-questionnaire
    mkdir /opt/go-questionnaire/responses

    sudo chown -R gquser:gquser /opt/go-questionnaire/*
    sudo chmod -R 644   /opt/go-questionnaire/*
    sudo chmod -R 755   /opt/go-questionnaire/go-questionnaire # make it executable
    sudo chmod -R 755   /opt/go-questionnaire/static
    sudo chmod -R 755   /opt/go-questionnaire/templates
    sudo chmod -R 755   /opt/go-questionnaire/responses

### Enable ports 80 and 443 for non-root

 Needs redo after *each* compilation.

     sudo setcap cap_net_bind_service=+eip /opt/go-questionnaire/go-questionnaire
     # 'e', 'i', and 'p' flags specify the (e)ffective, (i)nheritable and (p)ermitted sets.


## Prepare log file 

     sudo touch /var/log/go-questionnaire.log
     sudo chown gquser:gquser /var/log/go-questionnaire.log

     # reset log
     truncate   --size=0  /var/log/go-questionnaire.log

##  Prepare pid file 

     sudo mkdir /var/run/go-questionnaire
     sudo chown gquser:gquser /var/run/go-questionnaire/

     sudo touch       /var/run/go-questionnaire/prog.pid
     sudo chown gquser:gquser /var/run/go-questionnaire/prog.pid
     sudo rm          /var/run/go-questionnaire/prog.pid


## Make script reboot-hard - and accessible to all

Put the script ```go-questionnairectl``` to /etc/init.d

```go-questionnairectl``` is a start-stop script.
The source is in the same directory as this file.

    sudo mv ./go-questionnairectl  /etc/init.d/go-questionnairectl
    sudo chmod 755 /etc/init.d/go-questionnairectl
    fromdos /etc/init.d/go-questionnairectl   # remove windows newlines



###  chkconfig

Under debian, we do not need ```chkconfig``` - just put the script to init.d

      chkconfig: 2345 85 15
      description: 
              2,3,4,5      runlevel
              85           starting.
              15           stopping.
      [root@host ~]# chkconfig --add go-questionnairectl
      [root@host ~]# chkconfig --list | grep -i bspc





## Manage app with service commands

     go-questionnairectl status
     go-questionnairectl stop
     go-questionnairectl start


## Manage app manually

     cd /opt/go-questionnaire/
     ./go-questionnaire > /var/log/go-questionnaire.log 2>&1 &
     tail /var/log/go-questionnaire.log
     ps aux | grep go-questionnaire
     pkill go-questionnaire


## Denial of service consideration

* The hosting machine should be behind a firewall  
preventing denial-of-service attacks on network level.


## Golang installation

* Install and setup [golang](https://golang.org/doc/install)

* By default your `golang installation` will end up here:

      /usr/local/go  # under Linux
      c:\Go          # under Windows
      # Otherwise set $GOROOT to your different path

* By default your `source files` are assumed to be here:

      %USERPROFILE%\go
      ~\Go
      # Otherwise set $GOPATH to your source file directory

* Your compiled go programs end up here

      %USERPROFILE%\go\bin
      ~\go\bin
      # To have them always available:
      export PATH=$PATH:~/go/bin


* For details, refer the [golang install docs](https://golang.org/doc/install)



## Download and compile this application

* Source code is hosted at https://github.com/zew/go-questionnaire.

* Thus the source code should go to

      # mkdir ...
      cd ~/go/src/github.com/zew
      cd ..
      git clone https://github.com/zew/go-questionnaire
      cd go-questionnaire

* Fetch all required libraries with  

      `go get ./...`  

* Compile the application

      go build

#### Compiling from another machine

* You may cross compile under windows for linux using `crosscomp.bat`.

* Copy the new executable to yourhost.com using sftp.  
The new copy should retain execution privileges.  

#### Minimal directories

* Copy directories `static/...` and `templates/...` 

* If you want to run https,  
then put your `server.key` and `server.pem` files into the app dir.  

#### Create a configuration file

* Whenever you start the application,  
a file `config-example.json` is created.

* Derive your settings and save it as `config.json`  
into the app dir.


* Each new executable needs to be configured *again*  
to allow to use ports 80 and 443.  
See section **Enable ports 80 and 443** .

#### Adding an application user

* Whenever you start the application,  
a file `logins-example.json` is created.

* A new user must be entered into logins.json.  

    * `user` => Username, all lowercase, a-z, 0-9

    * `pass_initial` => Set an one-time password

    * `is_init_password` => Set to true

    * `email` => Optional email. So far unused.

 	* Optional `roles`

    ```
    "roles": {  
        "admin": "yes"  
    },
    ``` 


    * Login as admin call  `/logins-reload`

    * Or restart application

    * You may call `/logins-save`   
    to fill empty `pass_initial` fields with autogenerated passwords.






### Getting a trusted SSL certificate

Consider the golang [acme stuff](https://github.com/letsencrypt/boulder) for integration with letsEncrypt. Acme is even making provisions for automatic cert renewal.

[Some Golang SSL Info](https://gist.github.com/denji/12b3a568f092ab951456)

Activation of https via config setting `"tls": true,`

Generate private key for algorithm "RSA" ≥ 2048-bit

```
openssl genrsa -out server.key 2048
```

Key considerations for algorithm "ECDSA" ≥ secp384r1

List ECDSA supported curves:  `openssl ecparam -list_curves`

```
openssl ecparam -genkey -name secp384r1 -out server.key
```

Generation of self-signed (x509) public key based on the private key. PEM-encodings .pem|.crt

```
openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650
```

pem is a Privacy Enhanced Mail Certificate file

###	Checking the modulus

    openssl x509 -noout -modulus -in server.pem
    openssl rsa -check -noout -modulus -in server.key


### Use apache to run multiple instances under port 80

#### We must allow apache to use the network in order to proxy requests:

```
    sestatus -b | grep httpd_can
    setsebool -P httpd_can_network_connect=1
```

Put the app behind an apache virtual host.

Edit httpd.conf: 


```

# cache nothing ever
# serverfault.com/questions/4729/
<Location / >
   ExpiresActive On
   ExpiresDefault "now"
</Location>

# default virtual VirtualHost
<VirtualHost *:80>
    DocumentRoot "C:/xampp/htdocs"
</VirtualHost>

# enable mod_proxy_html.so
<VirtualHost *:80>
    ServerName go-questionnaire.myorg.net
    ProxyPreserveHost On
    ProxyPass        "/"   "http://127.0.0.1:8080/"
    ProxyPassReverse "/"   "http://127.0.0.1:8080/"    
</VirtualHost>

```

another example with multiple virtual hosts  
_and_ multiple instances of go-questionnaire

```
# cache nothing ever
# serverfault.com/questions/4729/
<Location / >
   ExpiresActive On
   ExpiresDefault "now"
</Location>

<VirtualHost *:80>
    ServerName some-other.myorg.net
    DocumentRoot "/var/www/some-other.myorg.net"
    <Directory /var/www/some-other.myorg.net/>
        Options Indexes FollowSymLinks MultiViews
        AllowOverride None
        Order allow,deny
        allow from all
    </Directory>
</VirtualHost>


# enable mod_proxy_html.so
<VirtualHost *:80>
    ServerName go-questionnaire.myorg.net
    # doc root is ignored
    DocumentRoot "/var/www/go-questionnaire.myorg.net"
    <Directory /var/www/exceldb.myorg.net/>
        Options Indexes FollowSymLinks MultiViews
        AllowOverride None
        Order allow,deny
        allow from all
    </Directory>
    ProxyPreserveHost On
    ProxyPass        "/app1"   "http://127.0.0.1:8080/app1"
    ProxyPass        "/app2"   "http://127.0.0.1:8081/app2"
    ProxyPassReverse "/app1"   "http://127.0.0.1:8080/app1"
    ProxyPassReverse "/app2"   "http://127.0.0.1:8081/app2"    
</VirtualHost>

```


[Back to documentation root](/{{AppPrefix}}/doc)

[Back to app](/{{AppPrefix}})