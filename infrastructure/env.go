package infrastructure

import (
	"os"
)

// Env has environment stored
type Env struct {
	ServerPort  string
	Environment string
	LogOutput   string
	DBUsername  string
	DBPassword  string
	DBHost      string
	DBPort      string
	DBName      string
	SentryDSN   string

	StorageBucketName string

	AdminEmail string
	AdminPass  string

	MailClientID     string
	MailClientSecret string
	MailAccesstoken  string
	MailRefreshToken string

	AWS_S3_REGION  string
	AWS_S3_BUCKET  string
	AWS_ACCESS_KEY string
	AWS_SECRET_KEY string

	TwilioBaseURL   string
	TwilioSID       string
	TwilioAuthToken string
	TwilioSMSFrom   string
}

// NewEnv creates a new environment
func NewEnv() Env {
	env := Env{}
	env.LoadEnv()
	return env
}

// LoadEnv loads environment
func (env *Env) LoadEnv() {
	env.ServerPort = os.Getenv("SERVER_PORT")
	env.Environment = os.Getenv("ENVIRONMENT")
	env.LogOutput = os.Getenv("LOG_OUTPUT")
	env.DBUsername = os.Getenv("DB_USERNAME")
	env.DBPassword = os.Getenv("DB_PASSWORD")
	env.DBHost = os.Getenv("DB_HOST")
	env.DBPort = os.Getenv("DB_PORT")
	env.DBName = os.Getenv("DB_NAME")
	env.SentryDSN = os.Getenv("SENTRY_DSN")
	env.StorageBucketName = os.Getenv("STORAGE_BUCKET_NAME")
	env.AdminEmail = os.Getenv("ADMIN_EMAIL")
	env.AdminPass = os.Getenv("ADMIN_PASS")
	env.MailClientID = os.Getenv("MAIL_CLIENT_ID")
	env.MailClientSecret = os.Getenv("MAIL_CLIENT_SECRET")
	env.MailAccesstoken = os.Getenv("MAIL_ACCESSTOKEN")
	env.MailRefreshToken = os.Getenv("MAIL_REFRESH_TOKEN")
	env.AWS_S3_REGION = os.Getenv("AWS_S3_REGION")
	env.AWS_S3_BUCKET = os.Getenv("AWS_S3_BUCKET")
	env.AWS_ACCESS_KEY = os.Getenv("AWS_ACCESS_KEY")
	env.AWS_SECRET_KEY = os.Getenv("AWS_SECRET_KEY")
	env.TwilioBaseURL = os.Getenv("TWILIO_BASE_URL")
	env.TwilioAuthToken = os.Getenv("TWILIO_AUTH_TOKEN")
	env.TwilioSID = os.Getenv("TWILIO_SID")
	env.TwilioSMSFrom = os.Getenv("TWILIO_SMSFROM")
}
