## Dockerfiles

This repository contains a collection of Dockerfiles for various programming languages, applications and development environments.

### Python

See [Python Dockerfile](python/Dockerfile) for more details.

Make sure to edit [pyproject.toml](python/pyproject.toml) in case you want to use a private package, e.g. from Azure DevOps:

```
[[tool.poetry.source]]
name = "private_repo"
url = "https://pkgs.dev.azure.com/{ORG}/_packaging/basic/pypi/simple/"
```

Or use the following command to add a new source:

```shell
poetry source add private_repo '{URL}'
```

#### Github Actions

See [.github/workflows/python-jobs.yaml](.github/workflows/python-jobs.yaml) for more details.

Currently the job fetches its secrets from Github Actions secrets. In an ideal world it would be better to use OIDC to login to your vault provider and fetch the secrets from there.

#### With private Azure DevOps package

```shell
azureDevopsResourceId="499b84ac-1321-427f-aa17-267ca6975798"
TOKEN=$(az account get-access-token --resource $azureDevopsResourceId | jq -r '.accessToken')
echo $TOKEN > /tmp/python-token
echo "basic" > /tmp/python-user

docker buildx build . \
    --load \
    --cache-from=type=registry,ref=ghcr.io/your/image:cache \
    --cache-to=type=registry,ref=ghcr.io/your/image:cache,mode=max \
    --platform linux/amd64 \
    --secret id=POETRY_AUTH_PASS,src=/tmp/python-token \
    --secret id=POETRY_AUTH_USER,src=/tmp/python-user \
    --target unit-test \
    --progress=plain \
    -t a-service-tests\

docker run -it a-service-tests
```
