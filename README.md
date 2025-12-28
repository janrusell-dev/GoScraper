<a id="readme-top"></a>

<details> <summary>Table of Contents</summary> <ol> <li> <a href="#about-the-project">About The Project</a> <ul> <li><a href="#built-with">Built With</a></li> </ul> </li> <li> <a href="#getting-started">Getting Started</a> <ul> <li><a href="#prerequisites">Prerequisites</a></li> <li><a href="#installation">Installation</a></li> </ul> </li> <li><a href="#usage">Usage</a></li></ol> </details>

<!-- ABOUT THE PROJECT -->

## About The Project

GoScraper is a modular full-stack utility designed for automated data extraction and local storage. It consists of a high-concurrency Go service that handles the heavy lifting of crawling and parsing, paired with a Next.js interface for controlling scrape parameters and handling browser-side data exports.

The project was built to experiment with Go's colly framework and to implement a clean bridge between a Go-based API and a TypeScript-based frontend.

Core Technical Specs:
The Backend: A Go REST API using colly for asynchronous scraping, sync.Mutex for thread-safe data collection, and built-in caching to prevent redundant network requests.

The Frontend: A Next.js (App Router) client-side dashboard that manages scraper state, pagination, and category selection.

Data Layer: Implements a strict "contract" between Go Structs and TypeScript Interfaces to ensure consistent data structures.

Export Engine: A custom utility layer in the frontend that transforms raw JSON responses into downloadable .json and .csv blobs without requiring server-side storage.

Why this structure?
Most scrapers are either simple scripts or complex apps with heavy databases. GoScraper sits in the middle:

Go provides the speed and concurrent visitors.

Next.js provides an immediate way to visualize the data before saving.

Client-Side Export keeps the backend stateless and fast—no need to manage file systems or cloud storage for the results.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Built With

[![Next][Next.js]][Next-url] [![Go][Go-badge]][Go-url] [![Tailwind][Tailwind-badge]][Tailwind-url] [![TypeScript][TS-badge]][TS-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Getting Started

To get a local copy up and running, follow these simple steps. Since this project is a monorepo (or split between frontend/backend), you will need both the Go runtime and Node.js environment.

### Prerequisites

- Go (Version 1.20 or higher) Download Go

- Node.js & npm (LTS Version) Download Node

### Installation

1. Clone the repo

```sh
git clone  git clone https://github.com/janrusell-dev/goscraper.git
```

2. Backend Setup

```sh
go mod tidy
go run cmd/main.go
```

The server should now be running at http://localhost:8080

3. Frontend Setup
   Install NPM packages

```sh
cd frontend
npm install
```

4. Env Configuration Create a .env file in /frontend

```sh
NEXT_PUBLIC_API_URL=http://yourapihost
```

5. Start frontend

```sh
npm run dev
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- USAGE EXAMPLES -->

## Usage

1. Select Category: Use the dropdown to filter specific data types.

2. Set Page: Define the pagination depth for the scraper.

3. Run: Click RUN SCRAPER to see real-time JSON results in the preview window.

4. Export: Click DOWNLOAD .JSON or DOWNLOAD .CSV to save the results to your local machine.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[Next.js]: https://img.shields.io/badge/next.js-000000?style=for-the-badge&logo=nextdotjs&logoColor=white
[Next-url]: https://nextjs.org/
[Go-badge]: https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white
[Go-url]: https://golang.org/
[Tailwind-badge]: https://img.shields.io/badge/Tailwind_CSS-38B2AC?style=for-the-badge&logo=tailwind-css&logoColor=white
[Tailwind-url]: https://tailwindcss.com/
[TS-badge]: https://img.shields.io/badge/TypeScript-007ACC?style=for-the-badge&logo=typescript&logoColor=white
[TS-url]: https://www.typescriptlang.org/
