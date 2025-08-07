# golang-lessons

## Repository Setup Setup

- [Containerfile](./.devcontainer/Containerfile)
- [Devcontainer](./.devcontainer/devcontainer.json)
- [Tasks](./.vscode/tasks.json)
- [Launch](./.vscode/launch.json)

## Application Resources

- **[config.default.yaml](config.default.yaml)**: File di configurazione dell'applicazione di default con il quale configura tutti i setup
- **Build Flags**:
  - -ldflags
  - -X main.AppName=`golang-lessons `
  - -X main.Version=`$(git describe --tags --abbrev=0 2>/dev/null || echo undefined)`
  - -X main.Commit=`$(git rev-parse --short HEAD)`
  - -X main.BuildTime=`$(date -u +%Y-%m-%dT%H:%M:%SZ)`
- **Program Flags**:
  - `--config-file-name`: Nome del file di configurazione (es: `config.docker`)
  - `--config-file-type`: Tipo di file di configurazione (es: `yaml`)
  - `--config-file-path`: Directory del file di configurazione (es: `/tmp/config`)
- **bin**: Cartella di output del `binary` generato

## Application Tree

- **main**: package principale dal quale l'applicazione viene avviata
- **util**: package per gestire `function`, `method`, `const` condivisi tra tutti i package dell'applicazione
- **logging**: package per la configurazione del logging `logrus`
- **config**: package per la configurazione dell'applicazione mediante `viper`
- **api**: package per esporre le api tramite `go-gin`