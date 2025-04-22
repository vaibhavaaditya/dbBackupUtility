/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"
	"github.com/vaibhavaaditya/dbBackupUtility/pkg/backup"
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var (
    dbType, host, user, password, dbName, output, savedConfig string
    port                                         			  int
)

func handleBackupResult(err error) {
    if err != nil {
        fmt.Println("Backup failed:", err)
    } else {
        fmt.Println("Backup completed successfully.")
    }
}

func dispatchBackup(dbType, host string, port int, user, password, dbName, output string) {
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
}


// backupCmd represents the backup command
var backupCmd = &cobra.Command{
	Use:   "backup",
	Short: "Backup the database",
	Long:  `Creates a backup of the configured database and saves it to the specified location.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Mode 3: Saved Config
		if savedConfig != "" {
			// Fetch config logic (mock for now)
			config := map[string]string{
				"dbType": "mysql",
				"host":   "localhost",
				"port":   "3306",
				"user":   "root",
				"dbName": "mydb",
				"output": "backup_from_saved.sql",
			}
		
		
			// Prompt for password only
			survey.AskOne(&survey.Password{Message: "Enter database password:"}, &password, survey.WithValidator(survey.Required))
		
		
			portInt, _ := strconv.Atoi(config["port"])
		
		
			dispatchBackup(config["dbType"], config["host"], portInt, config["user"], password, config["dbName"], config["output"])
			return
		}


		// Mode 1: If required flags are set (non-zero), use them directly
		flagUsed := cmd.Flags().Changed("type") || cmd.Flags().Changed("host") || cmd.Flags().Changed("user") ||
			cmd.Flags().Changed("password") || cmd.Flags().Changed("database")


		if flagUsed {
			dispatchBackup(dbType, host, port, user, password, dbName, output)
			return
		}


		// Mode 2: Prompt interactively
		var qs = []*survey.Question{
			{
				Name: "dbType",
				Prompt: &survey.Select{
					Message: "Choose database type:",
					Options: []string{"mysql", "postgres"},
				},
				Validate: survey.Required,
			},
			{
				Name:     "host",
				Prompt:   &survey.Input{Message: "Database host:", Default: "localhost"},
				Validate: survey.Required,
			},
			{
				Name:     "port",
				Prompt:   &survey.Input{Message: "Database port:", Default: "3306"},
				Validate: survey.Required,
			},
			{
				Name:     "user",
				Prompt:   &survey.Input{Message: "Database user:"},
				Validate: survey.Required,
			},
			{
				Name:     "password",
				Prompt:   &survey.Password{Message: "Database password:"},
				Validate: survey.Required,
			},
			{
				Name:     "dbName",
				Prompt:   &survey.Input{Message: "Database name:"},
				Validate: survey.Required,
			},
			{
				Name:     "output",
				Prompt:   &survey.Input{Message: "Output file path:", Default: "backup.sql"},
				Validate: survey.Required,
			},
		}


		answers := struct {
			DBType   string `survey:"dbType"`
			Host     string `survey:"host"`
			Port     string `survey:"port"`
			User     string `survey:"user"`
			Password string `survey:"password"`
			DBName   string `survey:"dbName"`
			Output   string `survey:"output"`
		}{}
		
		
		if err := survey.Ask(qs, &answers); err != nil {
			fmt.Println("Prompt failed:", err)
			return
		}


		portInt, err := strconv.Atoi(answers.Port)
		if err != nil {
			fmt.Println("Invalid port number.")
			return
		}


		dispatchBackup(answers.DBType, answers.Host, portInt, answers.User, answers.Password, answers.DBName, answers.Output)
	},
}

func init() {
	rootCmd.AddCommand(backupCmd)

    backupCmd.Flags().StringVar(&dbType, "type", "", "Database type (mysql|postgres)")
    backupCmd.Flags().StringVar(&host, "host", "", "Database host")
    backupCmd.Flags().IntVar(&port, "port", 3306, "Database port")
    backupCmd.Flags().StringVar(&user, "user", "", "Database username")
    backupCmd.Flags().StringVar(&password, "password", "", "Database password")
    backupCmd.Flags().StringVar(&dbName, "database", "", "Database name")
    backupCmd.Flags().StringVar(&output, "output", "backup.sql", "Output file path for backup")
    backupCmd.Flags().StringVar(&savedConfig, "savedconfig", "", "Use a saved config (provide config name)")
}