![image](https://hub.steampipe.io/images/plugins/turbot/circleci-social-graphic.png)

# CircleCI Plugin for Steampipe

Use SQL to query projects, pipelines, workflows, builds and more from CircleCI.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/circleci)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/circleci/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-circleci/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install circleci
```

Run a query:

```sql
select
  concat(username, '/', reponame) as repository,
  branch,
  status,
  build_url
from
  circleci_build
order by
  stop_time desc limit 10;
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-circleci.git
cd steampipe-plugin-circleci
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/circleci.spc
```

Try it!

```
steampipe query
> .inspect circleci
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). Contributions to the plugin are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-circleci/blob/main/LICENSE). Contributions to the plugin documentation are subject to the [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-circleci/blob/main/docs/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [CircleCI Plugin](https://github.com/turbot/steampipe-plugin-circleci/labels/help%20wanted)
