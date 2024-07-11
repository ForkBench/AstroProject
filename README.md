<a name="readme-top"></a>

<div align="center">
  <!-- You are encouraged to replace this logo with your own! Otherwise you can also remove it. -->
  <img src="doc/logo.png" alt="logo" width="140"  height="auto" />
  <br/>

<h3><b>Astro - A new Universe</b></h3>

</div>


# 📗 Table of Contents

- [📗 Table of Contents](#-table-of-contents)
- [📖 Astro ](#-astro-)
  - [🛠 Built With ](#-built-with-)
    - [Tech Stack ](#tech-stack-)
    - [Key Features ](#key-features-)
  - [💻 Getting Started ](#-getting-started-)
    - [Prerequisites](#prerequisites)
    - [Setup](#setup)
    - [Install](#install)
    - [Usage](#usage)
    - [Deployment](#deployment)
  - [🔭 Future Features ](#-future-features-)
  - [📝 License ](#-license-)

# 📖 Astro <a name="about-project"></a>


**Astro** is an Open-Source and free project to organize competitions and events.

## 🛠 Built With <a name="built-with"></a>

### Tech Stack <a name="tech-stack"></a>

The project is built using the following technologies:

<details>
  <summary>Back-End</summary>
  <ul>
    <li><a href="https://go.dev">Go</a></li>
  </ul>
</details>

<details>
<summary>Front-End</summary>
  <ul>
    <li><a href="https://svelte.dev">Svelte</a></li>
  </ul>
</details>

It is powered by the [Wails framework](http://wails.io).

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Key Features <a name="key-features"></a>

> 👷‍ The project is currently under development.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## 💻 Getting Started <a name="getting-started"></a>

To get a local copy up and running, follow these steps.

### Prerequisites

In order to run this project you need:

- [Go](https://golang.org/dl/)
- [Node.js](https://nodejs.org/en/download/)
- [WailsV3](https://v3alpha.wails.io/getting-started/installation/)

### Setup

Clone this repository to your desired folder:

```sh
  cd my-folder
  git clone git@gitlab.isima.fr:rovandemer/astroproject.git
```

### Install

Install this project with:

```sh
  mvn clean install
```

Then configure the environment variables:

```sh
  cp .env.template .env
```

### Usage

To run the project, execute the following command:

```sh
  docker-compose up
```

### Deployment

You can deploy this project using:

```sh
  mvn clean install
  docker-compose up
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## 🔭 Future Features <a name="future-features"></a>

> Describe 1 - 3 features you will add to the project.

- [ ] **Web UI based on SSE (Server-Sent Events)**
- [ ] **Login with 3rd party providers**

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## 📝 License <a name="license"></a>

This project is [Apache](LICENSE) licensed.

<p align="right">(<a href="#readme-top">back to top</a>)</p>