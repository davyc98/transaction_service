## Architectural Approach
### Doman Driven Architecture
Using Clean Code Architectural Approach, specifically addopting [go-clean-arch](https://github.com/bxcodec/go-clean-arch/tree/v1) 

This project has 4 Domain layer :

Models Layer
Repository Layer
Usecase Layer
Delivery Layer
The explanation about this project's structure can read from this medium's post :[Testing Clean Architecture on Golang](https://medium.com/hackernoon/golang-clean-archithecture-efd6d7c43047)


The app require to have this:
- Create table on your postgres with this attribute: 
 ``` 
  DB_HOST: 127.0.0.1
  DB_NAME: pos_majoo
  DB_USER: postgres
  DB_PORT: 5432
 ```
- Run this database script, under script/database.sql:
(the script contain creating new table in postgres database)
- This project need connection to database, please change your environment variable accordingly, or you can follow this value:
```
  DB_HOST: 127.0.0.1
  DB_NAME: pos_majoo
  DB_USER: postgres
  DB_PORT: 5432
  DB_PWD: root
 ```