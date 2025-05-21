# go-fiber
> Fiber Golang web application template 

## Project structure
```bash
└───app   
    ├───config         # Project virtual environments and default values
    ├───docs           # Swagger  
    ├───modules
    │   ├───handlers   # Your framework api code (routes, response&request types, etc.)
    │   └───services   # Your mainclasses 
    ├───static         # Web static directory
    └───views          # Html&css stuff if using render
```

## Environment

Store your examples in .env.dev, or default credential ex. *admin:admin*
```bash
│   .env.dev
│   .env.production <---- Never keep this file in repository  🛑
│   .gitignore      <---- Better to put it here               ✅
└───app
```

## Tests

```bash
├───module
│       config.go
│       service.go
│       service_test.go   <---- Testing package says it is ok
```
