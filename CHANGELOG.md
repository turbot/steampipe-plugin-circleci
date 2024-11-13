## v1.0.1 [2024-11-13]

_Bug fixes_

- Fixed the `JSON unmarshalling` error when querying the `trigger_parameters` column of the `circleci_pipeline` table. ([#53](https://github.com/turbot/steampipe-plugin-circleci/pull/53))

## v1.0.0 [2024-10-22]

There are no significant changes in this plugin version; it has been released to align with [Steampipe's v1.0.0](https://steampipe.io/changelog/steampipe-cli-v1-0-0) release. This plugin adheres to [semantic versioning](https://semver.org/#semantic-versioning-specification-semver), ensuring backward compatibility within each major version.

_Dependencies_

- Recompiled plugin with Go version `1.22`. ([#48](https://github.com/turbot/steampipe-plugin-circleci/pull/48))
- Recompiled plugin with [steampipe-plugin-sdk v5.10.4](https://github.com/turbot/steampipe-plugin-sdk/blob/develop/CHANGELOG.md#v5104-2024-08-29) that fixes logging in the plugin export tool. ([#48](https://github.com/turbot/steampipe-plugin-circleci/pull/48))

## v0.5.0 [2023-12-12]

_What's new?_

- The plugin can now be downloaded and used with the [Steampipe CLI](https://steampipe.io/docs), as a [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview), as a [SQLite extension](https://steampipe.io/docs//steampipe_sqlite/overview) and as a standalone [exporter](https://steampipe.io/docs/steampipe_export/overview). ([#27](https://github.com/turbot/steampipe-plugin-circleci/pull/27))
- The table docs have been updated to provide corresponding example queries for Postgres FDW and SQLite extension. ([#27](https://github.com/turbot/steampipe-plugin-circleci/pull/27))
- Docs license updated to match Steampipe [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-circleci/blob/main/docs/LICENSE). ([#27](https://github.com/turbot/steampipe-plugin-circleci/pull/27))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.8.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v580-2023-12-11) that includes plugin server encapsulation for in-process and GRPC usage, adding Steampipe Plugin SDK version to `_ctx` column, and fixing connection and potential divide-by-zero bugs. ([#26](https://github.com/turbot/steampipe-plugin-circleci/pull/26))

## v0.4.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#17](https://github.com/turbot/steampipe-plugin-circleci/pull/17))

## v0.4.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#15](https://github.com/turbot/steampipe-plugin-circleci/pull/15))
- Recompiled plugin with Go version `1.21`. ([#15](https://github.com/turbot/steampipe-plugin-circleci/pull/15))

## v0.3.0 [2023-06-30]

_What's new?_

- New tables added
  - [circleci_insights_workflow_run](https://hub.steampipe.io/plugins/turbot/circleci/tables/circleci_insights_workflow_run) ([#8](https://github.com/turbot/steampipe-plugin-circleci/pull/8)) (Thanks [@leonid-panich](https://github.com/leonid-panich) for the contribution!)

## v0.2.0 [2023-04-11]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which adds go-getter support to dynamic tables. ([#6](https://github.com/turbot/steampipe-plugin-circleci/pull/6))

## v0.1.0 [2023-01-27]

_What's new?_

- New tables added
  - [circleci_context](https://hub.steampipe.io/plugins/turbot/circleci/tables/circleci_context) ([#3](https://github.com/turbot/steampipe-plugin-circleci/pull/3))
  - [circleci_context_environment_variable](https://hub.steampipe.io/plugins/turbot/circleci/tables/circleci_context_environment_variable) ([#3](https://github.com/turbot/steampipe-plugin-circleci/pull/3))
  - [circleci_organization](https://hub.steampipe.io/plugins/turbot/circleci/tables/circleci_organization) ([#3](https://github.com/turbot/steampipe-plugin-circleci/pull/3))

## v0.0.1 [2023-01-10]

_What's new?_

- New tables added

  - [circleci_build](https://hub.steampipe.io/plugins/turbot/circleci/tables/circleci_build)
  - [circleci_pipeline](https://hub.steampipe.io/plugins/turbot/circleci/tables/circleci_pipeline)
  - [circleci_project](https://hub.steampipe.io/plugins/turbot/circleci/tables/circleci_project)
  - [circleci_workflow](https://hub.steampipe.io/plugins/turbot/circleci/tables/circleci_workflow)
