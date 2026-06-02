#!/usr/bin/env bash
set -euo pipefail

BASE_URL="${{{.BaseUrlEnvName}:-}}"
TOKEN="${{{.TokenEnvName}:-}}"
TIMEOUT="${GVA_TIMEOUT:-30}"
DATA=""
DATA_FILE=""
SHOW_HELP=0
URL_PATH="{{.Path}}"
QUERY_KEYS=({{- range .QueryParams }} "{{ .Name }}"{{- end }})
QUERY_VALUES=({{- range .QueryParams }} ""{{- end }})
HEADER_KEYS=({{- range .HeaderParams }} "{{ .Name }}"{{- end }})
HEADER_VALUES=({{- range .HeaderParams }} ""{{- end }})
{{- range .PathParams }}
{{ .VarName }}=""
{{- end }}

usage() {
  cat <<'EOF'
{{.ScriptName}} - {{.Summary}}

Usage:
  {{.ScriptName}} [options]

Options:
  --base-url <url>       Override base URL (default from {{.BaseUrlEnvName}})
  --token <token>        Override x-token (default from {{.TokenEnvName}})
  --data <json>          Raw JSON request body
  --data-file <path>     Read JSON request body from file
{{- range .PathParams }}
  --{{ .FlagName }} <value>    {{ .Description }}{{ if .Required }} (required){{ end }}
{{- end }}
{{- range .QueryParams }}
  --{{ .FlagName }} <value>    {{ .Description }}{{ if .Required }} (required){{ end }}
{{- end }}
{{- range .HeaderParams }}
  --{{ .FlagName }} <value>    {{ .Description }}{{ if .Required }} (required){{ end }}
{{- end }}
  --help                 Show this help message

Environment:
  {{.BaseUrlEnvName}}     Base server URL, e.g. http://127.0.0.1:8888
  {{.TokenEnvName}}       API token sent as x-token header
  GVA_TIMEOUT             curl timeout in seconds (default 30)

Example:
  {{.Example}}
EOF
}

if ! command -v curl >/dev/null 2>&1; then
  echo "curl is required" >&2
  exit 1
fi

while [[ $# -gt 0 ]]; do
  case "$1" in
    --base-url)
      BASE_URL="$2"
      shift 2
      ;;
    --token)
      TOKEN="$2"
      shift 2
      ;;
    --data)
      DATA="$2"
      shift 2
      ;;
    --data-file)
      DATA_FILE="$2"
      shift 2
      ;;
    --help|-h)
      SHOW_HELP=1
      shift
      ;;
{{- range .PathParams }}
    --{{ .FlagName }})
      {{ .VarName }}="$2"
      shift 2
      ;;
{{- end }}
{{- range $index, $param := .QueryParams }}
    --{{ $param.FlagName }})
      QUERY_VALUES[{{ $index }}]="$2"
      shift 2
      ;;
{{- end }}
{{- range $index, $param := .HeaderParams }}
    --{{ $param.FlagName }})
      HEADER_VALUES[{{ $index }}]="$2"
      shift 2
      ;;
{{- end }}
    *)
      echo "Unknown option: $1" >&2
      usage >&2
      exit 1
      ;;
  esac
done

if [[ "$SHOW_HELP" == "1" ]]; then
  usage
  exit 0
fi

if [[ -z "$BASE_URL" ]]; then
  echo "Missing base URL. Set {{.BaseUrlEnvName}} or pass --base-url." >&2
  exit 1
fi

if [[ -z "$TOKEN" ]]; then
  echo "Missing token. Set {{.TokenEnvName}} or pass --token." >&2
  exit 1
fi

{{- range .PathParams }}
if [[ -z "${{ .VarName }}" ]]; then
  echo "Missing required option: --{{ .FlagName }}" >&2
  exit 1
fi
{{- end }}
{{- range $index, $param := .QueryParams }}
{{- if $param.Required }}
if [[ -z "${QUERY_VALUES[{{ $index }}]}" ]]; then
  echo "Missing required option: --{{ $param.FlagName }}" >&2
  exit 1
fi
{{- end }}
{{- end }}
{{- range $index, $param := .HeaderParams }}
{{- if $param.Required }}
if [[ -z "${HEADER_VALUES[{{ $index }}]}" ]]; then
  echo "Missing required option: --{{ $param.FlagName }}" >&2
  exit 1
fi
{{- end }}
{{- end }}

if [[ -n "$DATA" && -n "$DATA_FILE" ]]; then
  echo "Use either --data or --data-file, not both." >&2
  exit 1
fi

if [[ -n "$DATA_FILE" ]]; then
  DATA="$(<"$DATA_FILE")"
fi

{{- range .PathParams }}
URL_PATH="${URL_PATH//\{\{.Name\}\}/${{ .VarName }}}"
{{- end }}

URL="${BASE_URL%/}${URL_PATH}"
QUERY_STRING=""
for i in "${!QUERY_KEYS[@]}"; do
  value="${QUERY_VALUES[$i]}"
  if [[ -n "$value" ]]; then
    if [[ -n "$QUERY_STRING" ]]; then
      QUERY_STRING+="&"
    fi
    QUERY_STRING+="${QUERY_KEYS[$i]}=$(printf '%s' "$value" | jq -sRr @uri 2>/dev/null || python - <<'PY'
import sys
from urllib.parse import quote
print(quote(sys.stdin.read().rstrip('\n'), safe=''))
PY
)"
  fi
done
if [[ -n "$QUERY_STRING" ]]; then
  URL+="?${QUERY_STRING}"
fi

CURL_ARGS=(
  --silent
  --show-error
  --fail-with-body
  --max-time "$TIMEOUT"
  -X "{{.Method}}"
  -H "x-token: ${TOKEN}"
)

for i in "${!HEADER_KEYS[@]}"; do
  value="${HEADER_VALUES[$i]}"
  if [[ -n "$value" ]]; then
    CURL_ARGS+=( -H "${HEADER_KEYS[$i]}: ${value}" )
  fi
done

if [[ -n "$DATA" ]]; then
  CURL_ARGS+=( -H "Content-Type: application/json" --data "$DATA" )
fi

curl "${CURL_ARGS[@]}" "$URL"
