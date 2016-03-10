Fibonacci Server
====================

## Prerequesites
* A Linux OS that supports Docker
* Docker (tested on v1.10.2)
* Make


## Typical Make Commands
### make run
Use this command to start the server. It is currently configured to listen at 
port 8080. This command will also automatically build the docker image if 
it hasn't been previously built.

### make test
Use this command to run the go test framework.

### make remove-all
Use this command to remove all of the automatically generated 
docker images and containers from your system. Think of this as a 
complete uninstall. Also use this command when making changes to the makefile.


## Advanced Make Commands
### make stop
Use this command to stop the server that was started with `make run`.

### make force-build
Use this command when you have modified the Dockerfile.
This will attempt to rebuild the docker image even if the image already exists. 


## Configurations
### Published Server Port
The published server port is configured from the FIB_PORT variable in the
makefile. This can be easily found at the top of the makefile. 
Run `make remove-all` after changing.
