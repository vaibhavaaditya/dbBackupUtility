/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/vaibhavaaditya/dbBackupUtility/pkg/backup"
	"github.com/spf13/cobra"
)

var (
    dbType, host, user, password, dbName, output string
    port                                         int
)

func handleBackupResult(err error) {
    if err != nil {
        fmt.Println("Backup failed:", err)
    } else {
        fmt.Println("Backup completed successfully.")
    }
}


// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup the database",
	Long:  `Creates a backup of the configured database and saves it to the specified location.`,
	Run: func(cmd *cobra.Command, args []string) {
        switch dbType {
        case "mysql":
            err := backup.BackupMySQL(host, port, user, password, dbName, output)
            handleBackupResult(err)
        case "postgres":
            // err := backup.BackupPostgres(host, port, user, password, dbName, output)
            // handleBackupResult(err)
        default:
            fmt.Printf("Unsupported database type: %s\n", dbType)
        }
    },
}

func init() {
	rootCmd.AddCommand(backupCmd)

    backupCmd.Flags().StringVar(&dbType, "type", "", "Database type (mysql)")
    backupCmd.Flags().StringVar(&host, "host", "localhost", "Database host")
    backupCmd.Flags().IntVar(&port, "port", 3306, "Database port")
    backupCmd.Flags().StringVar(&user, "user", "", "Database username")
    backupCmd.Flags().StringVar(&password, "password", "", "Database password")
    backupCmd.Flags().StringVar(&dbName, "database", "", "Database name")
    backupCmd.Flags().StringVar(&output, "output", "backup.sql", "Output file path for backup")


    backupCmd.MarkFlagRequired("type")
    backupCmd.MarkFlagRequired("user")
    backupCmd.MarkFlagRequired("password")
    backupCmd.MarkFlagRequired("database")
}