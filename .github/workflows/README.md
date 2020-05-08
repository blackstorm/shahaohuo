## on push
```yaml
on: [push]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      # https://github.community/t5/GitHub-Actions/get-action-job-id/td-p/43276#
      # https://help.github.com/en/actions/reference/context-and-expression-syntax-for-github-actions#github-context
      - uses: actions/checkout@v2
      - name: package image
        run: |
          docker build . --file ./build/server/Dockerfile --build-arg APP_ENV=dev --tag docker.pkg.github.com/$GITHUB_REPOSITORY/server:$GITHUB_RUN_ID
```