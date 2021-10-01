# Documentations

## Spécifique à GO
- Basic API : https://www.youtube.com/watch?v=W5b64DXeP0o&t=27s&ab_channel=TutorialEdge
- Mux documentation : https://www.gorillatoolkit.org/pkg/mux
- Mux examples : https://github.com/gorilla/mux#examples

## Spécifique à MongoDB
- Installation de mongoDB : https://dev.to/seanwelshbrown/installing-mongodb-on-windows-subsystem-for-linux-wsl-2-19m9
- ODM pour MongoDB : https://github.com/Kamva/mgm
- Première interaction avec MongoDB : https://www.loginradius.com/blog/async/mongodb-as-datasource-in-golang/

# Commandes utiles

Get une nouvelle bibliotheque :
```go 
go get 
``` 

Run le back : 
```go
go run . 
``` 

Permet de creer un module : 
```go
go mod init 
``` 

Permet de telecharger ce qu'il y a dans le module :
```go
go mod  
``` 

Refresh le `.mod` :
```go
go mod tidy  
```

# Autres

- Architecture de fichiers : https://github.com/bxcodec/go-clean-arch