## Docker workshop
This repository consists for 5 part excercies for undersanding docker basics
 * Part 1  - Introduces building an Image
 * Part 2 - Running your first container
 * Part 3 - Inspecting the Container
 * Part 4 - Network and Volumes
 * Part 5 - Containerize a RESTfull api (written in Sprinboot, Django, Express)

Each part Contains a Dockerfile and a solution folder, the docker file contains a series of Problems(steps) with hints.Solutions are provided in the solution folder. 
## Requirements
* docker-ce > 17.XX
## Docker installation
####  Linux
```
curl -fsSL get.docker.com -o get-docker.sh
sudo sh get-docker.sh

# Add current user to docker group
sudo usermod -aG docker $USER

#Logout and Login or Reboot
```
#### MAC OSX
```
https://download.docker.com/mac/stable/Docker.dmg

# Install the dmg
```
