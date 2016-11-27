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
* ~~Write configuration to JSON file~~
* JSON
  * **GET requests**
    * **Ready**
      * ~~Deployment counter~~
      * Manual override
        * Force zone on for duration
        * Force zone on until turned off
        * Reset configuration
    * ~~Configuration~~
    * ~~Array of programs~~
    * ~~Single program~~
    * ~~Array of steps in program~~
    * ~~Single step in program~~
  * **POST requests**
    * ~~Replace configuation~~
    * ~~Replace array of programs~~
    * ~~Replace single program~~
    * ~~Replace array of steps in program~~
    * ~~Replace single step~~
    * Manual override
      * Force zone on for duration
      * Force zone on until turned off
      * Reset configuration
* Raw
  * GET requests
    * Ready
      * Deployment counter
      * Manual override
        * Force zone on for duration
        * Force zone on until turned off
        * Reset configuration
    * Configuration
    * Array of programs
    * Single program
    * Array of steps in program
    * Single step in program
  * POST requests
    * Replace configuation
    * Replace array of programs
    * Replace single program
    * Replace array of steps in program
    * Replace single step
    * Manual override
      * Force zone on for duration
      * Force zone on until turned off
      * Reset configuration
* Determine whether a request is JSON or raw.
  * *Instead of using a leading byte in the request, we might be better off using a header or URL prefix to determine which one should be used*
* ~~Split source into multiple files~~

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

* Returns the program containing the provided program ID.

**Example:**

URL: `http://api.example.com/programs/1`

Response body:

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

* Returns an array of steps in the program containing the provided program ID.

---

`http://api.example.com/programs/{programId}/steps/{stepIndex}`

* Returns the step located at the provided index from the program containing the provided program ID.

**Example:**

URL: `http://api.example.com/programs/1/steps/0`

Response body:

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

`http://api.example.com/config`

* Replaces the entire configuration.

**Example:**

URL: `http://api.example.com/config`

Request body:

```
{
  "deploymentCounter": 0,
  "url": "localhost",
  "port": 4000,
  "programCount": 1,
  "programs": [
    {
      "id": 1,
      "enabled": true,
      "daysOfWeek": [
        false,
        true,
        true,
        true,
        true,
        true,
        false
      ],
      "stepCount": 1,
      "steps": [
        {
          "zones": [
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
          "startTime": 420,
          "duration": 300
        }
      ]
    }
  ]
}
```

Response body:

```
{
  "deploymentCounter": 1,
  "url": "localhost",
  "port": 4000,
  "programCount": 1,
  "programs": [
    {
      "id": 1,
      "enabled": true,
      "daysOfWeek": [
        false,
        true,
        true,
        true,
        true,
        true,
        false
      ],
      "stepCount": 1,
      "steps": [
        {
          "zones": [
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
          "startTime": 420,
          "duration": 300
        }
      ]
    }
  ]
}
```

---

`http://api.example.com/programs`

* Replaces the currently stored programs.

**Example:**

URL: `http://api.example.com/programs`

Request body:

```
[
  {
    "id": 1,
    "enabled": true,
    "daysOfWeek": [
      false,
      true,
      true,
      true,
      true,
      true,
      false
    ],
    "stepCount": 2,
    "steps": [
      {
        "zones": [
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
        "startTime": 420,
        "duration": 60
      },
      {
        "zones": [
          false,
          false,
          false,
          false,
          false,
          true,
          true,
          false,
          false
        ],
        "startTime": 420,
        "duration": 30
      }
    ]
  },
  {
    "id": 123,
    "enabled": true,
    "daysOfWeek": [
      false,
      false,
      false,
      false,
      false,
      true,
      false
    ],
    "stepCount": 1,
    "steps": [
      {
        "zones": [
          true,
          true,
          true,
          true,
          true,
          true,
          true,
          true,
          true
        ],
        "startTime": 420,
        "duration": 300
      }
    ]
  }
]
```

Response body:

