package dao

import "github.com/Viva-con-Agua/vcago"

var Logger = vcago.NewLoggingHandler("pool-core")

var Database = vcago.NewMongoDB("pool-core")
