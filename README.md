
# sucellus

[![go](https://github.com/joaomarcuslf/sucellus/actions/workflows/go.yml/badge.svg)](https://github.com/joaomarcuslf/sucellus/actions/workflows/go.yml)

## Motivation

I use my RaspberryPi to run some of my personal projects, like a cloud platform, I wanted to make it easier to add, stop and update those repos, this is Sucellus, it is my personal cloud platform, it must have this functionalities:

- Add a project
  - Providing Name
  - Providing URL
  - Providing Language
  - Providing Port
  - Providing Env Vars
  - Providing Pooling interval
- Start project (auto start on boot)
- Stop project (auto stop when updating project)
- Edit Project

## Getting Started

1. Copy ```sample.env``` to ```.env``` and rename the variables if you need
2. You can run this repo on vscode

![image](https://raw.githubusercontent.com/joaomarcuslf/sucellus/main/static/run-application.png)

Or, you can run by doing this:

```sh
make build-mongo # Only once

make start-mongo

make dev
```

You must always start mongo while in dev.

## Running Tests

```sh
make test
```

## Collaborators

<table>
  <tr>
    <td align="center">
      <a href="https://github.com/joaomarcuslf">
        <img src="https://avatars.githubusercontent.com/u/53450523?v=4" width="100px;" alt="Joaomarcuslf's Github picture"/><br>
        <sub>
          <b>joaomarcuslf</b>
        </sub>
      </a>
    </td>
  </tr>
</table>

[⬆ Scroll top](#sucellus)<br>
