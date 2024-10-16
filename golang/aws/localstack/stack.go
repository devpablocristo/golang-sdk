package sdkawslocal

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"

	"github.com/devpablocristo/sdk/golang/sdk/aws/localstack/ports"
)

var (
	instance  ports.Stack
	once      sync.Once
	initError error
)

type stack struct {
	config ports.Config
	awsCfg aws.Config
}

func newStack(c ports.Config) (ports.Stack, error) {
	once.Do(func() {
		svc := &stack{config: c}
		initError = svc.Connect()
		if initError != nil {
			instance = nil
		} else {
			instance = svc
		}
	})
	return instance, initError
}

func (s *stack) Connect() error {
	awsCfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(s.config.GetAWSRegion()),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			s.config.GetAWSAccessKeyID(), s.config.GetAWSSecretAccessKey(), "",
		)),
	)
	if err != nil {
		return err
	}
	s.awsCfg = awsCfg
	return nil
}

func (s *stack) GetCfg() aws.Config {
	return s.awsCfg
}

// // Example of creating a DynamoDB client with the stack-specific endpoint resolver
// func (s *stack) NewDynamoDBClient() *dynamodb.Client {
// 	return dynamodb.NewFromConfig(s.awsCfg, func(o *dynamodb.Options) {
// 		o.EndpointResolver = dynamodb.EndpointResolverFromURL(s.config.GetLocalStackEndpoint())
// 	})
// }

// // Similarly, for S3:
// func (s *stack) NewS3Client() *s3.Client {
// 	return s3.NewFromConfig(s.awsCfg, func(o *s3.Options) {
// 		o.EndpointResolver = s3.EndpointResolverFromURL(s.config.GetLocalStackEndpoint())
// 	})
// }
