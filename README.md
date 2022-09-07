# go-cleancodegen - generate a scaffold project

### Subjects

- Packaging and Continuous Integration
- Main Application
- Private Codes
- Controllers
- Database
- Database models
- Repository Implementation
- Application Service Implementation
- Templates (not in list above)
- Application service iterface
- UseCase
- Domain
- Entity
- Repository Interface
- ValueObject
- Tests

### Clean code folders and files
```zsh
build
cmd
internal/app/adapter/postgres/model
internal/app/adapter/mysql/model
internal/app/adapter/repository
internal/app/adapter/service
internal/app/application
internal/app/application/usecase
internal/app/application/service
internal/app/domain/model
internal/app/domain/repository
internal/app/domain/valueobject
pkg/utils
tests
LICENSE
README.md
main.go
.gitignore
```

###  Domain Layer

- The core of Clean Architecture. It says "Entities".

###  Application Layer

- The second layer from the core. It says "Use Cases".

###  Adapter Layer

- The third layer from the core. It says "Controllers / Gateways / Presenters".

###  External Layer

- The fourth layer from the core. It says "Devices / DB / External Interfaces / UI / Web".
  - **We DON'T write much codes in this layer.**

<p align="center">
  <img src="https://user-images.githubusercontent.com/19743841/93830264-afa9c480-fcaa-11ea-9589-7c5308c291f4.jpg">
</p>
<p align="center">
  <a href="https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html">The Clean Architecture</a>
</p>

## Versioning and license

Our version numbers follow the [semantic versioning specification](http://semver.org/). You can see the available versions by checking the [tags on this repository](https://github.com/thiagozs/go-cleancodegen/tags). For more details about our license model, please take a look at the [LICENSE](LICENSE) file.

**2022**, thiagozs.
