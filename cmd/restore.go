/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/vaibhavaaditya/dbBackupUtility/pkg/restore"
	"github.com/spf13/cobra"
)

var input string

func handleRestoreResult(err error) {
    if err != nil {
        fmt.Println("Restore failed:", err)
    } else {
        fmt.Println("Restore completed successfully.")
    }
}

// restoreCmd represents the restore command
var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore database from backup",
	Long:  `Restores the database using a backup file provided by the user.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch dbType {
        case "mysql":
            err := restore.RestoreMYSQL(host, port, user, password, dbName, input)
            handleRestoreResult(err)
        case "postgres":
            // err := restore.RestorePostgres(host, port, user, password, dbName, input)
            // handleRestoreResult(err)
        default:
            fmt.Printf("Unsupported database type: %s\n", dbType)
        }
	},
}

func init() {
    rootCmd.AddCommand(restoreCmd)

    restoreCmd.Flags().StringVar(&dbType, "type", "", "Database type (mysql)")
    restoreCmd.Flags().StringVar(&host, "host", "localhost", "Database host")
    restoreCmd.Flags().IntVar(&port, "port", 3306, "Database port")
    restoreCmd.Flags().StringVar(&user, "user", "", "Database username")
    restoreCmd.Flags().StringVar(&password, "password", "", "Database password")
    restoreCmd.Flags().StringVar(&dbName, "database", "", "Database name")
    restoreCmd.Flags().StringVar(&input, "input", "", "Input backup file path")


    restoreCmd.MarkFlagRequired("type")
    restoreCmd.MarkFlagRequired("user")
    restoreCmd.MarkFlagRequired("password")
    restoreCmd.MarkFlagRequired("database")
    restoreCmd.MarkFlagRequired("input")
}
