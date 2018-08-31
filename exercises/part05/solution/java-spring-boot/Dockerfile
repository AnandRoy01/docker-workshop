#
# Edwp-05 : 
#----------------------------------------------

# 1) Get the base image 
#    Hint : use FROM Keyword
#       For Java : 1.8 
#       For Python : 
#       For Node : node:8.11.4 
FROM java:8

# 3) Create a folder called app under /home
#     Hint : user mkdir keyword
RUN mkdir /home/app

# 4) Copy the code from source folder to /home/app/
#    Hint : use COPY keyword

COPY java-spring-boot/ /home/app/
# 5) Change the working directory to /home/app/
#    Hint : use WORKDIR cmd

WORKDIR /home/app/

# 6) Install the dependencies
#    Hint : use RUN keyword
#         For Java:
#         For Python:    
#         For Node: npm install

RUN ./gradlew clean build
# 7) Expose the server port (this is for dev reference, EXPOSE cmd will not
#      open up any ports)  
#    Hint : use EXPOSE keyword
EXPOSE 9000

# 8) Start the server with Entrypoint
#    Hint : use ENTRYPOINT keyword
#       For Java :
#       For Python :
#       For Node : node server.js
ENTRYPOINT ["java", "-jar", "./build/libs/dockerworkshop-0.0.1.jar"]
# 9) Create the image  
#   Hint : use docker build

# 10) Run the container and expose the port
#   Hint : use docker run 


