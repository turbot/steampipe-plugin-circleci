---
organization: Turbot
category: ["saas"]
icon_url: "/images/plugins/turbot/circleci.svg"
brand_color: "#04AA51"
display_name: "CircleCI"
name: "circleci"
description: "Steampipe plugin for querying resource projects, pipelines, builds and more from CircleCI."
og_description: "Query CircleCI with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/circleci-social-graphic.png"
---

# CircleCI + Steampipe

[CircleCI](https://www.circleci.com/) is the leading open source automation server, CircleCI provides hundreds of plugins to support building, deploying and automating any project.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example:

```sql
select
  username as "organization",
  reponame,
  branch,
  build_time_millis,
  status,
  author_name,
  build_url
from
  circleci_build;
```

```
+---------------+----------------+------------------------+-------------------+---------+-----------------+--------------------------------------------------------+
| organization  | reponame       | branch                 | build_time_millis | status  | author_name     | build_url                                              |
+---------------+----------------+------------------------+-------------------+---------+-----------------+--------------------------------------------------------+
| fluent-cattle | warm-beagle    | circleci-project-setup | 2794              | success | Noah            | https://circleci.com/gh/fluent-cattle/warm-beagle/52   |
| funny-liger   | sp-plugin-test | circleci-project-setup | 604185            | success | Luis            | https://circleci.com/gh/funny-liger/sp-plugin-test/90  |
| fluent-cattle | sp-plugin-test | circleci-project-setup | 6696              | failed  | Oliver          | https://circleci.com/gh/fluent-cattle/sp-plugin-test/8 |
+---------------+----------------+------------------------+-------------------+---------+-----------------+--------------------------------------------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/circleci/tables)**

## Get started

### Install

Download and install the latest CircleCI plugin:

```bash
steampipe plugin install circleci
```

### Credentials

| Item        | Description                                                                                                                                                                                                                                                                                 |
|-------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Credentials | CircleCI requires an [API token](https://www.circleci.com/doc/book/using/using-credentials/) for all requests.                                                                                                                                                                                 |
| Permissions | API tokens have the same permissions as the user who creates them, and if the user permissions change, the API token permissions also change.                                                                                                                                               |
| Radius      | Each connection represents a single CircleCI Installation.                                                                                                                                                                                                                                   |
| Resolution  | 1. With configuration provided in connection in steampipe _**.spc**_ config file.<br />2. With circleci environment variables.<br />3. An circleci.yaml file in a .circleci folder in the current user's home directory _**(~/.circleci/circleci.yaml or %userprofile\.circleci\circleci.yaml)**_. |

### Configuration

Installing the latest circleci plugin will create a config file (~/.steampipe/config/circleci.spc) with a single connection named circleci:

```hcl
connection "circleci" {
  plugin = "circleci"

  # Get your API token from CircleCI https://circleci.com/docs/api-developers-guide/#add-an-api-token

  # api_token = "1234ee38fc6943f6cb9537a564e9a6dac6ef1463"
}
```

By default, all options are commented out in the default connection, thus Steampipe will resolve your credentials using the same order as mentioned in [Credentials](#credentials). This provides a quick way to get started with Steampipe, but you will probably want to customize your experience using configuration options for querying multiple organizations, configuring credentials from your circleci configuration files, [environment variables](#credentials-from-environment-variables), etc.


## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-circleci
- Community: [Slack Channel](https://steampipe.io/community/join)

## Configuring CircleCI Credentials

### Credentials from Environment Variables

The CircleCI plugin will use the standard CircleCI environment variables to obtain credentials **only if other arguments (`api_token`) are not specified** in the connection:

#### API Token

```sh
export CIRCLE_TOKEN=1234ee38fc6943f6cb9537a564e9a6dac6ef1463
```
