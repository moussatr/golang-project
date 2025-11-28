# Flight Aggregator

L'objectif de l'API est de récupérer tous les vols vers une destination et de les trier par :

- prix
- date de départ
- temps de voyage

# Standardisation des données

## Problème initial

Les données provenant de `j-server1` et `j-server2` avaient des structures différentes :

- **j-server1** : Structure simple et plate avec des champs directs (`bookingId`, `passengerName`, `price`, etc.)
- **j-server2** : Structure imbriquée avec des objets (`reference`, `traveler.firstName/lastName`, `segments[]`, `total.amount`, etc.)

## Solution implémentée

### 1. Modèle standardisé (`server/model/flight.go`)

Un modèle unique `Flight` a été créé pour représenter les vols des deux sources. Ce modèle correspond à la structure de `j-server1` pour maintenir la compatibilité :

```go
type Flight struct {
    BokingId         string    `json:"bookingId"`
    Status           string    `json:"status"`
    PassengerName    string    `json:"passengerName"`
    FlightNumber     string    `json:"flightNumber"`
    DepartureAirport string    `json:"departureAirport"`
    ArrivalAirport   string    `json:"arrivalAirport"`
    DepartureTime    time.Time `json:"departureTime"`
    ArrivalTime      time.Time `json:"arrivalTime"`
    Price            float64   `json:"price"`
    Currency         string    `json:"currency"`
}
```

### 2. Modèle pour j-server2 (`server/model/flight_server2.go`)

Un modèle dédié `FlightServer2` a été créé pour représenter la structure brute des données de `j-server2` :

```go
type FlightServer2 struct {
    Reference string
    Traveler  TravelerServer2
    Segments  []SegmentServer2
    Total     TotalServer2
    Status    string
}
```

### 3. Fonction de transformation (`server/utils/transformer.go`)

Une fonction `TransformServer2ToFlight` convertit les données de `j-server2` vers le format standardisé :

- **`reference` → `bookingId`** : Le champ `reference` de j-server2 est mappé vers `bookingId`
- **`traveler.firstName + lastName` → `passengerName`** : Les noms sont combinés en une seule chaîne
- **`segments[0]` → départ** : Le premier segment fournit l'aéroport et l'heure de départ
- **`segments[-1]` → arrivée** : Le dernier segment fournit l'aéroport et l'heure d'arrivée
- **`total.amount` → `price`** : Le montant total devient le prix
- **`segments[0].flight.number` → `flightNumber`** : Le numéro de vol du premier segment est utilisé

### 4. Repository mis à jour (`server/repo/server2_repo.go`)

Le `Server2Repo` :

1. Récupère les données brutes de j-server2 au format `FlightServer2`
2. Transforme chaque vol via `TransformServer2ToFlight`
3. Retourne une liste de `Flight` standardisés

## Exécution des tests

### Prérequis

Assurez-vous que `gotestsum` est installé et accessible dans votre `PATH` :

```bash
go install gotest.tools/gotestsum@latest
```

Ajoutez `$HOME/go/bin` à votre `PATH` si nécessaire (dans `~/.zshrc` ou `~/.bashrc`) :

```bash
export PATH=$PATH:$HOME/go/bin
```

### Exécuter les tests

Depuis le répertoire `server/` :

**Option 1 : Avec gotestsum (mode watch - recommandé pour le développement)**

```bash
make test
```

Cette commande lance `gotestsum` en mode watch, qui surveille les fichiers et relance automatiquement les tests à chaque modification.

**Option 2 : Tests uniques avec go test**

```bash
go test ./...
```

**Option 3 : Tests avec affichage détaillé**

```bash
go test -v ./...
```

### Structure des tests

Les tests couvrent :

- **Tri par prix** : Vérification que les vols sont triés par prix croissant
- **Tri par date de départ** : Vérification que les vols sont triés par heure de départ
- **Tri par défaut** : Vérification que le tri par prix est utilisé par défaut
- **Gestion des erreurs** : Vérification que les erreurs des repositories sont propagées correctement

Les tests utilisent des mocks pour isoler la logique métier des appels HTTP réels.
### Membres du groupe: A.M. Ornelle Chougourou , Moussa Traore, Ryan MAMBOU , Yacine Boucenna
