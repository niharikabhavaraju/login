/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	common "login/common"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
)

var username, password string

var router = mux.NewRouter()

// loginCmd represents the login command

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login to your account",
	Long: `Login to your account by providing username and password  command:go run main.go login --username niharika --password 1234 

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login called")
		common.SetUsernamePassword(username, password)
		router.HandleFunc("/", common.LoginPageHandler)
		router.HandleFunc("/index", common.IndexPageHandler)
		router.HandleFunc("/invalid", common.InvalidPageHandler)
		router.HandleFunc("/login", common.LoginHandler).Methods("POST")
		router.HandleFunc("/logout", common.LogoutHandler).Methods("POST")
		http.Handle("/", router)
		http.ListenAndServe(":8080", nil)
		open.RunWith("http://localhost:8080/", "firefox")
		fmt.Println("Done")

	},
}

func init() {

	rootCmd.AddCommand(loginCmd)
	loginCmd.PersistentFlags().StringVar(&username, "username", "", "enter username")
	loginCmd.MarkPersistentFlagRequired("username")
	loginCmd.PersistentFlags().StringVar(&password, "password", "", "enter password")
	loginCmd.MarkPersistentFlagRequired("password")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
