# growth-pull

A dockercfg generator for Codeship.

Must have access to AWS credentials as per the SDK [documentation](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html).
Normally via environment variables:

- `AWS_DEFAULT_REGION=eu-west-1`
- `AWS_ACCESS_KEY_ID=`
- `AWS_SECRET_ACCESS_KEY=`

Will authenticate to that account's ECR, or you can provide a space separated list of Account IDs in `AWS_ECR_REGISTRY_IDS`.

You can optionally provide `DOCKER_HUB_USERNAME` and `DOCKER_HUB_PASSWORD` to also authenticate with Docker Hub.


## Example `codeship-services.yml`

```yaml
ecr:
  image: ghcr.io/tailsdotcom/growth-pull:latest
  encrypted_env_file: aws.env.encrypted

build:
  dockercfg_service: ecr
  image: 123.dkr.ecr.eu-west-1.amazonaws.com/foo
```
