# Finite Vault

Repository for the Finite Vault project.

A stop on the road to excellence.

Tech stack motivated by the eternal job search.

Frontend is VueJS, primary backend written in Golang. NestJS is secondary / to be done if the opportunity calls for it.

## Contributing

Just make a branch name something like `identifier-description`. For example, `mb-fix-login-bug` or `dendrobyte-page-state-update`. Please link PRs on relevant Trello tasks if you have a chance, but not totally crucial.

Let's merge all PRs into `develop` unless it needs to be deployed; even then we can create PRs for `develop` into `main` at various points. Simple "feature release" style I guess. Merges to `main` will trigger GitHub actions to deploy both backend (Golang to Digital Ocean droplet) and frontend (VueJS SPA to CloudFlare Pages).
