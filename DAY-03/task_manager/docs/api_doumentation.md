
## Endpoints

### Get All Tasks
- **Endpoint:** `GET /tasks`
- **Description:** Retrieves a list of all tasks.
- **Response:**
  - **Status Code:** `200 OK`
  - **Content-Type:** `application/json`
  - **Body:** 
    ```json
    [
      {
        "id": 1,
        "title": "Sample Task",
        "description": "This is a sample task",
        "status": "pending"
      },
      ...
    ]
    ```

### Get Task by ID
- **Endpoint:** `GET /tasks/:id`
- **Description:** Retrieves a specific task by its ID.
- **Path Parameters:**
  - `id` (integer): The ID of the task.
- **Response:**
  - **Status Code:** `200 OK`
  - **Content-Type:** `application/json`
  - **Body:**
    ```json
    {
      "id": 1,
      "title": "Sample Task",
      "description": "This is a sample task",
      "status": "pending"
    }
    ```
  - **Status Code:** `400 Bad Request`
    - **Body:**
      ```json
      {
        "error": "Invalid task ID"
      }
      ```
  - **Status Code:** `404 Not Found`
    - **Body:**
      ```json
      {
        "error": "Task not found"
      }
      ```

### Create Task
- **Endpoint:** `POST /tasks`
- **Description:** Creates a new task.
- **Request Body:**
  - **Content-Type:** `application/json`
  - **Body:**
    ```json
    {
      "title": "New Task",
      "description": "Description of the new task",
      "status": "pending"
    }
    ```
- **Response:**
  - **Status Code:** `201 Created`
  - **Content-Type:** `application/json`
  - **Body:**
    ```json
    {
      "id": 1,
      "title": "New Task",
      "description": "Description of the new task",
      "status": "pending"
    }
    ```
  - **Status Code:** `400 Bad Request`
    - **Body:**
      ```json
      {
        "error": "Error message detailing what went wrong"
      }
      ```

### Update Task
- **Endpoint:** `PUT /tasks/:id`
- **Description:** Updates an existing task by its ID.
- **Path Parameters:**
  - `id` (integer): The ID of the task to update.
- **Request Body:**
  - **Content-Type:** `application/json`
  - **Body:**
    ```json
    {
      "title": "Updated Task Title",
      "description": "Updated description",
      "status": "completed"
    }
    ```
- **Response:**
  - **Status Code:** `200 OK`
  - **Content-Type:** `application/json`
  - **Body:**
    ```json
    {
      "id": 1,
      "title": "Updated Task Title",
      "description": "Updated description",
      "status": "completed"
    }
    ```
  - **Status Code:** `400 Bad Request`
    - **Body:**
      ```json
      {
        "error": "Error message detailing what went wrong"
      }
      ```
  - **Status Code:** `404 Not Found`
    - **Body:**
      ```json
      {
        "error": "Task not found"
      }
      ```

### Delete Task
- **Endpoint:** `DELETE /tasks/:id`
- **Description:** Deletes a specific task by its ID.
- **Path Parameters:**
  - `id` (integer): The ID of the task to delete.
- **Response:**
  - **Status Code:** `204 No Content`
  - **Content-Type:** `application/json`
  - **Body:** `null`
  - **Status Code:** `400 Bad Request`
    - **Body:**
      ```json
      {
        "error": "Invalid task ID"
      }
      ```
  - **Status Code:** `404 Not Found`
    - **Body:**
      ```json
      {
        "error": "Task not found"
      }
      ```

## Error Handling
- **Common Error Responses:**
  - `400 Bad Request`: The request was invalid or cannot be processed.
  - `404 Not Found`: The requested resource was not found.

## Notes
- All date fields should follow the ISO 8601 format.
- The status field can be one of: "pending", "in-progress", "completed".

