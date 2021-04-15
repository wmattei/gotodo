# TODO em go


## Endpoints
[X] `[POST] /tasks`
Cria uma task nova.

[ ] `[GET] /tasks`
Lista todas as tasks. Com filtros por `done` e `title` (texto livre)

# junior
[ ] `[GET] /tasks/{TASK_ID}`
Busca uma task por ID.

# junior
[ ] `[PUT] /tasks/{TASK_ID}`
Atualiza o conteúdo de uma task pelo seu ID

[ ] `[DELETE] /tasks/{TASK_ID}`
Delete uma task pelo seu ID

## Task model
```json
{
  "title": "string",
  "description": "string",
  "done": "boolean",
}
```

##  TODO
- Body validation
- Env variables
- Logger
- Exception Handler
