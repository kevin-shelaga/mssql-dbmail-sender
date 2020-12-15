# Microsoft Sql Server Database Mail Sender

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/kevin-shelaga/mssql-dbmail-sender)
[![Go Report Card](https://goreportcard.com/badge/github.com/kevin-shelaga/mssql-dbmail-sender)](https://goreportcard.com/report/github.com/kevin-shelaga/mssql-dbmail-sender)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![codecov](https://codecov.io/gh/kevin-shelaga/mssql-dbmail-sender/branch/master/graph/badge.svg?token=D07EP88G53)](https://codecov.io/gh/kevin-shelaga/mssql-dbmail-sender)
![build](https://github.com/kevin-shelaga/mssql-dbmail-sender/workflows/build/badge.svg)
![Docker Pulls](https://img.shields.io/docker/pulls/kevinshelaga/mssql-dbmail-sender)
![Docker Image Size (latest by date)](https://img.shields.io/docker/image-size/kevinshelaga/mssql-dbmail-sender)

## What

Support Microsoft database mail in the cloud. Add the dbmail tables to your managed sql instance, run this go application in kubernetes.

## How

### Environment Variables

| Environment Variable | Type    | Description                                                                                                                                          |
|---|---|---|---|---|
|   |   |   |   |   |
|   |   |   |   |   |
|   |   |   |   |   |

### Run as cronjob in kubernetes

#### Use Kustomize

```sh
curl -s "https://raw.githubusercontent.com/\
kubernetes-sigs/kustomize/master/hack/install_kustomize.sh"  | bash

chmod u+x ./kustomize

cp -a ./manifests/. .

./kustomize edit set image kevinshelaga/mssql-dbmail-sender:tag=kevinshelaga/mssql-dbmail-sender:latest

./kustomize build . | kubectl -n monitoring apply -f -
```

## Whats left

### TODO

- [ ] More/better tests
- [ ] Helm chart
- [ ] Policies around contributions
