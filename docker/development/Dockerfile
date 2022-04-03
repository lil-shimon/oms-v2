FROM golang:latest

COPY server/ /go/src/server
WORKDIR /go/src/server

# gin install
RUN go get -u github.com/gin-gonic/gin && \ 
  # gorm install
  go get github.com/jinzhu/gorm && \
  go get github.com/go-sql-driver/mysql

EXPOSE 1323

CMD ["go", "run", "main.go"]