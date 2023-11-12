#base go image
FROM golang:1.18-alpine as builder

#Create a folder called app inside the docker container
RUN mkdir /app

#Copy everything from the current folder(broker-service) into
#the container's app folder
COPY . /app

#Make the app folder in the container the working directory.
#Every command will run agains the app folder by default.
WORKDIR /app

#Build the service and name the executable brokerApp.
#CGO_ENABLED=0 means there are no C libraries, just the standard go libraries.
RUN CGO_ENABLED=0 go build -o brokerApp ./cmd/api

#Make sure the compiled file is executable
RUN chmod +x /app/brokerApp

#build a tiny docker image
FROM alpine:latest

#Create a folder called app inside the docker container
RUN mkdir /app

#From our base container(above) copy the executable into the new container's app folder
COPY --from=builder /app/brokerApp /app

#Copy the binary into the app folder
COPY brokerApp /app

CMD [ "/app/brokerApp" ]