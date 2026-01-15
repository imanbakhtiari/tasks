# Cloud‑Native Infrastructure (Kubernetes + Crossplane + LocalStack)

Build a local, event‑driven pipeline on Kubernetes using AWS services inside LocalStack, provisioned via Crossplane:
- Producer -> SNS topic -> SQS queue -> Consumer -> DynamoDB table
- Simple UI (DynamoDB Admin) to view stored items

## What you need
- Region: eu‑central‑1
- LocalStack
- Crossplane
- AWS resources (provisioned via Crossplane):
  - SNS topic: `justtrack-dev-devops-producer-events`
  - SQS queue: `justtrack-dev-devops-consumer-events`
  - Subscription: queue subscribed to the topic
  - DynamoDB table: `justtrack-dev-devops-consumer-events` with hash key `Id` (string)
- Apps (deployed with Helm):
  - Producer: `ghcr.io/justtrackio/devopstest-producer:latest`
    - Environment variable `CLOUD_AWS_DEFAULTS_ENDPOINT` with a valid endpoint to connect to LocalStack
  - Consumer: `ghcr.io/justtrackio/devopstest-consumer:latest`
    - Environment variable `CLOUD_AWS_DEFAULTS_ENDPOINT` with a valid endpoint to connect to LocalStack
  - DynamoDB Admin UI

## Success criteria
- Producer, Consumer, and DynamoDB Admin pods are Running
- SNS topic, SQS queue, subscription, and DynamoDB table exist with the exact names above
- At least one item is written to the DynamoDB table `justtrack-dev-devops-consumer-events`
- DynamoDB Admin shows the table and items

## Submission (keep it simple and complete)
Provide a single folder that contains:
- Screenshot: DynamoDB Admin showing the table `justtrack-dev-devops-consumer-events`
- Code: All manifests, configs, and scripts you wrote
- Docs: A README with documentation (setup, decisions, issues, improvements, etc.)

## Dependencies
- Docker Desktop 4.40+ with Kubernetes enabled
- LocalStack Docker Desktop extension
- kubectl and Helm

Note: If you see `Failed to inspect image ... unable to convert a nil pointer to a runtime API image` when using Docker Desktop, disable Docker Desktop → General → “Use containerd for pulling and storing images”.

## Evaluation (what we look for)
- Kubernetes, Crossplane, and Helm proficiency
- Working end‑to‑end data flow and clear, minimal documentation
- Clean, reproducible setup

