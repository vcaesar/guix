FROM golang:1.9.4

RUN apt-get update  
# customize dependencies
RUN apt-get install -y -qq libxi-dev libxcursor-dev 
RUN apt-get install -y libxrandr-dev libxinerama-dev 
RUN apt-get install -y mesa-common-dev libgl1-mesa-dev libxxf86vm-dev
RUN go get -u github.com/vcaesar/guix