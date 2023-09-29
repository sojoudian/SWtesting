#Create multi stage Docker
#BUILD STAGE
FROM golang:1.21.1-alpine as builder

#Set the working directory
WORKDIR /app

COPY . .

#Build the app
RUN go build -o myapp

#Use a smaller image to reduce the size of the final container
FROM alpine:latest

#Set Working Directory
WORKDIR /root/

#Copy the binary from the builder stage
COPY --from=builder /app/myapp .

#Execute the app
CMD ["./myapp"]