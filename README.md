# growth-pull

A dockercfg generator for Codeship.

### Tooling

Unless noted these are pre-configured and used based on their documentation.

- **Planning and issues**:
  [Clubhouse](https://app.clubhouse.io/tails/stories/space/12868/everything).

- **Version control**:
  [Github](https://github.com/tailsdotcom/growth-pull),
  [Flow](https://guides.github.com/introduction/flow/).

- **Continuous integration**:
  [Codeship](https://app.codeship.com/projects/7f08da42-b905-4f2e-a09a-bed0ac32973e).

- **Deployment**:
  Tag a release in the from `v0.1.0` and wait for Github Actions.

### Settings

Must have access to AWS credentials as per the SDK [documentation](https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html).
Normally via environment variables:

- `AWS_REGION=eu-west-1`
- `AWS_ACCESS_KEY_ID=`
- `AWS_SECRET_ACCESS_KEY=`

Will authenticate to that account's ECR, or you can provide a space separated list of Account IDs in `AWS_ECR_REGISTRY_IDS`.

You can optionally provide `DOCKER_HUB_USERNAME` and `DOCKER_HUB_PASSWORD` to also authenticate with Docker Hub.
