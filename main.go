package main

import (
  "fmt"
  "github.com/spf13/viper"
)

func main() {

  viper.SetConfigName("app")     // no need to include file extension
  viper.AddConfigPath("config")  // set the path of your config file

  err := viper.ReadInConfig()
  if err != nil {
    fmt.Println("Config file not found...")
  } else {
    dev_server := viper.GetString("development.server")
    dev_connection_max := viper.GetInt("development.connection_max")
    dev_enabled := viper.GetBool("development.enabled")
    dev_port := viper.GetInt("development.port")

    fmt.Printf("\nDevelopment Config found:\n server = %s\n connection_max = %d\n" +
        " enabled = %t\n" +
        " port = %d\n",
        dev_server,
        dev_connection_max,
        dev_enabled,
        dev_port)

    prod_server := viper.GetString("production.server")
    prod_connection_max := viper.GetInt("production.connection_max")
    prod_enabled := viper.GetBool("production.enabled")
    prod_port := viper.GetInt("production.port")

    fmt.Printf("\nProduction Config found:\n server = %s\n connection_max = %d\n" +
        " enabled = %t\n" +
        " port = %d\n",
        prod_server,
        prod_connection_max,
        prod_enabled,
        prod_port)
  }

}