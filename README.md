# Projet-Forum

## Sommaire
- Présentation du projet]
- Comment installer et déployer le projet ?


------------------------------------------------------------------------------------------------------------------------------------

## Présentation du projet

  Tout d'abord parlons de notre groupe ! Il est constitué de 4 membres ! Yasser, Mathis, Nathan et Enzo.

  Maintenant parlons projet :) 

  En bref l'objectif du projet est de concevoir un forum qui permet de :
  - Créer des posts (texte et/ou images)
  - Chaque post aura une ou plusieurs   catégories (parmi une liste  pré-déterminée)
  - Laisser d'autres utilisateurs y réagir en postant des commentaires (ou réponse).
  - Ajouter des votes (Likes et dislikes)
  - Les posts et commentaires seront accessible par tout le monde, mais participer nécessitera de créer un compte utilisateur.
  - On pourra filtrer les posts par catégorie, pour ne voir que les posts qu'on a sois-même likés et postés

  Le projet se fera en plusieur grandes étapes: 
  - Conception du site (premières séances)
  - Planification (premières séances)
  - Auto-formation et ateliers (au fur et à mesure du développement)
  - Développement (le gros du temps passé)
  - Soutenance

  En plus du projet s'ajoute des ateliers pour nous aider à faire nos premiers pas sur les points suivants :
  - Conception du site web et d'applications
  - Introduction aux bases de données relationnelles, modèle de données, utilisation de SQLite et langage SQL
  - Gestion de projet et planification
  - Sensibilisation à la sécurité dans le développement web
  - Utilisation Docker et déploiement 


------------------------------------------------------------------------------------------------------------------------------------
## Comment installer et déployer le projet ?
Télecharger  <a href ="https://github.com/YasserSeryas/Projet-Forum/archive/refs/heads/main.zip">ici </a>
### Comment installer le projet ?
Décompressez le zip et l'ouvrir sur Visual Studio code 


### Comment déployer le projet ?
Sur VsCode et en etant sur la racine du projet tapez "go run main.go" depuis le terminal 

Suivez le chemin ecrit sur le terminal
------------------------------------------------------------------------------------------------------------------------------------


## La structure du projet et son architecture


   Le projet contient 4 dossiers (templates, static, src, BDD) et 2 fichiers (README.md, main.go) à la racine : 
  - Le dossier "templates" comporte tout les fichiers html, c'est donc ici que se situe toutes nos pages internet.
  - Le dossier "static" comporte les dossiers "css" (il nous permet de designer nos pages internets),"img" (pour les logo de notre site web) et "js" (pour la sidebar).
  - Le dossier "src" comporte tout nos fichier ".go" , ces fichiers ont pour utilités de gérer toute la partie back de notre projet.
  - Le dossier "BDD" comporte la base de donnée fait sur DbSchema.
  - Le fichier "README.md" comporte l'installation et le déploiment du projet ainsi que la structure du projet et son architecture.
  - Le fichier "main.go" contient 
  - Le fichier
  
  Cette organisation permet de se retrouver plus rapidement dans le projet et de retrouver n'importe qu'el fichier sans accros.
