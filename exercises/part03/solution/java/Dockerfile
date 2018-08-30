# Edwp-03
#

# 1) Create a base ubuntu 18.04 layer
#    Hint : use FROM
FROM ubuntu:16.04
# 2) Create intermediate layer by installing git
#    Hint :apt-get update  
#          apt-get install -y git

RUN apt-get update -y
#RUN apt-get install -y git 
# 3) Create inermediate layer by installing one  

#      a) java 1.8
        # apt-get update
        # apt install openjdk-8-jdk

#      b) python 3.6
        # apt-get update 
        # apt-get install -y python3.6  
#
#      c) node 8.11.4 (Fix this)
        # apt-get install -y nodejs
	# curl -sL https://deb.nodesource.com/setup_8.x | bash -
        # apt-get install -y npm 

RUN apt-get install -y curl
RUN curl -sL https://deb.nodesource.com/setup_8.x | bash -
RUN apt-get update       
RUN apt-get install -y nodejs
RUN node --version

# 3) Build an image named "dwp01" 
#    Hint : use "docker <build> .. command
#


# 4) Create the container and run in daemon mode
#    Hint : use "docker run -d <image name> "
#

# 5) Inspect the container check the versions for one of 
#	the java , python, node
#    Hint : use "docker exec -it <cnt> /bin/bash"
#

