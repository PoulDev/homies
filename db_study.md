# Traccia delle query al DB

## Login
1. `POST /auth/login`
-> `coll:users`

## Register
1. `POST /auth/register`
-> `coll:avatars`
-> `coll:users`

## App Homepage
> ![WARNING] DA OTTIMIZZARE
> fa davvero troppe query


1. `GET /users/me`
-> `coll:users`

2. `GET /avatars/XX`
-> `coll:avatars`

3. `GET /user/me/house`
-> `coll:houses`
-> `coll:users # get house members`


