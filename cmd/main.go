package main

import (
	"github.com/jinzhu/gorm"
	"github.com/wr2net/imersao-fullcylce-codepix/application/grpc"
	"github.com/wr2net/imersao-fullcylce-codepix/infrastructure/db"
	"os"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 500051)
}
