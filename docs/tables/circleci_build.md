# Table: circleci_build

A CircleCI build is a result of a single execution of a workflow.

## Examples

### Last 10 successful builds

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
  circleci_build
where
  status = 'success'
order by
  stop_time desc
limit 10;
```

### Number of failed builds in a repository

```sql
select
  concat(username,'/',reponame) as repository,
  count(1) as failed_builds
from
  circleci_build b
where
  status = 'failed'
group by
  concat(username,'/',reponame)
order by
  failed_builds desc;
```

### Average execution time duration of successful builds of a repository (in seconds)

```sql
select
  ROUND(avg(build_time_millis/1000)) as average_duration
from
  circleci_build
where
  status = 'success' and
  username = 'fluent-cattle' and
  reponame = 'sp-plugin-test'
group by
  status;
```
