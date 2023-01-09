# Table: circleci_pipeline

CircleCI pipelines are the highest-level unit of work, encompassing a project’s full .circleci/config.yml file.

## Examples

### Error details of the pipelines of a project

```sql
select
  id,
  number,
  errors
from
  circleci_pipeline
where
  project_slug = 'gh/fluent-cattle/sp-plugin-test' and
  errors is not null and
  jsonb_array_length(errors) > 0;
```

### Number of pipelines per project

```sql
select
  pr.slug,
  count(pl.*)
from
  circleci_project pr
join
  circleci_pipeline pl
on
  pr.slug = pl.project_slug
group by
  pr.slug;
```

### Pipelines of a project by state

```sql
select
  state,
  count(*)
from
  circleci_pipeline
where
  project_slug = 'gh/fluent-cattle/sp-plugin-test'
group by
  state;
```
