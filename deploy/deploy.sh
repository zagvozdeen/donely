#!/usr/bin/env bash

export YC_REGISTRY=cr.yandex/crp1pfesq62cdgibjrnd
export YC_INSTANCE_ID=fhmf6br96f40r8r8fd6p
export APP_DOMAIN=donely.ru

set -euo pipefail

script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
root_dir="$(cd "$script_dir/.." && pwd)"

project="${PROJECT:-donely}"
registry="${YC_REGISTRY:-${REGISTRY:-}}"
instance_id="${YC_INSTANCE_ID:-${YC_VM_ID:-}}"
tag="${TAG:-$(git -C "$root_dir" rev-parse --short HEAD)}"
platform="${PLATFORM:-linux/amd64}"
app_domain="${APP_DOMAIN:-}"
api_env_file="${API_ENV_FILE:-/opt/donely/api.env}"

if [[ -z "$registry" ]]; then
	echo "Set YC_REGISTRY, for example: cr.yandex/<registry-id>" >&2
	exit 1
fi

if [[ -z "$instance_id" ]]; then
	echo "Set YC_INSTANCE_ID to the Container Solution VM id" >&2
	exit 1
fi

if [[ -z "$app_domain" ]]; then
	echo "Set APP_DOMAIN, for example: example.com" >&2
	exit 1
fi

for command in docker yc git; do
	if ! command -v "$command" >/dev/null 2>&1; then
		echo "Required command is missing: $command" >&2
		exit 1
	fi
done

registry="${registry%/}"
api_image="${registry}/${project}-api:${tag}"
web_image="${registry}/${project}-web:${tag}"
compose_file="$(mktemp "${TMPDIR:-/tmp}/${project}-compose.XXXXXX.yaml")"

cleanup() {
	rm -f "$compose_file"
}
trap cleanup EXIT

cd "$root_dir"

echo "Building ${api_image}"
docker build --platform "$platform" -f deploy/api.Dockerfile -t "$api_image" .

echo "Building ${web_image}"
docker build --platform "$platform" -f deploy/web.Dockerfile -t "$web_image" .

echo "Pushing images"
docker push "$api_image"
docker push "$web_image"

while IFS= read -r line || [[ -n "$line" ]]; do
	line="${line//\$\{API_IMAGE\}/$api_image}"
	line="${line//\$\{WEB_IMAGE\}/$web_image}"
	line="${line//\$\{APP_DOMAIN\}/$app_domain}"
	line="${line//\$\{API_ENV_FILE\}/$api_env_file}"
	printf '%s\n' "$line"
done < deploy/docker-compose.prod.yaml.tpl > "$compose_file"

echo "Updating Container Solution VM ${instance_id}"
yc compute instance update-container "$instance_id" --docker-compose-file="$compose_file"

echo "Deployed ${project}:${tag}"
