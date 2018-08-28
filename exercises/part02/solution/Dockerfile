# 1) set base image layer from ubuntu 18.04
FROM ubuntu:18.04

# 2) COPY "hello_world" binary to /home/ directory
COPY ./hello_world /home/

# 3) Change workdir to /home/
WORKDIR /home/

# 4) set entrypoint for running the process
ENTRYPOINT ["./hello_world"]

# 5) Build the image with the name "dwp02"
#     "docker build -t dwp02 ." 

# 6) Create a container using  dwp02 image
#    "docker run dwp02"

# 7) Run docker in daemon mode
#   "docker run -d dwp02"

# 8) Check the status of the container
#   "docker ps"
