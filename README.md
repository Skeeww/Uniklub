# Uniklub

![Go](https://img.shields.io/badge/go-black?style=for-the-badge&logo=go)
![Postgresql](https://img.shields.io/badge/postgresql-black?style=for-the-badge&logo=postgresql)
![Gin](https://img.shields.io/badge/gin-black?style=for-the-badge&logo=gin)
![JWT](https://img.shields.io/badge/jwt-black?style=for-the-badge&logo=json-web-tokens)

Uniklub est une application de gestion de clubs permettant aux utilisateurs de gérer des clubs, des utilisateurs et des inventaires associés. Elle fournit une API RESTful sécurisée pour effectuer diverses opérations.

## Fonctionnalités

- Gestion des utilisateurs : création, mise à jour et suppression.
- Gestion des clubs : ajout, mise à jour, récupération et suppression.
- Authentification JWT pour sécuriser les endpoints.
- Gestion des éléments et des catégories associés aux clubs.

## Prérequis

- **Go** version 1.24 ou supérieure.
- **PostgreSQL** pour la base de données.

## Installation

1. Clonez le dépôt :

   ```bash
   git clone <url-du-dépôt>
   cd Uniklub
   ```

2. Installez les dépendances :

   ```bash
   go mod tidy
   ```

3. Configurez la base de données PostgreSQL :

   - Créez une base de données nommée `app`.
   - Mettez à jour les informations de connexion dans `SetupDatabase` dans `setup.go` si nécessaire.

4. Modifiez le `Makefile` pour configurer les migrations de base de données :

   - Remplacez les informations de connexion par celles correspondant à votre environnement PostgreSQL dans les cibles `migrate-up` et `migrate-down`.

5. Appliquez les migrations :

   ```bash
   make migrate-up
   ```

## Lancer le serveur

1. Compilez le projet :

   ```bash
   make main
   ```

2. Lancez le serveur :

   ```bash
   ./build/uniklub-server
   ```

Le serveur sera accessible par défaut sur `http://localhost:8080`.

## Endpoints principaux

### Authentification

- **POST** `/v1/auth/` : Connexion utilisateur.
- **GET** `/v1/auth/me` : Récupérer les informations de l'utilisateur connecté.

### Utilisateurs

- **POST** `/v1/users/` : Créer un utilisateur.

### Clubs

- **GET** `/v1/clubs/` : Récupérer tous les clubs.
- **GET** `/v1/clubs/:id` : Récupérer un club par ID.
- **POST** `/v1/clubs/` : Ajouter un club.
- **PUT** `/v1/clubs/:id` : Mettre à jour un club.
- **DELETE** `/v1/clubs/:id` : Supprimer un club.

## Tests

Pour exécuter les tests :

```bash
go test -v ./...
```

## Contribution

1. Forkez le dépôt.
2. Créez une branche pour vos modifications :

   ```bash
   git checkout -b feature/nom-de-la-fonctionnalité
   ```

3. Faites vos modifications et committez-les :

   ```bash
   git commit -m "Description des modifications"
   ```

4. Poussez vos modifications :

   ```bash
   git push origin feature/nom-de-la-fonctionnalité
   ```

5. Ouvrez une Pull Request.

## Licence

![GitHub License](https://img.shields.io/github/license/Skeeww/Uniklub?style=for-the-badge&logo=github)

Ce projet est sous licence GNU General Public License v3.0 or later. Consultez le fichier `LICENSE` pour plus d'informations.