```
{
  "deploymentCounter": 5,
  "url": "localhost",
  "port": 4000,
  "programCount": 2,
  "programs": [
    {
      "id": 1,
      "enabled": true,
      "daysOfWeek": [
        false,
        true,
        true,
        true,
        true,
        true,
        false
      ],
      "stepCount": 2,
      "steps": [
        {
          "zones": [
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
          "startTime": 420,
          "duration": 60
        },
        {
          "zones": [
            false,
            false,
            false,
            false,
            false,
            true,
            true,
            false,
            false
          ],
          "startTime": 420,
          "duration": 30
        }
      ]
    },
    {
      "id": 123,
      "enabled": true,
      "daysOfWeek": [
        false,
        false,
        false,
        false,
        false,
        true,
        false
      ],
      "stepCount": 1,
      "steps": [
        {
          "zones": [
            true,
            true,
            true,
            true,
            true,
            true,
            true,
            true,
            true
          ],
          "startTime": 420,
          "duration": 300
        }
      ]
    }
  ]
}
```

---

`http://api.example.com/programs/{programId}`

* Replaces the program containing the provided program ID.

**Example:**

URL: `http://api.example.com/programs/1`

Request body:

```
{
  "id": 1,
  "enabled": true,
  "daysOfWeek": [
    false,
    false,
    false,
    false,
    false,
    false,
    true
  ],
  "stepCount": 1,
  "steps": [
    {
      "zones": [
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
      "startTime": 420,
      "duration": 60
    }
  ]
}
```

Response body:

```
{
  "deploymentCounter": 5,
  "url": "localhost",
  "port": 4000,
  "programCount": 1,
  "programs": [
    {
      "id": 1,
      "enabled": true,
      "daysOfWeek": [
        false,
        false,
        false,
        false,
        false,
        false,
        true
      ],
      "stepCount": 1,
      "steps": [
        {
          "zones": [
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
          "startTime": 420,
          "duration": 60
        }
      ]
    }
  ]
}
```

---

`http://api.example.com/programs/{programId}/steps`

* Replaces the steps in the program containing the provided program ID.

**Example:**

URL: `http://api.example.com/programs/1/steps`

Request body:

```
[
  {
    "zones": [
      true,
      true,
      true,
      true,
      true,
      true,
      true,
      true,
      true
    ],
    "startTime": 420,
    "duration": 600
  }
]
```

Response body:

```
{
  "deploymentCounter": 5,
  "url": "localhost",
  "port": 4000,
  "programCount": 1,
  "programs": [
    {
      "id": 1,
      "enabled": true,
      "daysOfWeek": [
        false,
        true,
        true,
        true,
        true,
        true,
        false
      ],
      "stepCount": 1,
      "steps": [
        {
          "zones": [
            true,
            true,
            true,
            true,
            true,
            true,
            true,
            true,
            true
          ],
          "startTime": 420,
          "duration": 600
        }
      ]
    }
  ]
}
```

---

`http://api.example.com/programs/{programId}/steps/{stepIndex}`

* Replaces the step located at the provided index in the program containing the provided program ID.

**Example:**

URL: `http://api.example.com/programs/1/steps/0`

Request body:

```
{
  "zones": [
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
  "startTime": 420,
  "duration": 300
}
```

Response body:

```
{
  "deploymentCounter": 5,
  "url": "localhost",
  "port": 4000,
  "programCount": 1,
  "programs": [
    {
      "id": 1,
      "enabled": true,
      "daysOfWeek": [
        false,
        true,
        true,
        true,
        true,
        true,
        false
      ],
      "stepCount": 5,
      "steps": [
        {
          "zones": [
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
          "startTime": 420,
          "duration": 300
        },
        {
          "zones": [
            false,
            false,
            false,
            false,
            false,
            true,
            true,
            false,
            false
          ],
          "startTime": 420,
          "duration": 30
        },
        {
          "zones": [
            false,
            false,
            false,
            false,
            false,
            false,
            false,
            true,
            true
          ],
          "startTime": 420,
          "duration": 15
        },
        {
          "zones": [
            false,
            false,
            false,
            false,
            true,
            false,
            false,
            false,
            false
          ],
          "startTime": 480,
          "duration": 60
        },
        {
          "zones": [
            false,
            false,
            false,
            false,
            false,
            true,
            true,
            true,
            true
          ],
          "startTime": 1200,
          "duration": 120
        }
      ]
    }
  ]
}
```
