#!/usr/bin/env python3
"""
Test Account Manager for Agent Monster Game

Provides CLI tools for:
- Creating and managing test accounts
- Listing test accounts
- Generating authentication tokens
- Binding GitHub accounts
"""

import requests
import json
import argparse
import sys
from typing import Dict, Optional, List
from tabulate import tabulate
from datetime import datetime


class TestAccountManager:
    def __init__(self, api_url: str = "http://localhost:8080"):
        self.api_url = api_url
        self.base_url = f"{api_url}/api/auth"

    def create_test_account(self, username: str, password: str) -> Optional[Dict]:
        """Create a new test account"""
        try:
            response = requests.post(
                f"{self.base_url}/register-test",
                json={"username": username, "password": password},
                timeout=10,
            )
            response.raise_for_status()
            return response.json()
        except requests.exceptions.RequestException as e:
            print(f"Error creating account: {e}")
            return None

    def login(
        self, username: str, password: str, client_name: str = "cli"
    ) -> Optional[Dict]:
        """Login with test account"""
        try:
            response = requests.post(
                f"{self.base_url}/login",
                json={
                    "username": username,
                    "password": password,
                    "client_name": client_name,
                },
                timeout=10,
            )
            response.raise_for_status()
            return response.json()
        except requests.exceptions.RequestException as e:
            print(f"Error logging in: {e}")
            return None

    def login_with_token(
        self, server_token: str, client_name: str = "cli"
    ) -> Optional[Dict]:
        """Login with server token"""
        try:
            response = requests.post(
                f"{self.base_url}/login",
                json={"server_token": server_token, "client_name": client_name},
                timeout=10,
            )
            response.raise_for_status()
            return response.json()
        except requests.exceptions.RequestException as e:
            print(f"Error logging in with token: {e}")
            return None

    def list_test_accounts(self) -> Optional[List[Dict]]:
        """List all test accounts"""
        try:
            response = requests.get(f"{self.base_url}/test-accounts", timeout=10)
            response.raise_for_status()
            data = response.json()
            return data.get("accounts", []) if data.get("success") else None
        except requests.exceptions.RequestException as e:
            print(f"Error listing accounts: {e}")
            return None

    def create_github_binding_token(
        self, github_id: int, github_login: str
    ) -> Optional[Dict]:
        """Create a GitHub binding token"""
        try:
            response = requests.post(
                f"{self.base_url}/create-player-token",
                json={"github_id": github_id, "github_login": github_login},
                timeout=10,
            )
            response.raise_for_status()
            return response.json()
        except requests.exceptions.RequestException as e:
            print(f"Error creating player token: {e}")
            return None

    def bind_github_account(
        self, player_token: str, client_name: str = "cli"
    ) -> Optional[Dict]:
        """Bind GitHub account using player token"""
        try:
            response = requests.post(
                f"{self.base_url}/bind-github",
                json={"player_token": player_token, "client_name": client_name},
                timeout=10,
            )
            response.raise_for_status()
            return response.json()
        except requests.exceptions.RequestException as e:
            print(f"Error binding GitHub account: {e}")
            return None

    def validate_token(self, access_token: str) -> Optional[Dict]:
        """Validate an access token"""
        try:
            response = requests.post(
                f"{self.base_url}/validate-token",
                json={"access_token": access_token},
                timeout=10,
            )
            response.raise_for_status()
            return response.json()
        except requests.exceptions.RequestException as e:
            print(f"Error validating token: {e}")
            return None


def print_account_table(accounts: List[Dict]):
    """Pretty print accounts table"""
    if not accounts:
        print("No accounts found.")
        return

    table_data = []
    for acc in accounts:
        table_data.append(
            [
                acc.get("username", "N/A"),
                acc.get("player_id", "N/A")[:12] + "...",
                acc.get("status", "N/A"),
                acc.get("created_at", "N/A")[:10],
            ]
        )

    headers = ["Username", "Player ID", "Status", "Created Date"]
    print(tabulate(table_data, headers=headers, tablefmt="grid"))


