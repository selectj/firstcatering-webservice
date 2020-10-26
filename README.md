# First Catering RESTful API
L4 Software Developer Synoptic Project by **Joseph Metcalfe**

## How to Install & Run the Application
To install & run my application you will need Go installed. Please see the following on how to install Go for your system: https://golang.org/doc/
Please note these instructions are for Linux and may differ ever so slightly to a Windows installation. 
After installation as mentioned in the Go documentation, please ensure the environment varaible $GOPATH is set.

You will then need to get the following Go packages, by running the following:  
`go get github.com/gorilla/mux`  
`go get github.com/go-sql-driver/mysql`  


Place all the `.go` files inside a directory, navigate to that directory and run `go build` to compile the application into an executable program named `api`.

### Example
```bash
mkdir /var/firstcatering-webservice
go build /var/firstcatering-webservice
```

(Optional) Now to create the systemd service to allow the application to run as a daemon:
```bash
nano /lib/systemd/system/fcws.service
```

I used the following for my systemd service:
```
[Unit]
Description=First Catering RESTful API

[Service]
ExecStart=/var/firstcatering-webservice/api
Restart=on-failure
RestartSec=5s
Type=simple

[Install]
WantedBy=multi-user.target
```

Then reload systemd and enable & start the service:
```bash
systemctl daemon-reload
systemctl enable fcws.service
systemctl start fcws.service
```

Should you wish to run the application without creating it as a systemd service, you could just run it from the command line in a screen session or something similar:
```bash
screen
/var/firstcatering-webservice/api
Ctrl+A then D #Detach from screen session
screen -r #To resume screen session
```