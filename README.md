# dbBackupUtility

A command-line utility built with Go to easily backup and restore MySQL and PostgreSQL databases. It supports direct command-line flags, interactive prompts, and saved configurations for streamlined workflows.

## Features

*   **Backup:** Create database backups for MySQL and PostgreSQL.
*   **Restore:** Restore databases from backup files for MySQL and PostgreSQL.
*   **Configuration Management:** Save and manage connection/backup settings for different database environments.
*   **Multiple Interaction Modes:**
    *   Use command-line flags for automation and scripting.
    *   Use interactive prompts for guided setup.
    *   Use saved configurations for frequently used settings.
*   **Logging:** Logs backup and restore actions (details like DB type, host, status, etc.).

## Installation

*(Assuming standard Go installation)*

1.  Clone the repository:
    ```bash
    git clone <your-repository-url>
    cd dbBackupUtility
    ```
2.  Build the executable:
    ```bash
    go build -o dbBackupUtility .
    ```
3.  (Optional) Move the executable to a directory in your PATH, e.g.:
    ```bash
    sudo mv dbBackupUtility /usr/local/bin/
    ```

## Usage

The utility follows a standard command-line structure:

```bash
dbBackupUtility [command] [flags]