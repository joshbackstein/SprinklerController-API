# Sprinkler Controller - API

The sprinkler controller API is written in Go. It will sit on the server and take requests to display and change different aspects of the current setup of the sprinkler system.

## Current Progress

### Key:

* Not done
* **In progress**
* *Notes or unclear on what to do*
* ~~Done~~

### List:

* ~~Read configuration from JSON file~~
* ~~Write configuration to JSON file~~
* ~~JSON~~
  * ~~GET requests~~
    * ~~Ready~~
      * ~~Deployment counter~~
    * ~~Configuration~~
    * ~~Master control~~
    * ~~Array of programs~~
    * ~~Single program~~
    * ~~Array of steps in program~~
    * ~~Single step in program~~
    * ~~Manual override~~
      * ~~Force zone on until turned off~~
      * ~~Force zone on for duration~~
  * ~~POST requests~~
    * ~~Replace configuation~~
    * ~~Replace master control~~
    * ~~Replace array of programs~~
    * ~~Replace single program~~
    * ~~Replace array of steps in program~~
    * ~~Replace single step~~
    * ~~Manual override~~
      * ~~Force zone on until turned off~~
      * ~~Force zone on for duration~~
* ~~Raw~~
  * *Arduino will be using JSON instead of raw data.*
  * ~~GET requests~~
    * ~~Ready~~
      * ~~Deployment counter~~
    * ~~Configuration~~
    * ~~Master control~~
    * ~~Array of programs~~
    * ~~Single program~~
    * ~~Array of steps in program~~
    * ~~Single step in program~~
    * ~~Manual override~~
      * ~~Force zone on until turned off~~
      * ~~Force zone on for duration~~
  * ~~POST requests~~
    * ~~Replace configuation~~
    * ~~Replace master control~~
    * ~~Replace array of programs~~
    * ~~Replace single program~~
    * ~~Replace array of steps in program~~
    * ~~Replace single step~~
    * ~~Manual override~~
      * ~~Force zone on until turned off~~
      * ~~Force zone on for duration~~
* ~~Determine whether a request is JSON or raw~~
  * ~~*Instead of using a leading byte in the request, we might be better off using a header or URL prefix to determine which one should be used*~~
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

* Returns the entire configuration of the current deployment, including the deployment counter, the local host and port the API is listening on, the external host and port that should be used to access the API, an array of the currently stored programs, and an array of the currently stored manual overrides.

---

`http://api.example.com/master`

* Returns the current master control configuration.

---

`http://api.example.com/programs`

* Returns an array of the currently stored programs.

---

`http://api.example.com/programs/{programIndex}`

* Returns the program at the provided index.

**Example:**

URL: `http://api.example.com/programs/0`

Response body:

```
{
  "enabled": true,
  "name": "Program #1",
  "daysOfWeek": [
    false,
    true,
    true,
    true,
    true,
    true,
    false
  ],
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

---

`http://api.example.com/programs/{programIndex}/steps`

* Returns an array of steps in the program at the provided index.

---

`http://api.example.com/programs/{programIndex}/steps/{stepIndex}`

* Returns the step at the provided index from the program at the provided index.

**Example:**

URL: `http://api.example.com/programs/0/steps/0`

Response body:

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
  "duration": 60
}
```

---

`http://api.example.com/overrides`

* Returns an array of the currently stored manual overrides.

---

`http://api.example.com/overrides/{overrideIndex}`

* Returns the manual override at the provided index.

**Example:**

URL: `http://api.example.com/overrides/0`

Response body:

```
{
  "enabled": false,
  "untilTurnedOff": true,
  "duration": 0
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
  "localHost": "localhost",
  "localPort": 4000,
  "host": "api.example.com",
  "port": 80,
  "master": {
    "enabled": true
  },
  "programs": [
    {
      "enabled": true,
      "name": "Program #1",
      "daysOfWeek": [
        false,
        true,
        true,
        true,
        true,
        true,
        false
      ],
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
  ],
  "overrides": [
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    }
  ]
}
```

Response body:

```
{
  "deploymentCounter": 1,
  "localHost": "localhost",
  "localPort": 4000,
  "host": "api.example.com",
  "port": 80,
  "master": {
    "enabled": true
  },
  "programs": [
    {
      "enabled": true,
      "name": "Program #1",
      "daysOfWeek": [
        false,
        true,
        true,
        true,
        true,
        true,
        false
      ],
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
  ],
  "overrides": [
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    }
  ]
}
```

---

`http://api.example.com/master`

* Replaces the current master control configuration.

**Example:**

URL: `http://api.example.com/master`

Request body:

```
{
  "enabled": false
}
```

Response body:

```
{
  "deploymentCounter": 2,
  "localHost": "localhost",
  "localPort": 4000,
  "host": "api.example.com",
  "port": 80,
  "master": {
    "enabled": false
  },
  "programs": [
    {
      "enabled": true,
      "name": "Program #1",
      "daysOfWeek": [
        false,
        true,
        true,
        true,
        true,
        true,
        false
      ],
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
  ],
  "overrides": [
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
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
    "enabled": true,
    "name": "Program #1",
    "daysOfWeek": [
      false,
      true,
      true,
      true,
      true,
      true,
      false
    ],
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
    "enabled": true,
    "name": "Program #2",
    "daysOfWeek": [
      false,
      false,
      false,
      false,
      false,
      true,
      false
    ],
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
  "deploymentCounter": 2,
  "localHost": "localhost",
  "localPort": 4000,
  "host": "api.example.com",
  "port": 80,
  "master": {
    "enabled": true
  },
  "programs": [
    {
      "enabled": true,
      "name": "Program #1",
      "daysOfWeek": [
        false,
        true,
        true,
        true,
        true,
        true,
        false
      ],
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
      "enabled": true,
      "name": "Program #2",
      "daysOfWeek": [
        false,
        false,
        false,
        false,
        false,
        true,
        false
      ],
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
  ],
  "overrides": [
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    }
  ]
}
```

