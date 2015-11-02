#!/bin/bash
BASE_URL="https://api.github.com/repos/$CIRCLE_PROJECT_USERNAME/$CIRCLE_PROJECT_REPONAME"

function upload_asset() {
UPLOAD_URL=$1
FILE=$2
NAME=$(basename $FILE)
  curl -s -XPOST \
    -d "@$FILE" \
    -H "Authorization: token $GITHUB_AUTH" \
    -H "Content-Type:application/octet-stream" \
    "$UPLOAD_URL?name=$NAME"
}

function release_object() {
TAG=$1
CHANGES=$2
echo "{
  \"tag_name\": \"$TAG\",
  \"name\": \"$TAG\",
  \"draft\": true,
  \"body\": \"$CHANGES\"
}"
}

function tag_changes() {
  CURRENT_TAG=$1
  PREVIOUS_TAG=$(git tag | tail -n2| head -n1)
  if [[ -z "$PREVIOUS_TAG" ]]; then PREVIOUS_TAG='master'; fi
  git log --oneline --pretty='* %s' "${PREVIOUS_TAG}..${CURRENT_TAG}"
}

function create_release_for_tag() {
TAG=$1
CHANGES=$(tag_changes "$TAG")
RELEASE_BODY=$(release_object "$TAG" "$CHANGES")
curl -i -XPOST \
  -H "Authorization: token $GITHUB_AUTH" \
  -H 'Content-Type: application/json; charset=utf-8' \
  -d "$RELEASE_BODY" \
  "$BASE_URL/releases"
}

RELEASE=$(create_release_for_tag "$CIRCLE_TAG")
UPLOAD_URL=$(echo $RELEASE | grep -Po '(?<="upload_url": ")[^{]+' | head -n1)
RELEASE_ID=$(echo $RELEASE | grep -Po '(?<="id": )[0-9]+' | head -n1)

for asset in $PWD/dist/*; do
  ASSET=$(upload_asset $UPLOAD_URL $asset)
  ASSET_ID=$(echo $ASSET | grep -Po '(?<="id":)[0-9]+' | head -n1 )
  echo "Uploaded $(basename $asset) with id $ASSET_ID in release ${RELEASE_ID}."
done
