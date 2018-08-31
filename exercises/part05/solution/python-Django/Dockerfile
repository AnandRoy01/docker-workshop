#
# Edwp-05 : 
#----------------------------------------------

# 1) Get the base image 
#    Hint : use FROM Keyword
#       For Java : java:8 
#       For Python : python:3.6 
#       For Node : node:8.11.4 
FROM python:3.6

# 2) Create a folder called app under /home
#     /home/app
#     Hint : user mkdir keyword

RUN mkdir /home/app
# 3) Copy the code from source code folder to /home/app/
#    Hint : use COPY keyword
#    
#    For Java : copy java-spring-boot
#    For Python: copy python-Django 
#    For JS : copy js-Express
#
COPY python-Django/ /home/app/
# 4) Change the working directory to /home/app/
#    Hint : use WORKDIR cmd

WORKDIR /home/app/
# 5) Install the dependencies
#    Hint : use RUN keyword
#         For Java: ./gradlew clean build
#         For Python: pip3 install -r requirements.txt   
#         For Node: npm install

RUN pip3 install -r requirements.txt
# 6) Expose the server port (this is for dev reference, EXPOSE cmd will not
#      open up any ports)  
#    Hint : use EXPOSE keyword

EXPOSE 9000
# 7) Start the server with Entrypoint
#    Hint : use ENTRYPOINT keyword
#       For Java : java -jar ./build/libs/dockerworkshop-0.0.1.jar
#       For Python : python3.6 manage.py runserver 0.0.0.0:9000
#       For Node : node server.js
ENTRYPOINT ["python", "manage.py", "runserver", "0.0.0.0:9000"] 

# 8) Create the image  
#   Hint : use docker build

# 9) Run the container and expose the port
#   Hint : use docker run 


