# edwp-03
#

# 1) Create a base ubuntu 16.04 layer
#    Hint : use FROM
FROM ubuntu:16.04

# 2) Create intermediate layer by installing git
#    Hint :apt-get update  
RUN apt-get update

# 3) Create inermediate layer by installing one  

# python 3.6
#    Hint: apt-get update 
#          apt-get install -y python3.6  

RUN apt-get install -y python3
RUN apt-get update       
RUN python3 -V 
RUN apt-get install -y python3-pip

RUN pip3 install numpy
RUN apt-get install -y build-essential libssl-dev libffi-dev python-dev
RUN apt-get install -y python3-venv
RUN mkdir environments
RUN cd environments
RUN python3 -m venv my_env
RUN ls my_env
RUN . my_env/bin/activate

# 3) Build an image named "dwp03" 
#    Hint : use "docker <build> .. command
#


# 4) Create the container and run in daemon mode
#    Hint : use "docker run -d <image name> "
#

# 5) Inspect the container check the versions for one of 
#	the java , python, node
#    Hint : use "docker exec -it <cnt> /bin/bash"
#

