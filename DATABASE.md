# MongoDB DataBase Setup

NAME: `roommates`
Collections:
    - `users`


## users Collection
il valore `email` deve essere unico
```
db.users.createIndex({ email: 1 }, { unique: true })
```


