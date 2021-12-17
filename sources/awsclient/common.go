package awsclient

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	log "github.com/sirupsen/logrus"
)

// CreateNewSession return new AWS session
func CreateNewSession(accessKey, secretKey, sessionToken, region string) *session.Session {
	var credentialsAWS *credentials.Credentials

	// Use separate call for AWS credentials defined in config.yaml
	// Otherwise environment variables will be used
	if accessKey != "" && secretKey != "" {
		log.Info("Using AccessKey or SecretKey defined in config")
		credentialsAWS = credentials.NewStaticCredentials(accessKey, secretKey, sessionToken)
	}

	sess := session.Must(session.NewSession(&aws.Config{
		Region:      &region,
		Credentials: credentialsAWS,
	}))

	return sess
}
