# Groupie-Trackers# Groupie-Tracker - API Pokémon TCG

## Description du projet
Ce projet est une application qui exploite l'API du [Pokémon Trading Card Game (TCG)](https://pokemontcg.io/) pour afficher et organiser des cartes Pokémon. Il permet aux utilisateurs de rechercher des cartes, d'afficher leurs détails et d'explorer différents ensembles de cartes.

## Fonctionnalités
- Recherche de cartes par nom, type et rareté
- Affichage des détails d'une carte (nom, description, image...)
- Filtrage et tri des cartes selon différents critères
- Gestion d'une collection personnelle de cartes

## Installation
### Prérequis
- [Golang](https://go.dev/) installé sur votre machine
- Accès à l'API de Pokémon TCG

### Installation et exécution
```bash
# Cloner le projet
git clone <https://github.com/LilianLPV/Groupie-Trackers.git>


# Installer les dépendances
go mod tidy

# Lancer le projet
go run main.go
```

## Routes implémentées
| Méthode | Route | Description |
|----------|-------|-------------|
| GET | `/search` | Récupère la liste des cartes |
| GET | `/selectid` | Affiche les détails d'une carte spécifique |
| GET | `/erreur` | Affiche les détails de l'erreur |
| GET | `/favorites` | Affiche les favoris |
| GET | `/aboutit` | Affiche la page a propos avec le READ.ME dedans |



## API Utilisée
Le projet exploite l'API publique [Pokémon TCG](https://pokemontcg.io/), qui propose différents endpoints permettant de récupérer des données sur les cartes et ensembles disponibles.

## FAQ - Gestion du projet

### 1. Comment avons-nous décomposé le projet ?
Le projet a été divisé en plusieurs phases :
- **Phase 1 : Analyse** - Comprendre les besoins et explorer l'API.
- **Phase 2 : Conception** - Définition de l'architecture et des routes.
- **Phase 3 : Développement** - Implémentation des fonctionnalités principales.
- **Phase 4 : Tests et finalisation** - Corrections et optimisations.

### 2. Comment j'ai réparti les tâches ?
Nous avons utilisé la méthode Agile avec des tâches définies.

### 3. Comment avons-nous géré notre temps ?
Nous avons priorisé les tâches critiques (connexion API, affichage des cartes) et avons réparti le travail en sprints hebdomadaires.

### 4. Quelle stratégie avons-nous adoptée pour nous documenter ?
Nous avons combiné la documentation officielle de l'API avec des recherches sur Golang et les bonnes pratiques en développement web.


