# 🚀 Azure AKS GitOps with Argo CD

A hands-on **DevOps & GitOps** project demonstrating the deployment of a containerized Go application to **Azure Kubernetes Service (AKS)** using **Docker, Azure Container Registry (ACR), Kubernetes, GitHub, and Argo CD**.

The project demonstrates:

- 📦 Containerized application deployment
- ☁️ Azure Kubernetes Service
- 🐳 Docker multi-stage builds
- 🗄️ Azure Container Registry
- ☸️ Kubernetes deployments and services
- 🔄 GitOps-based automated synchronization
- 🔍 Configuration drift detection
- ❤️ Argo CD self-healing
- 🌐 Azure LoadBalancer-based application exposure

---

## 🏗️ Architecture

```mermaid
flowchart LR
    DEV[👨‍💻 Developer] --> GIT[🐙 GitHub Repository]

    DEV --> DOCKER[🐳 Docker Multi-Stage Build]
    DOCKER --> ACR[📦 Azure Container Registry]

    GIT --> ARGO[🔄 Argo CD]

    ARGO --> AKS[☁️ Azure Kubernetes Service]
    ACR --> AKS

    AKS --> NS[📁 azure-web Namespace]

    NS --> DEPLOY[☸️ Deployment]
    DEPLOY --> POD1[🟢 Pod 1]
    DEPLOY --> POD2[🟢 Pod 2]

    NS --> SVC[🌐 LoadBalancer Service]
    SVC --> USERS[👥 Application Users]

 ---

📋 Project Overview

This project implements an end-to-end GitOps workflow for deploying a Go web application on Azure Kubernetes Service.
