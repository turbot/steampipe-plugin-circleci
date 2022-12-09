# Table: circleci_build

Result of a single execution of a pipeline.

## Examples

### Amount of failed builds by repository

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
