# Table: circleci_workflow

Workflows define a list of jobs and their run order.

## Examples

### Workflow status of a pipeline

```sql
select
  id,
  created_at,
  status
from
  circleci_workflow
where
  pipeline_id = 'f43cc52a-c7eb-4a72-a05f-399c8577bb3e'
order by
  created_at desc;
```

### Workflow duration of a pipeline

```sql
select
  id,
  project_slug,
  extract(seconds
from
  (
    stopped_at - created_at
  )
) as duration_in_seconds
from
  circleci_workflow
where
  pipeline_id = 'f43cc52a-c7eb-4a72-a05f-399c8577bb3e';
```
