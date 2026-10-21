# AWS Elastic Beanstalk deployment notes

This project is prepared for the Elastic Beanstalk Go platform (Amazon Linux 2023), without Docker.

## Files added/updated
- `.env` is intentionally excluded from this deployment bundle; configure database settings in Elastic Beanstalk environment properties.
- `Buildfile`: compiles the repository-root Go application to `bin/application`.
- `Procfile`: starts `bin/application` as the web process.
- `config/config.go`: supports `SERVER_PORT`, then `PORT`, then default `5000`; supports `DB_SSLMODE` (default `require`).

## Before pushing to GitHub
1. Remove `.env` from Git tracking and do not commit real database credentials:
   `git rm --cached .env` (only if it is already tracked).
2. Ensure `.env` is ignored by `.gitignore`.
3. If `.env` was previously pushed with real credentials, rotate the database password.
4. Commit and push the deployment files.

## Elastic Beanstalk environment properties
Configure these in Elastic Beanstalk > Environment > Configuration > Updates, monitoring, and logging / Software (the exact panel label may vary by platform):

- `DB_HOST`: your existing RDS endpoint hostname (no `https://`, no port suffix)
- `DB_PORT`: `5432`
- `DB_USER`: your database username
- `DB_PASSWORD`: your database password (prefer Secrets Manager where supported)
- `DB_NAME`: your database name
- `DB_SSLMODE`: `require`
- `SERVER_PORT`: `5000`

Do not put actual secrets in this README or source control.

## Deploy using CodePipeline
Create a CodePipeline with GitHub as the source (main branch) and Elastic Beanstalk as the deploy provider. Select the Elastic Beanstalk application and environment. The source artifact is deployed as an Elastic Beanstalk source bundle.

## Database/networking
- Keep RDS private.
- The RDS security group must allow inbound TCP 5432 from the security group used by the Elastic Beanstalk instances.
- Do not use `localhost` as `DB_HOST` in AWS.
- Ensure the Elastic Beanstalk environment has network access to the RDS endpoint.
- `main.go` runs `AutoMigrate`; review schema changes and use controlled migrations for production.

## Endpoints
- REST: `/employees`, `/employees/:id`
- GraphQL: `/graphql`
- GraphQL UI: `/playground`
