package config

const username = "root"
const password = ""
const host = "localhost"
const port = "3306"
const dbName = "dreamTracker"
const charset = "utf8"

const ConnectionString = username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=" + charset
const DriverName = "mysql"
