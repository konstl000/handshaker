# Handshaker

This is a simple go server that listens on a port provided as an env variable and returns "OK". It can be used to pass health checks for load balancers which can be sometimes quite useful in multicontainer environments.
Can be also used for test purposes.

### Build
Use `docker build . -t some:tag` to build it in a docker container. 
You can also use the `deploy.sh` script to build and run the container locally (to check what it does). 
