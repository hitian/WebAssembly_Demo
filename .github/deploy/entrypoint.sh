#!/bin/sh
set -e
remote_repo="https://${GITHUB_TOKEN}@github.com/${GITHUB_REPOSITORY}.git" && \
remote_branch="gh-pages" && \
rm -rf gh-pages && git clone --branch $remote_branch $remote_repo gh-pages  && \
git config --global user.name "${GITHUB_ACTOR}" && \
git config --global user.email "${GITHUB_ACTOR}@users.noreply.github.com" && \
cp index.html main.wasm wasm_exec.js README.md gh-pages && \
cd gh-pages && git add . && \
git commit -m 'auto build' && \
git push $remote_repo $remote_branch:$remote_branch > /dev/null 2>&1   && \
echo 'Deploy SUCCESS!'