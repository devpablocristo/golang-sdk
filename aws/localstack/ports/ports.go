package ports

import (
	"github.com/aws/aws-sdk-go-v2/aws"
)

type Stack interface {
	Connect() error
	GetCfg() aws.Config
}

type Config interface {
	Validate() error
	GetAWSAccessKeyID() string
	GetAWSSecretAccessKey() string
	GetAWSRegion() string
	GetLocalStackEndpoint() string
}
