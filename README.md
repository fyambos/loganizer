# 📊 GoLog Analyzer (`loganizer`)

### Description
`loganizer` est un outil en ligne de commande écrit en Go.  
Il permet d’analyser en parallèle plusieurs fichiers de logs définis dans un fichier de configuration JSON.  
Chaque log est vérifié, et un rapport est généré au format JSON avec le statut de l’analyse.

---

## 🚀 Fonctionnalités
- Lecture d’un fichier de configuration (`config.json`).
- Analyse concurrente de plusieurs fichiers de logs grâce aux **goroutines** et **WaitGroup**.
- Simulation d’un délai d’analyse (entre 50 et 200 ms).
- Gestion des erreurs (fichier introuvable, chemin invalide).
- Résultats affichés sur la console.
- Exportation des résultats dans un fichier JSON (`report.json`).

---

## 📂 Exemple de configuration

**config.json**
```json
[
  {
    "id": "web-server-1",
    "path": "test_logs/access.log",
    "type": "nginx-access"
  },
  {
    "id": "app-backend-2",
    "path": "test_logs/errors.log",
    "type": "custom-app"
  },
  {
    "id": "db-server-3",
    "path": "test_logs/mysql_error.log",
    "type": "mysql-error"
  },
  {
    "id": "invalid-path",
    "path": "/non/existent/log.log",
    "type": "generic"
  }
]
```

---

## ⚙️ Installation

```bash
# Initialiser le module Go
go mod init loganizer

# Installer Cobra
go get github.com/spf13/cobra@latest
```

---

## 🖥️ Utilisation

### Lancer une analyse

```bash
go run . analyze -c config.json -o report.json
```

### Flags disponibles
- `-c, --config` : chemin vers le fichier de configuration JSON (par défaut: `config.json`)
- `-o, --output` : chemin vers le fichier de sortie JSON (par défaut: `report.json`)

---

## 📋 Exemple d’exécution

### Commande
```bash
go run . analyze -c config.json -o ./reports/report.json
```

### Résultat console
```
Analyze lancé avec config = config.json  output = ./reports/report.json
Config chargée avec 6 entrées
[FAILED] db-server-3 (test_logs/mysql_error.log): Fichier introuvable.
  Détails erreur: GetFileAttributesEx test_logs/mysql_error.log: The system cannot find the file specified.
[FAILED] invalid-path (/non/existent/log.log): Fichier introuvable.
  Détails erreur: GetFileAttributesEx /non/existent/log.log: The system cannot find the path specified.
[OK] corrupted-log (test_logs/corrupted.log): Analyse terminée avec succès.
[OK] web-server-1 (test_logs/access.log): Analyse terminée avec succès.
[OK] app-backend-2 (test_logs/errors.log): Analyse terminée avec succès.
[OK] empty-log (test_logs/empty.log): Analyse terminée avec succès.
Rapport écrit dans ./reports/report.json
```

### Résultat JSON (`report.json`)
```json
[
  {
    "log_id": "db-server-3",
    "file_path": "test_logs/mysql_error.log",
    "status": "FAILED",
    "message": "Fichier introuvable.",
    "error_details": "GetFileAttributesEx test_logs/mysql_error.log: The system cannot find the file specified."
  },
  {
    ...
  }
]
```

---

## 🏗️ Architecture du projet

```
.
├── cmd/
│   ├── root.go       # Commande racine loganizer
│   └── analyze.go    # Commande analyze
├── reports           # Les rapports
├── config.json       # Exemple de fichier de configuration
├── test_logs/        # Logs d'exemple pour les tests
└── go.mod
```

Le dossier "internal" servira lorsque le code sera refactoré.

---

## ✅ Prochaines améliorations possibles
- [ ] Refactorer le code depuis cmd/analyzer.go vers les autres fichiers pour propreté
- [ ] Ajouter des dates et heures pour les logs/rapport
- [ ] Filtrer les résultats avec un flag `--status OK|FAILED`