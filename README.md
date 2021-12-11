# First Catering RESTful API
L4 Software Developer Synoptic Project by **Joseph Metcalfe**

## Project Requirements
Your task is to create a RESTful web service for First Catering Ltd that will allow
Bows Formula One employees to use their existing data cards in the kiosks to
register and top up with money.

- If the card is not registered on the system, the card owner will be required to provide
basic employee information:
    - unique employee ID;
    - name;
    - email,
    - mobile number.
- A four-digit pin number chosen by the employee should be used for further security.
- The data used for providing the services must be held in a database.
- When a card is presented to the system and the service finds that the card is already
registered, the system will show a welcome message with the user’s name
associated with the card.
- If the user’s card is not registered, then the system will respond requesting that the
card needs to be registered.
- It is envisaged that when the user taps their card a second time the system informs
the user and says “Goodbye”.
- The application should timeout after a number of minutes of inactivity.
- You only need to provide the REST API, which should conform to industry standards.

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
