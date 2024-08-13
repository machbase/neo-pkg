[![ci](https://github.com/machbase/neo-pkg/actions/workflows/ci.yml/badge.svg)](https://github.com/machbase/neo-pkg/actions/workflows/ci.yml)
[![rebuild-cache](https://github.com/machbase/neo-pkg/actions/workflows/rebuild-cache.yml/badge.svg)](https://github.com/machbase/neo-pkg/actions/workflows/rebuild-cache.yml)

# machbase-neo add-on packages roster

> [!WARNING]  
> This work is currently under development and the pull requests might be rejected.

## How to register your web application

1. Ensure your package is available (see the section below)
2. Check out this repo.
3. Add your project meta information to `projects/{pkg-name}/package.yml` by referring to [projects/neo-pkg-web-example](./projects/neo-pkg-web-example) as a template.
The package names starting with `machbase` and `neo` are reserved and will be rejected or changed to other random names. Please ensure that you do not use these names for your packages. 
4. Submit a Pull Request with your changes.

## neo pkg web application

The neo pkg web application should be:

1. GitHub **Public** Repo.
2. Released and marked "latest" with semantic versioning (e.g. 1.2.3 or v1.2.3)
3. It should have `LICENSE` file which is recognizable by https://spdx.org/licenses/
4. The web doc base should be `/web/apps/{pkg-name}/`
