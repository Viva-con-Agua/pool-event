package dao

import "github.com/Viva-con-Agua/vcago"

var Logger = vcago.NewLoggingHandler("pool-event")

var Database = vcago.NewMongoDB("pool-event")
