# Table: circleci_context_environment_variable

CircleCI context environment variables store customer data that is used by projects.

## Examples

### Environment variables in a context

```sql
select
  context_id,
  variable,
  created_at,
  updated_at
from
  circleci_context_environment_variable
where
  context_id = '60d77d33-a62c-4167-90be-3e02ee38a75b';
```

### Environment variables not updated for more than 90 days

```sql
select
  context_id,
  variable,
  created_at,
  updated_at 
from
  circleci_context_environment_variable
where
  context_id = '60d77d33-a62c-4167-90be-3e02ee38a75b'
  and updated_at < current_date - interval '90' day;
```
