# S3 State Bucket for Pulumi (Go)

This directory contains Pulumi code written in Go to create an S3 bucket named exactly `pulumi-state-bucket-jason4151` for storing Pulumi state in AWS us-east-2. It uses a local state backend initially, then migrates to the S3 bucket.

## Prerequisites
- Pulumi CLI installed (`brew install pulumi`)
- Go installed (`brew install go`, verify with `go version`)
- AWS CLI configured (`aws configure`)
- Git installed

## Files
- `main.go`: Creates the S3 bucket.
- `Pulumi.yaml`: Project config.
- `go.mod`: Go module dependencies.
- `README.md`: This file.

## Deployment Steps
### 1. Set Local Backend
```bash
cd s3_state_bucket
pulumi login --local
```
### 2. Initialize Go Project
```bash
echo 'name: s3_state_bucket
runtime: go' > Pulumi.yaml
go mod init s3_state_bucket
```
### 3. Install Dependencies
```bash
go mod tidy
```
### 5. Set AWS Region
```bash
pulumi config set projectName s3_state_bucket
pulumi config set aws:region us-east-2
```
### 6. Deploy
```bash
pulumi stack init dev
pulumi up
```
### 7. Migrate to S3
```bash
pulumi stack export --file s3-state.json
pulumi login s3://pulumi-state-bucket-jason4151
pulumi stack init dev
pulumi stack import --file s3-state.json
rm s3-state.json
```

