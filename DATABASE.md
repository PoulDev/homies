# MongoDB DataBase Setup

NAME: `roommates`
Collections:
    - `users`


## users Collection
il valore `email` deve essere unico
il valore `house` deve essere un index altrimenti le ricerche diventano pesantissime
```
db.users.createIndex({ email: 1 }, { unique: true })
db.users.createIndex({ house: 1 })
```

## houses Collection 
il valore `owner` deve essere unico. Ogni persona puo' creare ed essere in una sola casa.
```
db.houses.createIndex({ owner: 1 }, { unique: true })
```

# CodeBase
SE le funzioni del database contengono parametri ID allora avranno due versioni:
nomeFunzione _(privata)_:
    gli ID vengono passati come `primitive.ObjectId`
NomeFunzione _(pubblica)_:
    Si limita a convertire i parametri ID da stringa a ObjectId,
    poi utilizza `nomeFunzione` _(privata)_.

Solo internamente il DB utilizza tipi e metodi MongoDB.
Questo per permettere di cambiare il database potendo riscrivere solamente `pkg/db`,
lasciando il resto della codebase invariata.

> [!WARNING]
> Per evitare troppe richieste al database il server **NON** controlla
> l'uid del jwt token per ogni singola richiesta.
> questo significa che anche se l'account di un utente viene eliminato/bannato/sospeso
> questo potra' continuare a fare richieste.
> Anche se apparentemente e' uno schifo al momento non ci serve nessun tipo di ban/sospensione,
> e considerando la schifezza di server che abbiamo lo ritengo un trade-off che dobbiamo fare.
> Per la questione account eliminato: 
>   - Una volta eliminato viene effettuato il logout
>   - i login ovviamente diventano impossibili.
>   - un attaccante potrebbe usare account eliminati per eseguire azioni
>   - L'account eliminato rimane quindi utilizzabile per 21 giorni ( poi non potra' fare il `/auth/renew` )

