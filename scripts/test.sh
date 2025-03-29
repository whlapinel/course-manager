NAME="$1"

if [ -z "$NAME" ]; then
    echo "❌ Usage: $0 <path-to-backup.tar.gz>"
    exit 1
fi

echo "you typed " "$NAME" 