def main():
    parser = argparse.ArgumentParser(
        description="Agent Monster Test Account Manager",
        formatter_class=argparse.RawDescriptionHelpFormatter,
        epilog="""
Examples:
  # Create a test account
  python3 test_account_manager.py create --username test_player_1 --password password123

  # Login with test account
  python3 test_account_manager.py login --username test_player_1 --password password123

  # List all test accounts
  python3 test_account_manager.py list

  # Create GitHub binding token
  python3 test_account_manager.py create-github-token --github-id 12345 --github-login myusername

  # Bind GitHub account
  python3 test_account_manager.py bind-github --player-token <token> --client-name claude

  # Validate access token
  python3 test_account_manager.py validate --token <access_token>
        """,
    )

    parser.add_argument(
        "--api-url",
        default="http://localhost:8080",
        help="API base URL (default: http://localhost:8080)",
    )

    subparsers = parser.add_subparsers(dest="command", help="Command to execute")

    # Create account command
    create_parser = subparsers.add_parser("create", help="Create a test account")
    create_parser.add_argument(
        "--username", required=True, help="Username for the test account"
    )
    create_parser.add_argument(
        "--password", required=True, help="Password for the test account"
    )

    # Login command
    login_parser = subparsers.add_parser("login", help="Login with test account")
    login_parser.add_argument("--username", required=True, help="Username")
    login_parser.add_argument("--password", required=True, help="Password")
    login_parser.add_argument("--client-name", default="cli", help="Client name")

    # Login with token command
    token_login_parser = subparsers.add_parser(
        "login-token", help="Login with server token"
    )
    token_login_parser.add_argument("--token", required=True, help="Server token")
    token_login_parser.add_argument("--client-name", default="cli", help="Client name")

    # List command
    subparsers.add_parser("list", help="List all test accounts")

    # Create GitHub token command
    github_token_parser = subparsers.add_parser(
        "create-github-token", help="Create GitHub binding token"
    )
    github_token_parser.add_argument(
        "--github-id", type=int, required=True, help="GitHub user ID"
    )
    github_token_parser.add_argument(
        "--github-login", required=True, help="GitHub username"
    )

    # Bind GitHub command
    bind_parser = subparsers.add_parser("bind-github", help="Bind GitHub account")
    bind_parser.add_argument("--player-token", required=True, help="Player token")
    bind_parser.add_argument("--client-name", default="cli", help="Client name")

    # Validate token command
    validate_parser = subparsers.add_parser("validate", help="Validate access token")
    validate_parser.add_argument(
        "--token", required=True, help="Access token to validate"
    )

    args = parser.parse_args()

    if not args.command:
        parser.print_help()
        return

    manager = TestAccountManager(api_url=args.api_url)

    if args.command == "create":
        print(f"Creating test account: {args.username}")
        result = manager.create_test_account(args.username, args.password)
        if result and result.get("success"):
            print(f"✓ Account created successfully!")
            print(f"  Player ID: {result.get('player_id')}")
            print(f"  Server Token: {result.get('server_token')}")
            print(f"  Access Token: {result.get('access_token')}")
        else:
            print(
                f"✗ Failed to create account: {result.get('error') if result else 'Unknown error'}"
            )

    elif args.command == "login":
        print(f"Logging in: {args.username}")
        result = manager.login(args.username, args.password, args.client_name)
        if result and result.get("success"):
            print(f"✓ Login successful!")
            print(f"  Player ID: {result.get('player_id')}")
            print(f"  Access Token: {result.get('access_token')}")
        else:
            print(
                f"✗ Login failed: {result.get('error') if result else 'Unknown error'}"
            )

    elif args.command == "login-token":
        print(f"Logging in with token...")
        result = manager.login_with_token(args.token, args.client_name)
        if result and result.get("success"):
            print(f"✓ Login successful!")
            print(f"  Player ID: {result.get('player_id')}")
            print(f"  Access Token: {result.get('access_token')}")
        else:
            print(
                f"✗ Login failed: {result.get('error') if result else 'Unknown error'}"
            )

    elif args.command == "list":
        print("Fetching test accounts...")
        accounts = manager.list_test_accounts()
        if accounts is not None:
            print(f"Found {len(accounts)} test accounts:")
            print()
            print_account_table(accounts)
        else:
            print("✗ Failed to list accounts")

    elif args.command == "create-github-token":
        print(f"Creating GitHub binding token for: {args.github_login}")
        result = manager.create_github_binding_token(args.github_id, args.github_login)
        if result and result.get("success"):
            print(f"✓ Player token created!")
            print(f"  Token: {result.get('player_token')}")
            print(f"  Expires at: {result.get('expires_at')}")
            print(f"  Message: {result.get('message')}")
        else:
            print(
                f"✗ Failed to create token: {result.get('error') if result else 'Unknown error'}"
            )

    elif args.command == "bind-github":
        print(f"Binding GitHub account...")
        result = manager.bind_github_account(args.player_token, args.client_name)
        if result and result.get("success"):
            print(f"✓ GitHub account bound successfully!")
            print(f"  Player ID: {result.get('player_id')}")
            print(f"  Server Token: {result.get('server_token')}")
            print(f"  Access Token: {result.get('access_token')}")
        else:
            print(
                f"✗ Failed to bind account: {result.get('error') if result else 'Unknown error'}"
            )

    elif args.command == "validate":
        print(f"Validating token...")
        result = manager.validate_token(args.token)
        if result and result.get("success"):
            print(f"✓ Token is valid!")
            print(f"  Player ID: {result.get('player_id')}")
        else:
            print(
                f"✗ Token is invalid: {result.get('error') if result else 'Unknown error'}"
            )


if __name__ == "__main__":
    main()
