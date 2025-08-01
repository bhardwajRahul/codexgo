<h1 align="center">

<!-- [![README Logo](https://raw.githubusercontent.com/bastean/codexgo/main/assets/readme/logo.png)](https://github.com/bastean) -->

[![README Logo](assets/readme/logo.png)](https://github.com/bastean/codexgo)

</h1>

<div align="center">

> Example CRUD project applying Hexagonal Architecture, Domain-Driven Design (DDD), Event-Driven Architecture (EDA), Command Query Responsibility Segregation (CQRS), Behavior-Driven Development (BDD), Continuous Integration (CI), and more... in Go.

</div>

<br />

<div align="center">

[![License MIT](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/bastean/codexgo/v4)](https://goreportcard.com/report/github.com/bastean/codexgo/v4)
[![Commitizen friendly](https://img.shields.io/badge/commitizen-friendly-brightgreen.svg)](https://github.com/commitizen/cz-cli)
[![Release It!](https://img.shields.io/badge/%F0%9F%93%A6%F0%9F%9A%80-release--it-orange.svg)](https://github.com/release-it/release-it)

</div>

<div align="center">

[![Upgrade workflow](https://github.com/bastean/codexgo/actions/workflows/upgrade.yml/badge.svg)](https://github.com/bastean/codexgo/actions/workflows/upgrade.yml)
[![CI workflow](https://github.com/bastean/codexgo/actions/workflows/ci.yml/badge.svg)](https://github.com/bastean/codexgo/actions/workflows/ci.yml)
[![Release workflow](https://github.com/bastean/codexgo/actions/workflows/release.yml/badge.svg)](https://github.com/bastean/codexgo/actions/workflows/release.yml)

</div>

<div align="center">

[![Go Reference](https://pkg.go.dev/badge/github.com/bastean/codexgo/v4.svg)](https://pkg.go.dev/github.com/bastean/codexgo/v4)
[![GitHub Releases](https://img.shields.io/github/v/release/bastean/codexgo.svg)](https://github.com/bastean/codexgo/releases)

</div>

## Showcase

<div align="center">

<img src="assets/readme/desktop-home.png" />

<img src="assets/readme/desktop-dashboard.png" />

<img width="49%" src="assets/readme/mobile-home.png" />

<img width="49%" src="assets/readme/mobile-dashboard.png" />

<img src="assets/readme/mail-confirm-account.png" />

<img src="assets/readme/mail-reset-password.png" />

</div>

## CLI

### Run (Demo)

```bash
go run github.com/bastean/codexgo/v4/cmd/codexgo@latest -demo
```

> [!NOTE]
>
> - Demo version does not require any configuration, because the required **ENV** values are already preset.
>   - Here we can find the **ENV** preset values that are used in the [Demo](internal/pkg/service/env/demo.go).
> - `In-Memory` implementation will be used for EventBus, CommandBus, QueryBus and Database (`BadgerDB`).
> - Links to confirm and recover the account are sent through the `Terminal` with the following messages:
>   - _"Hi \<username\>, please confirm your account through this link: \<link\>"_.
>   - _"Hi \<username\>, please reset your password through this link: \<link\>"_.

### Installation

```bash
go install github.com/bastean/codexgo/v4/cmd/codexgo@latest
```

```bash
codexgo -h
```

```text
              _________               ________________
_____________ ______  /_____ ____  __ __  ____/__  __ \
_  ___/_  __ \_  __  / _  _ \__  |/_/ _  / __  _  / / /
/ /__  / /_/ // /_/ /  /  __/__>  <   / /_/ /  / /_/ /
\___/  \____/ \__,_/   \___/ /_/|_|   \____/   \____/

Example CRUD project applying Hexagonal Architecture, DDD, EDA, CQRS, BDD, CI, and more... in Go.

Usage: codexgo [flags]

  -demo
        Use preset ENV values
  -env string
        Path to custom ENV file
```

### Usage

> [!NOTE]
>
> - We can use the [.env.demo.cli](deployments/.env.demo.cli) file or we can create our own `.env` file where we define our own values.
>   - In the [.env.example.cli](deployments/.env.example.cli) file, we can see the values that can be used.
>     - If `CODEXGO_SMTP_*` is omitted, the links to confirm and recover the account are sent through the `Terminal` with the following messages:
>       - _"Hi \<username\>, please confirm your account through this link: \<link\>"_.
>       - _"Hi \<username\>, please reset your password through this link: \<link\>"_.
>     - We can define our own **SMTP** configuration by simply modifying the `CODEXGO_SMTP_*` variables, then we will be able to receive the links by mail.
>     - If `CODEXGO_BROKER_*` is omitted, an in-memory EventBus will be used.
>     - `In-Memory` implementation will be used for CommandBus and QueryBus.
>     - If `CODEXGO_DATABASE_*` is omitted, a `BadgerDB` in-memory database will be used.
>     - We can use a file as a database instead of memory by defining the file name in the `CODEXGO_DATABASE_BADGERDB_DSN` variable.

```bash
codexgo -env path/to/.env
```

## Docker

### Run (Demo)

> [!NOTE]
>
> - [System Requirements](#locally)
> - In the Demo version ([.env.demo](deployments/.env.demo)):
>   - `RabbitMQ` implementation will be used for the EventBus.
>     - **URL:** `http://localhost:15672`
>     - **User/Password:** `codexgo-demo`
>   - `In-Memory` implementation will be used for CommandBus and QueryBus.
>   - `MongoDB` implementation will be used as Database.
>     - **Compass:** `mongodb://codexgo-demo:codexgo-demo@localhost:27017`
>   - `Mailpit` will be used as SMTP to receive emails with the links to confirm and recover the account.
>     - **URL:** `http://localhost:8025`
>   - `codexGO` Server.
>     - **URL:** `http://localhost:8080`

```bash
task demo
```

## Features

### Project Layout

- Based on [Standard Go Project Layout](https://github.com/golang-standards/project-layout).

### Git

- Hooks managed by [husky](https://github.com/typicode/husky):
  - Pre-Push:
    - Scanning Repository for secrets using [TruffleHog CLI](https://github.com/trufflesecurity/trufflehog) and [Trivy](https://github.com/aquasecurity/trivy)
  - Pre-Commit: [lint-staged](https://github.com/lint-staged/lint-staged)
    - Scanning files for secrets using [TruffleHog CLI](https://github.com/trufflesecurity/trufflehog?tab=readme-ov-file#8-scan-individual-files-or-directories)
    - Formatting files using [GolangCI-Lint](https://golangci-lint.run/welcome/quick-start/#formatting) and [Prettier](https://prettier.io/docs/install).
  - Commit-Msg: [commitlint](https://github.com/conventional-changelog/commitlint)
    - Check [Conventional Commits](https://www.conventionalcommits.org) rules
- Commit message helper using [czg](https://github.com/Zhengqbbb/cz-git).
  - Interactive prompt that allows you to write commits following the [Conventional Commits](https://www.conventionalcommits.org) rules:
    ```bash
    task commit
    ```

### Scanners

- [TruffleHog CLI](https://github.com/trufflesecurity/trufflehog): Secrets.
- [Trivy](https://github.com/aquasecurity/trivy): Secrets, Vulnerabilities and Misconfigurations.
- [OSV-Scanner](https://github.com/google/osv-scanner): Vulnerabilities.

### Linters/Formatters

- `*.go`: [GolangCI-Lint](https://github.com/golangci/golangci-lint).
- `*.templ`: [templ fmt](https://templ.guide/developer-tools/cli#formatting-templ-files).
- `*.feature` (Gherkin): [Cucumber extension](https://cucumber.io/docs/tools/general).
- `*.*`: [Prettier cli/extension](https://prettier.io/docs/install).

### Debuggers

- `*.go`: [deadcode](https://pkg.go.dev/golang.org/x/tools/cmd/deadcode).

### Tests

- Random data generator: [Gofakeit](https://github.com/brianvoe/gofakeit).
- Unit/Integration: [Testify](https://github.com/stretchr/testify).
- Acceptance: [Testify](https://github.com/stretchr/testify), [Godog (Cucumber)](https://github.com/cucumber/godog) and [Playwright](https://github.com/playwright-community/playwright-go).

### Releases

- Automatically managed by [Release It!](https://github.com/release-it/release-it):
  - Before/After Hooks for:
    - Linting
    - Testing
  - Bump version based on [Conventional Commits](https://www.conventionalcommits.org) and [SemVer](https://semver.org/):
    - CHANGELOG generator
    - Commits and Tags generator
    - GitHub Releases

### GitHub

- Actions for:
  - Setup Languages and Dependencies ([setup](.github/actions/setup/action.yml))
- Workflows running:
  - Automatically (Triggered by **Push** or **Pull requests**) ([ci](.github/workflows/ci.yml)):
    - Secrets
    - Vulnerabilities
    - Misconfigurations
    - Linting
    - Testing
  - Manually (Using the **Actions tab** on GitHub):
    - Upgrade Dependencies ([upgrade](.github/workflows/upgrade.yml))
    - Automate Release ([release](.github/workflows/release.yml))
- Issue Templates **(Defaults)**.

### Devcontainer

- Multiple [Features](.devcontainer/devcontainer.json) already pre-configured:
  - Go
  - Task
  - Docker in Docker
  - Light-weight Desktop (Fluxbox)
  - SSH
- [Extensions](.devcontainer/devcontainer.json) and their respective [settings](.vscode/settings.json) to work with:
  - Go
  - Task
  - Docker
  - templ
  - Cucumber
    - Gherkin
  - YAML
  - TOML
  - Prettier
  - Better Comments
  - Comment Divider
  - Todo Tree
  - cSpell
  - Fomantic Snippets

### Docker

- [Dockerfile](deployments/Dockerfile)
  - **Multi-stage builds**:
    - Development
    - Testing
    - Build
    - Production
- [Compose](deployments/docker-compose.yml)
  - Switched by ENVs and Profiles.

### Broker

- Message Routing Key based on [AsyncAPI Topic Definition](https://github.com/fmvilas/topic-definition).

### Server

- [Progressive Web App (PWA)](https://developer.mozilla.org/en-US/docs/Web/Progressive_web_apps/Guides/Making_PWAs_installable)
  - [Manifest](internal/app/server/static/manifest.json)

### Security

- Server log files.
- Requests **Rate Limiting**.
- Data **authentication** via **JWT** managed by **Session Cookies**.
- [Captcha](https://github.com/mojocn/base64Captcha) implementation in forms.
- Form validation at the client using [Fomantic - Form Validation](https://fomantic-ui.com/behaviors/form.html).
  - On the server, the validations are performed using the **Value Objects** defined in the **Context**.
- Account confirmation and recovery via **Mail** or **Terminal**.
  - Limited to a certain number of attempts in a given time interval.
- Password hashing using [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt).

### Tools

- [cDeps](https://github.com/bastean/x)
  - Copies the files required by the browser dependencies from the **node_modules** folder and places them inside the **static** folder on the server.
- [syncENV](https://github.com/bastean/x)
  - Synchronize all **.env\*** files in the directory using an **.env** model.

### Scripts

- [run](deployments/run.sh)
  - Display the logs and redirect them to a file whose name depends on the time at which the service was run.
  - Used in Production Image.

## Domain > (Infrastructure | Application) > Presentation

### Bounded Context (App/Business/Department) > Modules (Troubleshooting) > Layers (Domain, Infrastructure & Application)

- **Domain (Logic Core)**
  - Value Objects (Entities)
    - Mother Creators
    - Unit Tests
  - Messages
    - Mother Creators
  - Aggregates (Sets of Entities)
    - Aggregate Root (Core Set)
    - Mother Creators
  - Role Interfaces (Ports)
    - Repository
    - Broker
  - Model Interfaces
    - Use Cases
    - Handlers/Consumers
  - Services (Abstract Logic)
  - Errors (Management)
- **Infrastructure (Port Adapters)**
  - Persistence
    - Repository Mocks
    - Implementations (Adapters)
    - Integration Tests
  - Communication
    - Broker Mocks
    - Implementations (Adapters)
    - Integration Tests
- **Application (Orchestration of Domain Logic)**
  - Use Cases
    - Implementations
  - Commands
    - Mother Creators
  - Queries/Responses
    - Mother Creators
  - Handlers/Consumers
    - Implementations
    - Unit Tests

### Services > App > (Presentation)

- **Presentation (Consumers of Bounded Context Modules)**
  - Services (Mapping)
    - Centralize Imports
    - Initializations
  - Server
    - Templates
    - Handlers
      - API
      - Views
    - Routes
      - API `/v*`
      - Views
    - Features (Gherkin)
      - Acceptance Tests
        - API
        - Views

## Workflow

### Idea

The system allows users to register a new account, log in and update their data or permanently delete their account, as well as verify and reset their account through a link sent to their email.

### Functionality

It is a monolith where CRUD operations can be performed from different presentations to the same database, this allows us to manage users from the different presentations available, in addition to having a messaging system that allows to communicate the events occurred, thus avoiding a coupling to the source of the same.

### Folders

1. `pkg/context/(modules)`
   - It is the logical core that contains all the necessary functionalities that are agnostic of any **presentation**.

2. `internal/pkg/service`
   - It is responsible for initializing all **context** functionalities so that they are ready for use, as well as for **"mapping"** certain values to centralize all imports required for **presentations** in a single place.

3. `internal/app/(presentations)`
   - These **applications** will be used as **presentations** in order to serve the functionalities to an end user.

### Idiomatic

- **Domain**
  - `errors.New*()`, `errors.BubbleUp()` & `errors.Panic()`
    - Only in the `Domain` layer and in the `*_test.go` files can we throw `errors.Panic()`.
- **Infrastructure**
  - `New*()`, `Open()` & `Close()`
    - `session`
  - `errors.New*()` & `errors.BubbleUp()`
- **Application**
  - `Run()`, `Handle()` & `On()`
  - `errors.New*()` & `errors.BubbleUp()`
- **Presentation**
  - **Modules**
    - `Start()` & `Stop()`
    - `errors.BubbleUp()`
  - **Services / Apps**
    - `Init()`, `Up()` & `Down()`
      - `log.[Wrap]()`
    - `errors.New*()` & `errors.BubbleUp()`
      - In `Apps` we will handle `Bubble Errors`.
- **Main**
  - `log.Fatal()` & `log.[Wrap]()`
    - Only `main()` can use `log.Fatal()`.
- **Logs**
  - `[embed]`
    - We use `[]` to **"embed"** external values such as error messages, fields, etc... inside our messages.
- **ENVs**
  - `os.[Getenv/LookupEnv]()`
    - Only handle `ENVs` directly in the `Presentation` layer and in the `*_test.go` files.
      - At the `Infrastructure` layer, `ENVs` are received via arguments through function parameters.
- **Blocks**
  - `const`, `var`, & `type`
    - We will group only those that are declared on a single line.

### ...`v0` > `dev0.1.0` > `ci/dev0.1.0` > `main` > `v0`...

Create `v0` branch from `main`.

```bash
task git-v0
```

Create development branch `dev0.1.0` from `v0`.

```bash
task git-dev0.1.0
```

Create branch `ci/dev0.1.0` from `dev0.1.0` to ensure that the workflows run correctly with the new changes before merging them with `main`.

```bash
task git-ci/dev0.1.0
```

Once the workflows have been successfully passed, the new changes from `ci/dev0.1.0` will be merged into `main`.

```bash
task git-main-ci/dev0.1.0
```

After releasing the new version `v0.1.0`, the `main` and `v0` branches in our local repository will be updated.

```bash
task git-pull-v0
```

To end the cycle, the `dev0.1.0` and `ci/dev0.1.0` branches will be deleted.

```bash
task git-cleanup-dev0.1.0
```

## First Steps

### Clone

#### HTTPS

```bash
git clone https://github.com/bastean/codexgo.git && cd codexgo
```

#### SSH

```bash
git clone git@github.com:bastean/codexgo.git && cd codexgo
```

### Initialize

#### Dev Container **(Recommended)**

1. System Requirements
   - [Docker](https://docs.docker.com/get-docker)
     - [Dev Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)

2. Start VS Code

   ```bash
   code .
   ```

3. Open Command Palette
   - F1

4. Run

   ```text
   Dev Containers: Reopen in Container
   ```

5. SSH **(Optional)**
   - We can connect to our `Dev Container` via `SSH` in the following ways:
     - If we have [Task](https://taskfile.dev/installation) installed on our host, being in the root of the repository

       ```bash
       task connect-2222-vscode-localhost
       ```

     - Using the SSH Client of our host

       ```bash
       ssh -p 2222 -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -o GlobalKnownHostsFile=/dev/null vscode@localhost
       ```

     - **Password:** `vscode`

6. Desktop **(Optional)**
   - We can connect to our `Dev Container` via `VNC` using our web browser.
     - **URL:** `http://localhost:6080`
     - **Password:** `vscode`
       - No password is required to connect from the **local host**.
     - Change the default applications of the context menu **(Optional)**

       ```bash
       task desktop
       ```

       - **File Manager:** `xfe`
       - **Terminal:** `terminator`
       - **Web Browser:** `chromium`

#### Locally

1. System Requirements
   - [Go](https://go.dev/doc/install)
   - [Task](https://taskfile.dev/installation)
   - [Docker](https://docs.docker.com/get-docker)

2. Run

   ```bash
   task init
   ```

### ZIP

> [!NOTE]
>
> - [System Requirements](#locally)
> - We need to change `<user>` and `<repository>` with our own values.

```bash
curl -sSfLO https://github.com/bastean/codexgo/archive/refs/heads/main.zip \
&& unzip main.zip \
&& mv codexgo-main <repository> \
&& rm main.zip \
&& cd <repository> \
&& task genesis \
&& git commit -m "feat(genesis): codexgo" \
&& git branch -M main \
&& git remote add github https://github.com/<user>/<repository>.git \
&& git push -u github main \
&& git status
```

<details>

<summary>Repository initialization from ZIP <b>(Optional)</b></summary>

## Design

#### Icons (Template: [Icon Safe Zone SVG](https://www.w3.org/TR/appmanifest/images/icon-safe-zone.svg))

- `favicon-512.png` (512x512)

#### Screenshots

- `desktop.png` (1536x860)
- `mobile.png` (425x800)

#### Readme

- `logo.png`

#### Social Preview (Template: [GitHub Repository > Settings > General > Social preview > Download template]())

- `social-preview.png`

#### Mocks

- Wireframe
- UML

## Repository

#### Genesis

- Run initialization from `ZIP` ([codexGO](https://github.com/bastean/codexgo#zip))

- Cleanup
  - `CHANGELOG`
  - Everything not required

- Add `<repository>` > `assets/`
  - Compress `assets`

- Update `#colors` > `#<repository>`

- Update (`codexgo/v*`|`codexgo`|`codexGO`) > `<repository>`
  - `Source Code`, `Files` & `Folders`

- Update `README`
  - Update `Logo`
  - Update `Description`
  - Add `Disclaimer`

    ```markdown
    > [!IMPORTANT]
    >
    > `<repository>` is still in the early stages of development.
    ```

  - Update `Tech Stack`
  - Remove all other sections, except:
    - `Logo`, `Description`, `Disclaimer`, `Badges`, `Tech Stack`, `Contributing` & `License`
  - Change `v*.*.*` > `v0.0.0`

- Commit, Upgrade & Push

  ```bash
  git commit -m "chore: update codexgo to <repository>" \
  && task upgrade \
  && git push \
  && git status
  ```

## GitHub

#### Settings

- Rulesets
  - `New branch ruleset`
    - Ruleset Name \*
      - Branches
    - Enforcement status
      - Active
    - Target branches
      - Include by pattern
        - `main`
        - `v*`
    - Branch rules
      - Restrict deletions
      - Require signed commits **(Optional)**
      - Block force pushes
  - `New tag ruleset`
    - Ruleset Name \*
      - Tags
    - Enforcement status
      - Active
    - Target tags
      - Include by pattern
        - `v*`
    - Tag rules
      - Restrict deletions
      - Require signed commits **(Optional)**
      - Block force pushes

- Secrets and variables **(Optional)**
  - Actions
    - New repository secret
      - `BOT_GPG_PASSPHRASE`
      - `BOT_GPG_PRIVATE_KEY`

        ```bash
        gpg --armor --export-secret-key [Pub_Key_ID (*-BOT)]
        ```

- Social preview
  - Upload an image…

#### Actions

- Workflows
  - Release
    - Run workflow
      - Use workflow from
        - Branch: `main`
      - Status
        - `alpha`
      - Bump
        - `minor`
    - Pre-release
      - `v0.1.0-alpha.0`

</details>

### GitHub Actions **(Optional)**

> [!IMPORTANT]
> By default, the user used to create commits and tags in workflows is `github-actions[bot]`, whose creations are not signed. If we want the creation of commits and tags to be signed, we must create the specified **"Repository Secrets"**, this way by selecting `GPG Sign` when manually launching a workflow everything will be created using the information obtained from the **"GPG Key"**.

#### Settings tab

##### Secrets and variables

- Actions
  - New repository secret
    - `BOT_GPG_PASSPHRASE`

    - `BOT_GPG_PRIVATE_KEY`

      ```bash
      gpg --armor --export-secret-key [Pub_Key_ID (*-BOT)]
      ```

### Run

#### ENVs

> [!IMPORTANT]
> If we want to modify the values inside some `.env` file, here [.env.example.demo](deployments/.env.example.demo) we can see the values that can be used.

#### Development ([.env.demo.dev](deployments/.env.demo.dev))

> [!NOTE]
>
> - `RabbitMQ`
>   - **URL:** `http://localhost:15672`
>   - **User/Password:** `codexgo-dev`
> - `MongoDB`
>   - **Compass:** `mongodb://codexgo-dev:codexgo-dev@localhost:27017`
> - `Mailpit`
>   - **URL:** `http://localhost:8025`
> - `codexGO` Server
>   - **URL:** `http://localhost:8080`
> - `codexGO` Live-Reloading Server
>   - **URL:** `http://localhost:8090`

```bash
task compose-dev
```

#### Tests

##### Unit

```bash
task test-unit
```

```bash
task test-flaky
```

##### Integration ([.env.demo.test](deployments/.env.demo.test) | [.env.demo.test.integration](deployments/.env.demo.test.integration))

```bash
task compose-test-integration
```

##### Acceptance ([.env.demo.test](deployments/.env.demo.test) | [.env.demo.test.acceptance](deployments/.env.demo.test.acceptance))

```bash
task compose-test-acceptance
```

##### Unit / Integration / Acceptance ([.env.demo.test](deployments/.env.demo.test))

```bash
task compose-tests
```

#### Production ([.env.demo.prod](deployments/.env.demo.prod))

> [!NOTE]
>
> - We must define our own **SMTP** configuration in the `CODEXGO_SMTP_*` variables, to receive the links by mail.
>   - By default, as there are no values set, the links will be sent through the terminal.
> - `RabbitMQ`
>   - **URL:** `http://localhost:15672`
>   - **User/Password:** `codexgo`
> - `MongoDB`
>   - **Compass:** `mongodb://codexgo:codexgo@localhost:27017`
> - `codexGO` Server
>   - **URL:** `http://localhost:8080`

```bash
task compose-prod
```

## Tech Stack

#### Base

- [Go](https://go.dev)
- [Gin](https://gin-gonic.com)
- [templ](https://templ.guide)
  - [Fomantic-UI](https://fomantic-ui.com)
- [RabbitMQ](https://www.rabbitmq.com/tutorials/tutorial-one-go)
- [MongoDB](https://www.mongodb.com/docs/drivers/go)
- [BadgerDB](https://docs.hypermode.com/badger)

#### Please see

- [go.mod](go.mod)
- [package.json](package.json)

## Contributing

Contributions and Feedback are always welcome!

- [Open a new issue](https://github.com/bastean/codexgo/issues/new/choose)

## License

[MIT](LICENSE)
