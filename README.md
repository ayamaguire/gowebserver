# gowebserver
Little webserver in Go by Aya Maguire.
Usage: Put webserv.go in its own folder in your src folder. 
Run it with "go run webserv.go "

It will listen on port 8080. Send a post request like:
curl --data "password=MyPassword" http://localhost:8080

It will hash your password, wait five seconds, then respond telling you it has done so.

To shut down, send the following request:
curl --data "Shutdown" http://localhost:8080
