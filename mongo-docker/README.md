# MongoDB

## Installation

- Pour installer Mongo DB, lancer la suite de commandes suivante
  ```sh
  sudo apt update
  sudo apt-get install mongodb
  ```
> La commande précédente a également installé `mongo` la CLI de Mongo DB, que nous allons donc pouvoir utiliser par la suite

- Pour lancer le container `Mongo DB` e diriger dans le dossier mongo-docker, et lancer la commande
```sh
docker-compose up -d
```

- Afin de rentrer en mode interactif avec la DB, lancer la commande :
```sh
mongo --host mongodb://root:example@localhost:27018
```

- Commandes utiles de la CLI mongo :
```bash
show databases # Show all databases inside the motor
use <db> # Connect to the specified <db>
show collections # Show all collections inside the used <db>
db.<collection>.find() # Find all elements inside the specified <collection>
```