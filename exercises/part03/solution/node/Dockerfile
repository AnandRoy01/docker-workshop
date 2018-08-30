# edwp-03
#

# 1) Create a base ubuntu 16.04 layer
#    Hint: use FROM
FROM ubuntu:16.04

# 2) Create intermediate layer by installing git
#    Hint: apt-get update  
#          apt-get install -y git

RUN apt-get update -y
RUN apt-get install -y git 

# 3) Create inermediate layer by installing one  
#      c) node 8.11.4 (Fix this)
	# curl -sL https://deb.nodesource.com/setup_8.x | bash -
        # apt-get install -y nodejs
        # apt-get install -y npm 

RUN apt-get install -y curl
RUN curl -sL https://deb.nodesource.com/setup_8.x | bash -
RUN apt-get update       
RUN apt-get install -y nodejs
RUN node --version

# 3) Build an image named "edwp03" 
#    Hint : use "docker <build> .. command
#

# 4) Create the container and run in daemon mode
#    Hint : use "docker run -d <image name> "
#

# 5) Inspect the container check the versions for node
#    Hint : use "docker exec -it <cnt> /bin/bash"
#

