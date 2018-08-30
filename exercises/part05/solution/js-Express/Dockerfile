# Edwp-05 : 
#----------------------------------------------

# 1) Get the base image 
#    Hint : use FROM Keyword
#       For Java : java:8
#       For Python : 
#       For Node : node:8.11.4 

FROM node:8.11.4

# 2) Create a folder called app under /home
#     Hint : user mkdir keyword
RUN mkdir /home/app

# 3) Copy the code from source folder to /home/app/
#    Hint : use COPY keyword

COPY js-Express/ /home/app/
# 4) Change the working directory to /home/app/
#    Hint : use WORKDIR cmd

WORKDIR /home/app/

# 5) Install the dependencies
#    Hint : use RUN keyword
#         For Java:
#         For Python:    
#         For Node: npm install

RUN npm install
# 6) Expose the server port (this is for dev reference, EXPOSE cmd will not
#      open up any ports)  
#    Hint : use EXPOSE keyword
EXPOSE 9000

# 7) Start the server with Entrypoint
#    Hint : use ENTRYPOINT keyword
#       For Java :
#       For Python :
#       For Node : node server.js
ENTRYPOINT ["node", "server.js"]
# 8) Create the image  
#   Hint : use docker build

# 9) Run the container and expose the port
#   Hint : use docker run 


