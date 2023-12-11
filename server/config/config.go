package config

import "os"

var Version = "1.2.3"
var Env = os.Getenv("ENV")
var Port = os.Getenv("PORT")

var PostgresHost = os.Getenv("POSTGRES_HOST")
var PostgresUser = os.Getenv("POSTGRES_USER")
var PostgresPassword = os.Getenv("POSTGRES_PASSWORD")
var PostgresPort = os.Getenv("POSTGRES_PORT")

var DiscordPrefix = os.Getenv("DISCORD_PREFIX")
var DiscordToken = os.Getenv("DISCORD_TOKEN")

var FirebaseProjectID = os.Getenv("FIREBASE_PROJECT_ID")
var FirebaseServiceAccountEncoded = os.Getenv("FIREBASE_SERVICE_ACCOUNT")
