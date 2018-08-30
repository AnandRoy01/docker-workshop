# Edwp-04
#

#Part - A

# 1) Create a base ubuntu 16.04 layer
#    Hint : use FROM
FROM ubuntu:16.04

# 2) Create a directory called app at home
#    Hint : mkdir -p  /home/app
RUN mkdir -p /home/app

# 3) Copy hello_server to /home/app         
#    Hint : use COPY command
COPY ./hello_server /home/app/

# 4) Change working dir to /home/app/
#   Hint use WORKDIR command
WORKDIR /home/app/

# 4) Expose the port 9000 (this will not acutally open/unblock the port)
#    Hint : use EXPOSE command
EXPOSE 9000

# 5) Specify the entry point to start the server
#    Hint : use ENTRYPOINT
ENTRYPOINT ["./hello_server"]

# 6) Build the image dwp-04
#    Hint : docker build -t <name> .

# 7) Run two containers
#    Hint : docker run -d <image> 

# 8) Inspect the network, make sure both the container
#    are in the same bridge
#    Hint : use docker network ls

#-------------------------------------------------------

# Part - B (Stop both the containers)
# Hint : use docker stop

# 9) Create a directory called "cm" at /home/app
#   Inside the container
# Hint : use  RUN with mkdir 
RUN mkdir /home/cm

# 9) Create a volume called "Mordor"
# Hint : Use docker volume create <name>

# 10) Inspect the volume Mordor
# Hint : docker volume inspect

# 11) Run the container and mount the volume mordor 
#  under /home/app/cm
# Hint :  docker run with --mount source=mordor, target=/home/app/cm

# 12) Inspect the container and create a dummy file under /home/app/cm
# Hint : use docker exec -it <image> /bin/bash

# 13) Check if the newly created file is persisted in  after stopping the container
# Hint : docker inspect
 


