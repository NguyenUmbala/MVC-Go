FROM golang:1.10

# Install gin-gonic/gin
RUN go get github.com/gin-gonic/gin && go get github.com/jinzhu/gorm && go get github.com/jinzhu/gorm/dialects/sqlite

# Expose the app on port 8080
EXPOSE 8080

//Set the entry point of the container
CMD ["go", "run"]
