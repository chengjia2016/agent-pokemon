#!/bin/bash

# Daily settlement script for Agent Monster Judge Server
# Runs at midnight to generate daily leaderboard and sync to GitHub

JUDGE_SERVER_URL="http://localhost:8080"

echo "=== Agent Monster Daily Settlement ==="
echo "Time: $(date)"

# Generate daily leaderboard
echo "Generating daily leaderboard..."
curl -X POST "$JUDGE_SERVER_URL/api/leaderboard/daily" \
  -H "Content-Type: application/json"

if [ $? -eq 0 ]; then
  echo "✓ Daily leaderboard generated successfully"
else
  echo "✗ Failed to generate daily leaderboard"
  exit 1
fi

echo "=== Settlement Complete ==="
