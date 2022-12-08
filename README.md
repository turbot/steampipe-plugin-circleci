![image](https://hub.steampipe.io/images/plugins/turbot/circleci-social-graphic.png)

# CircleCI Plugin for Steampipe

Use SQL to query projects, pipelines, workflows, builds and more from CircleCI.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/circleci)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/circleci/tables)
- Community: [Slack Channel](https://steampipe.io/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-circleci/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install circleci
```

Run a query:

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

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-circleci/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [CircleCI Plugin](https://github.com/turbot/steampipe-plugin-circleci/labels/help%20wanted)
