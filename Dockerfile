FROM golang:1.22

#Set destination for COPY
WORKDIR /app

#Copy the source code
COPY ./ ./

# Doanload Go modules
RUN go mod download



#Build
RUN CGO_ENABLED=0 GOOS=linux go build -o docker-mongo-crud

EXPOSE 8080

RUN chmod +x docker-mongo-crud
#Run
CMD ["./docker-mongo-crud"]