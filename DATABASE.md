> [!WARNING]
> Per evitare troppe richieste al database il server **NON** controlla
> che l'uid del jwt token sia valido per ogni singola richiesta.
> questo significa che anche se l'account di un utente viene eliminato/bannato/sospeso potra' comunque continuare a fare richieste.
> Anche se e' uno schifo al momento non ci serve nessun tipo di ban/sospensione,
> e considerando la schifezza di server che abbiamo lo ritengo un trade-off che dobbiamo accettare.
>
> Alcune note sull'account eliminato/bannato/sospeso che puo' continuare ad essere utilizzato: 
>   - in caso di eliminazione account viene effettuato il logout.
>   - i login ovviamente diventano impossibili.
>   - L'account sarebbe quindi utilizzabile per 21 giorni, fino a quando non sara' necessario rinnovare il token, cosa che ovviamente l'API blocchera'.


