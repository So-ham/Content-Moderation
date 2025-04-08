# Content-Moderation

- `POST /signup` – Register new user
- `POST /login` – Authenticate user
- `POST /comment` – Submit a new comment
- `POST /review` – Submit a new review
- Get`/posts`  - Get a set of static posts to comment / review on


## Deployment Flow

1. Build and tag Docker images for each service
2. Push containers to Docker Hub or preferred registry
3. Deploy using Kubernetes manifests
