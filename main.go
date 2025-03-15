package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Crear un bucket en S3
		bucket, err := s3.NewBucket(ctx, "myBucket", nil)
		if err != nil {
			return err
		}

		// Exportar el nombre del bucket
		ctx.Export("bucketName", bucket.ID())
		return nil
	})
}


