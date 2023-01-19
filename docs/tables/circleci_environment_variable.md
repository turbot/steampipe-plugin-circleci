# Table: circleci_environment_variable

CircleCI environment variables store customer data that is used by projects.

## Examples

### Environment variables in a context

```sql
select
  context_id,
  variable,
  created_at,
  updated_at
from
  circleci_environment_variable
where
  context_id = '60d77d33-a62c-4167-90be-3e02ee38a75b';
```

### Environment variables not updated for more then 90 days
```sql
select
  context_id,
  variable,
  created_at,
  updated_at
from
  circleci_environment_variable
where
  context_id = '60d77d33-a62c-4167-90be-3e02ee38a75b' and
  updated_at < current_date - interval '90' day;
```


