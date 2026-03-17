# k8s-booking-app

A small **microservices booking platform** built to learn **Kubernetes,
FastAPI and containerized backend architecture**.

The goal of this project was to understand how modern backend systems
run in a Kubernetes cluster using multiple services, containerized
workloads and infrastructure tooling.

The project is intentionally structured like a real production backend
with separate services, database migrations, authentication and
Kubernetes deployment manifests.

------------------------------------------------------------------------

# Tech Stack

Backend

-   FastAPI (Python)
-   Go
-   SQLAlchemy
-   Alembic
-   JWT authentication
-   bcrypt password hashing

Infrastructure

-   Kubernetes
-   Docker
-   Skaffold
-   Docker Compose

Supporting tools

-   Mailhog (email testing)
-   Makefile for development commands

------------------------------------------------------------------------

# Architecture

The application consists of multiple backend services that communicate
within a Kubernetes cluster.

Client │ ├── accounts-service (FastAPI) │ ├── authentication │ ├── user
accounts │ └── email verification │ └── trips-service (Go) └── trip
management

Both services run as containers and are deployed through Kubernetes
manifests.

Infrastructure includes:

-   Deployments
-   Services
-   Ingress
-   Migration jobs
-   Namespaces

------------------------------------------------------------------------

# Repository Structure

k8s-booking-app/

accounts-service/ FastAPI service handling authentication and users
trips-service/ Go service responsible for trip management

k8s/ ├── accounts-service Kubernetes deployment manifests ├──
trips-service ├── ingress ├── jobs migration jobs ├── mailhog └──
namespaces

.github/ workflows/ CI workflows

docker-compose.yml Local development environment skaffold.yml Kubernetes
development workflow Makefile development commands

------------------------------------------------------------------------

# Services

## accounts-service

The accounts service is built with **FastAPI** and handles:

-   user registration
-   login
-   email verification
-   JWT token generation

The codebase follows a layered architecture.

accounts-service/src/

controllers/ services/ repositories/ models/ middleware/ dto/ utils/

Database access is handled through **SQLAlchemy** using a repository
pattern.

Password hashing is implemented with **bcrypt**.

Database migrations are managed using **Alembic**.

------------------------------------------------------------------------

## trips-service

The trips service is written in **Go** and handles trip related logic.

Structure:

trips-service/src/

auth/ config/ controllers/ database/ errors/ middleware/ models/ router/
server/ utils/

Authentication between services is handled using **JWT tokens**.

------------------------------------------------------------------------

# Running Locally

The project can be run locally using Docker Compose.

Start all services:

docker-compose up --build

This will start:

-   accounts service
-   trips service
-   database
-   mailhog (for testing email)

------------------------------------------------------------------------

# Running With Kubernetes

You can deploy the project to a Kubernetes cluster.

Apply the manifests:

kubectl apply -f k8s/

Check running pods:

kubectl get pods

Check services:

kubectl get services

------------------------------------------------------------------------

# Development Workflow

The project includes **Skaffold** for a smoother Kubernetes development
experience.

Run:

skaffold dev

Skaffold will:

-   build container images
-   deploy services to Kubernetes
-   watch for file changes
-   redeploy automatically

------------------------------------------------------------------------

# Database Migrations

Database migrations are handled using **Alembic**.

Example command:

make migrate

Migration jobs are also defined in the Kubernetes manifests.

------------------------------------------------------------------------

# Learning Goals

This project was created to learn and experiment with:

-   Kubernetes fundamentals
-   deploying microservices in Kubernetes
-   containerized backend development
-   FastAPI service architecture
-   Go microservices
-   service-to-service authentication using JWT
-   infrastructure tooling such as Skaffold and Docker

------------------------------------------------------------------------

# Possible Improvements

Future improvements could include:

-   API gateway
-   distributed tracing
-   service mesh
-   observability (Prometheus / Grafana)
-   CI/CD deployment pipelines
-   horizontal pod autoscaling

------------------------------------------------------------------------

# License

MIT
