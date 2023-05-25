### TaskManager Application

#### Architecture 

* golang backend 
    * 5 handlers: (CREATE,UPDATE,GET,DELETE,GETALL)
    * Database connection: planetscale db
    * Middleware - error handling + logging
    * Cors configuration
* svelte frontend
* Deployed on netlify