---

`http://api.example.com/programs/{programIndex}`

* Replaces the program at the provided index.

**Example:**

URL: `http://api.example.com/programs/0`

Request body:

```
{
  "enabled": true,
  "name": "Program #1",
  "daysOfWeek": [
    false,
    false,
    false,
    false,
    false,
    false,
    true
  ],
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
  "deploymentCounter": 2,
  "localHost": "localhost",
  "localPort": 4000,
  "host": "api.example.com",
  "port": 80,
  "master": {
    "enabled": true
  },
  "programs": [
    {
      "enabled": true,
      "name": "Program #1",
      "daysOfWeek": [
        false,
        false,
        false,
        false,
        false,
        false,
        true
      ],
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
  ],
  "overrides": [
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    }
  ]
}
```

---

`http://api.example.com/programs/{programIndex}/steps`

* Replaces the steps in the program at the provided index.

**Example:**

URL: `http://api.example.com/programs/0/steps`

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
  "deploymentCounter": 2,
  "localHost": "localhost",
  "localPort": 4000,
  "host": "api.example.com",
  "port": 80,
  "master": {
    "enabled": true
  },
  "programs": [
    {
      "enabled": true,
      "name": "Program #1",
      "daysOfWeek": [
        false,
        true,
        true,
        true,
        true,
        true,
        false
      ],
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
  ],
  "overrides": [
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    }
  ]
}
```

---

`http://api.example.com/programs/{programIndex}/steps/{stepIndex}`

* Replaces the step at the provided index in the program at the provided index.

**Example:**

URL: `http://api.example.com/programs/0/steps/0`

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
  "deploymentCounter": 2,
  "localHost": "localhost",
  "localPort": 4000,
  "host": "api.example.com",
  "port": 80,
  "master": {
    "enabled": true
  },
  "programs": [
    {
      "enabled": true,
      "name": "Program #1",
      "daysOfWeek": [
        false,
        true,
        true,
        true,
        true,
        true,
        false
      ],
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
  ],
  "overrides": [
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    }
  ]
}
```


---

`http://api.example.com/overrides`

* Replaces the currently stored manual overrides.

**Example:**

URL: `http://api.example.com/overrides`

Request body:

```
[
  {
    "enabled": true,
    "untilTurnedOff": false,
    "duration": 3600
  },
  {
    "enabled": true,
    "untilTurnedOff": true,
    "duration": 0
  },
  {
    "enabled": true,
    "untilTurnedOff": false,
    "duration": 60
  },
  {
    "enabled": true,
    "untilTurnedOff": false,
    "duration": 120
  },
  {
    "enabled": true,
    "untilTurnedOff": false,
    "duration": 180
  },
  {
    "enabled": true,
    "untilTurnedOff": false,
    "duration": 240
  },
  {
    "enabled": true,
    "untilTurnedOff": false,
    "duration": 300
  },
  {
    "enabled": true,
    "untilTurnedOff": false,
    "duration": 360
  },
  {
    "enabled": true,
    "untilTurnedOff": false,
    "duration": 420
  }
]
```

Response body:

```
{
  "deploymentCounter": 2,
  "localHost": "localhost",
  "localPort": 4000,
  "host": "api.example.com",
  "port": 80,
  "master": {
    "enabled": true
  },
  "programs": [
    {
      "enabled": true,
      "name": "Program #1",
      "daysOfWeek": [
        false,
        true,
        true,
        true,
        true,
        true,
        false
      ],
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
  ],
  "overrides": [
    {
      "enabled": true,
      "untilTurnedOff": false,
      "duration": 3600
    },
    {
      "enabled": true,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": true,
      "untilTurnedOff": false,
      "duration": 60
    },
    {
      "enabled": true,
      "untilTurnedOff": false,
      "duration": 120
    },
    {
      "enabled": true,
      "untilTurnedOff": false,
      "duration": 180
    },
    {
      "enabled": true,
      "untilTurnedOff": false,
      "duration": 240
    },
    {
      "enabled": true,
      "untilTurnedOff": false,
      "duration": 300
    },
    {
      "enabled": true,
      "untilTurnedOff": false,
      "duration": 360
    },
    {
      "enabled": true,
      "untilTurnedOff": false,
      "duration": 420
    }
  ]
}
```

---

`http://api.example.com/overrides/{overrideIndex}`

* Replaces the manual override at the provided index.

**Example:**

URL: `http://api.example.com/overrides/0`

Request body:

```
{
  "enabled": true,
  "untilTurnedOff": false,
  "duration": 60
}
```

Response body:

```
{
  "deploymentCounter": 2,
  "localHost": "localhost",
  "localPort": 4000,
  "host": "api.example.com",
  "port": 80,
  "master": {
    "enabled": true
  },
  "programs": [
    {
      "enabled": true,
      "name": "Program #1",
      "daysOfWeek": [
        false,
        true,
        true,
        true,
        true,
        true,
        false
      ],
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
  ],
  "overrides": [
    {
      "enabled": true,
      "untilTurnedOff": false,
      "duration": 60
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    },
    {
      "enabled": false,
      "untilTurnedOff": true,
      "duration": 0
    }
  ]
}
```
