package main

import (
    "github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        // Create the S3 bucket with a specific name
        bucket, err := s3.NewBucket(ctx, "stateBucket", &s3.BucketArgs{
            Bucket:      pulumi.String("pulumi-state-bucket-jason4151"),
            Acl:         pulumi.String("private"),
            Versioning:  s3.BucketVersioningArgs{Enabled: pulumi.Bool(true)},
            Tags:        pulumi.StringMap{"Name": pulumi.String("pulumi-state-bucket-jason4151")},
        })
        if err != nil {
            return err
        }

        // Export the bucket name
        ctx.Export("state_bucket_name", bucket.ID())
        return nil
    })
}
