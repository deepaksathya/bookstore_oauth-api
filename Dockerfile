FROM golang:1.12.13

#Configure the repo url so we can configure our work directory:
ENV REPO_URL=bookstore_oauth-api

# Setup our $GOPATH
ENV GOPATH=/app
ENV APP_PATH=$GOPATH/$REPO_URL

# /app/bookstore_oauth-api/src

ENV WORKPATH=$APP_PATH/src

# Copy the entire source code from the current directory to $WORKPATH
COPY src/ $WORKPATH
WORKDIR $WORKPATH
RUN pwd
RUN ls
RUN go build -o oauth-api .

#Expose port 8081 to the world:
EXPOSE 8081
CMD [“./oauth-api”]
