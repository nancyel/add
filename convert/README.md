### Workflow

![workflow]("workflow.png")

### How to get started

1.  Clone this repo.

2.  Add an auth token to your .npmrc file

```
npm config set @nancyel:registry=https://npm.pkg.github.com
npm config set -- '//npm.pkg.github.com/:_authToken=' "<GITHUB_TOKEN>"
```

3.  Run `yarn add`

4.  Run `yarn convert`

5.  Check the root directory for the generated files (e.g. `lines_{time}.json`)
