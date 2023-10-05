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
