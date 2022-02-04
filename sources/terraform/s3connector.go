package terraform

import (
	"bytes"
	"fmt"

	"github.com/hashicorp/terraform/states/statefile"
	"github.com/isan-rivkin/kubefigure/sources/awsclient"
)

// data "terraform_remote_state" "some_state" {
// 	backend = "s3"

// 	config = {
// 	  bucket = "..."
// 	  key    = "..."
// 	  region = "us-east-1"
// 	}
//   }

type S3RemoteStateConnector struct {
	bucket string
	key    string
	// todo: since the awsClient already has session the region here is useless it can error if sess.region =! region
	region   string
	s3Client *awsclient.S3Api
}

func NewDefaultS3RemoteStateConnector(bucket, key, region string) RemoteStateConnector {
	s3Api := awsclient.NewS3API(awsclient.CreateNewSession("", "", "", region))
	return NewS3RemoteStateConnector(s3Api, bucket, key, region)
}

func NewS3RemoteStateConnector(s3api *awsclient.S3Api, bucket, key, region string) RemoteStateConnector {
	return &S3RemoteStateConnector{
		s3Client: s3api,
		region:   region,
		bucket:   bucket,
		key:      key,
	}
}

func (sc *S3RemoteStateConnector) DownloadAsStatefile() (*statefile.File,[]byte,  error) {
	data, err := sc.Download()

	if err != nil {
		return nil,nil, err
	}

	reader := bytes.NewReader(data)
	file, err := statefile.Read(reader)

	return file,data, err
}

func (sc *S3RemoteStateConnector) Download() ([]byte, error) {
	data, err := sc.s3Client.GetObject(sc.bucket, sc.key)

	if err != nil {
		return nil, fmt.Errorf("could not download s3 state, Error, %v", err)
	}
	return data, nil

}
