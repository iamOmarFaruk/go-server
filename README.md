# 🌍 Travel & Stories Platform

> *Experience the journey. Share the adventure. Built with modern Go & Tailwind CSS.*

![Project Banner](https://images.unsplash.com/photo-1469854523086-cc02fe5d8800?q=80&w=2000&auto=format&fit=crop)

## ✨ Overview

Welcome to **Travel & Stories**, a next-generation web application designed for sharing immersive travel experiences. Built with performance and aesthetics in mind, this project demonstrates a robust **Go (Golang)** backend seamlessly integrated with a modern **Vite + Tailwind CSS v4** frontend.

Whether you're exploring the interactve **Travel Guide** or reading through captivating **Stories**, this platform delivers a premium, smooth, and responsive user experience.

## 🚀 Key Features

*   **⚡ Blazing Fast Backend**: Powered by Go and the `chi` router for high-performance HTTP handling.
*   **🎨 Premium UI/UX**: Crafted with **Tailwind CSS v4**, featuring glassmorphism, smooth animations, and a curated color palette (Emerald, Stone, Sky).
*   **📱 Fully Responsive**: A mobile-first design approach ensures the site looks stunning on any device.
*   **🖼️ Dynamic Content**: Grid layouts for stories with beautiful Unsplash image integrations.
*   **🔁 Modern Dev Workflow**: Hot-reloading enabled via **Air** (Go) and **Vite** (Frontend) for an instant feedback loop.
*   **🐳 Dockerized**: Fully containerized environment for consistent development and deployment.
*   **🔒 Secure**: Custom middleware implementation for route protection (e.g., Age Verification).

## 🛠️ Tech Stack

### Backend
*   **Go (Golang)**: Core language.
*   **Chi**: Lightweight, idiomatic router.
*   **Air**: Live reload for Go development.

### Frontend
*   **HTML5 Templates**: Efficient server-side rendering.
*   **Tailwind CSS v4**: Utility-first styling with modern configuration.
*   **Vite**: Next-generation frontend tooling for asset bundling.
*   **JavaScript (ES6+)**: Interactive client-side enhancements.

### DevOps
*   **Docker & Docker Compose**: Containerization.
*   **Make**: Automation commands.

## 🏁 Getting Started

Follow these steps to get the project running on your local machine.

### Prerequisites
*   [Go](https://go.dev/) (1.21+)
*   [Node.js](https://nodejs.org/) & npm
*   [Docker](https://www.docker.com/) (Optional, for containerized run)

### 💻 Local Development

1.  **Clone the repository**
    ```bash
    git clone https://github.com/iamOmarFaruk/go-server.git
    cd go-server
    ```

2.  **Install Frontend Dependencies**
    ```bash
    npm install
    ```

3.  **Run the Application**
    
    *Option A: Using Makefile*
    ```bash
    # Terminal 1: Start Go Server
    make run
    
    # Terminal 2: Start Vite (for frontend assets)
    npm run dev
    ```

    *Option B: Docker (All-in-One)*
    ```bash
    docker compose up --build
    ```

4.  **Visit the App**
    Open [http://localhost:8080](http://localhost:8080) to view the application.

## 📂 Project Structure

```
├── cmd/                # Application entry points
├── internal/           # Private application code (Handlers, Middleware)
├── templates/          # HTML templates (Pages, Layouts, Partials)
├── static/             # Static assets (Compiled CSS/JS)
├── src/                # Source frontend assets (Tailwind, JS)
├── docker-compose.yml  # Docker configuration
└── Makefile            # Command automation
```

## 👥 Author

<div align="center">

**Omar Faruk**

[![GitHub](https://img.shields.io/badge/GitHub-iamOmarFaruk-181717?style=for-the-badge&logo=github)](https://github.com/iamOmarFaruk)
[![Website](https://img.shields.io/badge/Website-omarfaruk.dev-blueviolet?style=for-the-badge)](https://omarfaruk.dev)

*Crafting robust software with a touch of elegance.*

</div>

---

<!--
 ┌── o m a r ──┐
 │ gh@iamOmarFaruk
 │ omarfaruk.dev
 │ Created: 2025-12-21
 │ Updated: 2025-12-21
 └─ go-server ───┘
-->