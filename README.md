Go Client for ICanCode event
==========================

Installation
------------

1. Requirements: golang v1.12 and above
2. To init project

    ```bash
    $ cd project_folder
    $ go run cmd/bot/main.go
    ```
    
4. After registering in ICanCode server copy host, username and code from browser's address bar into `handler.go` related variables
3. Your code should be implemented in `pkg/handler.go` 
5. See api docs `http://${DOJOSERVER}/codenjoy-contest/resources/icancode/landing-icancode-contest.html`