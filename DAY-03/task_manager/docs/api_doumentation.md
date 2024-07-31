# Task Management API Documentation

## Endpoints

### GET /api/v1/tasks
- Description: Get a list of all tasks.
- Response: 200 OK
```json
[
    {
        "id": 1,
        "title": "Task 1",
        "description": "Description 1",
        "due_date": "2023-10-01",
        "status": "pending"
    }
]