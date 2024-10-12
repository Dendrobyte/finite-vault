# Finite Vault

Repository for the Finite Vault project, **now deprecated to pursue the bigger project of Infinite Game.** You can find the development site here: http://infinitegame.markobacon.com/ (if it isn't Finite Vault, then we've moved on to Infinite Game to production!)

It's buggy and feature lacking, but it hit my MVP. The expenses reverse when user data updates, you can't change your own daily number, expenses log every expense (aren't paginated), etc. Feel free to reach out on Twitter @Mobkinz78 if you have questions or anything.

## PREVIOUS README ##

A stop on the road to excellence.

Tech stack motivated by the eternal job search.

Frontend is VueJS, primary backend written in Golang. NestJS is secondary / to be done if the opportunity calls for it.

## Contributing

Just make a branch name something like `identifier-description`. For example, `mb-fix-login-bug` or `dendrobyte-page-state-update`. Please link PRs on relevant Trello tasks if you have a chance, but not totally crucial.

Let's merge all PRs into `develop` unless it needs to be deployed; even then we can create PRs for `develop` into `main` at various points. Simple "feature release" style I guess. Merges to `main` will trigger GitHub actions to deploy both backend (Golang to Digital Ocean droplet) and frontend (VueJS SPA to CloudFlare Pages).
