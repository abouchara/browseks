package eks

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/eks"
)

type EKSClient struct {
	service *eks.EKS
}

// NewEKSClient creates a new EKS client with given AWS credentials and region.
//
func NewEKSClient(accessKeyID, secretAccessKey, sessionToken, region string) (*EKSClient, error) {
	creds := credentials.NewStaticCredentials(accessKeyID, secretAccessKey, sessionToken)
	sess, err := session.NewSession(&aws.Config{
		Credentials: creds,
		Region:      aws.String(region),
	})
	if err != nil {
		return nil, err
	}

	return &EKSClient{
		service: eks.New(sess),
	}, nil
}

// ListClusters lists the Amazon EKS clusters in your AWS account in the specified region.
//
func (c *EKSClient) ListClusters() ([]string, error) {

	input := &eks.ListClustersInput{}

	result, err := c.service.ListClusters(input)
	if err != nil {
		return nil, err
	}

	return aws.StringValueSlice(result.Clusters), nil
}

// Additional functions to interact with EKS can be added below...
