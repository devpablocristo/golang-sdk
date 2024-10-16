package sdkawslocal

import (
	"fmt"

	"github.com/devpablocristo/golang/sdk/pkg/aws/localstack/ports"
)

type configSdk struct {
	AWSAccessKeyID     string
	AWSSecretAccessKey string
	AWSRegion          string
	LocalStackEndpoint string
}

func newConfig(awsAccessKeyID, awsSecretAccessKey, awsRegion, localStackEndpoint string) ports.Config {
	return &configSdk{
		AWSAccessKeyID:     awsAccessKeyID,
		AWSSecretAccessKey: awsSecretAccessKey,
		AWSRegion:          awsRegion,
		LocalStackEndpoint: localStackEndpoint,
	}
}

func (c *configSdk) Validate() error {
	if c.AWSAccessKeyID == "" {
		return fmt.Errorf("AWS_ACCESS_KEY_ID is required")
	}
	if c.AWSSecretAccessKey == "" {
		return fmt.Errorf("AWS_SECRET_ACCESS_KEY is required")
	}
	if c.AWSRegion == "" {
		return fmt.Errorf("AWS_REGION is required")
	}
	if c.LocalStackEndpoint == "" {
		return fmt.Errorf("LOCALSTACK_ENDPOINT is required")
	}
	return nil
}

func (c *configSdk) GetAWSAccessKeyID() string {
	return c.AWSAccessKeyID
}

func (c *configSdk) GetAWSSecretAccessKey() string {
	return c.AWSSecretAccessKey
}

func (c *configSdk) GetAWSRegion() string {
	return c.AWSRegion
}

func (c *configSdk) GetLocalStackEndpoint() string {
	return c.LocalStackEndpoint
}
