# Sprinkler Controller - API

The sprinkler controller API is written in Go. It will sit on the server and take requests to display and change different aspects of the current setup of the sprinkler system.

## Current Progress

### Key:

* Not done
* **In progress**
* *Unclear on what to do*
* ~~Done~~

### List:

* ~~Read configuration from JSON file~~
* JSON
  * ~~GET requests~~
    * ~~Deployment counter~~
    * ~~Configuration~~
    * ~~Array of programs~~
    * ~~Single program~~
    * ~~Array of steps in program~~
    * ~~Single step in program~~
  * **POST requests**
    * Replace array of programs
    * Replace single program
    * Replace array of steps in program
    * Replace single step
* Raw
  * GET requests
    * Deployment counter
    * Configuration
    * Array of programs
    * Single program
    * Array of steps in program
    * Single step in program
  * POST requests
    * Replace array of programs
    * Replace single program
    * Replace array of steps in program
    * Replace single step
* Determine whether a request is JSON or raw.
  * *Instead of using a leading byte in the request, we might be better off using a header or URL prefix to determine which one should be used*
* Split source into multiple files

## Usage

### General

Requests will be done using the GET and POST methods on a given path. The currently implemented paths and how to use them are listed below.

### Methods

#### GET

`http://api.example.com/`

* Returns the current deployment counter.

---

`http://api.example.com/config`

* Returns the entire configuration of the current deployment, including the deployment counter, the URL and port the API is listening on, the program count, and an array of the currently stored programs.

---

`http://api.example.com/programs`

* Returns an array of the currently stored programs.

---

`http://api.example.com/programs/{programId}`

* Returns the program with the provided program ID.

**Example:**

URL: `http://api.example.com/programs/1`

Response:

```
{
  "id:1,
  "enabled":true,
  "daysOfWeek":[
    false,
    true,
    true,
    true,
    true,
    true,
    false
  ],
  "stepCount":1,
  "steps":[
    {
      "zones":[
        true,
        true,
        true,
        true,
        true,
        false,
        false,
        false,
        false
      ],
      "startTime":420,
      "duration":60
    }
  ]
}
```

---

`http://api.example.com/programs/{programId}/steps`

* Returns an array of steps in the program with the provided program ID.

---

`http://api.example.com/programs/{programId}/steps/{stepIndex}`

* Returns the step located at the provided index from the program with the provided program ID.

**Example:**

URL: `http://api.example.com/programs/1/steps/0`

Response:

```
{
  "zones":[
    true,
    true,
    true,
    true,
    true,
    false,
    false,
    false,
    false
  ],
  "startTime":420,
  "duration":60
}
```

---

#### POST

`http://api.example.com/`

* Not implemented