package handlers

import (
    "api/internal/config"
)

var base_url string

func getBaseURL() string {
    if base_url == "" {
        config, err := config.LoadConfig()
        if err != nil {
            return ""
        }
        base_url = config.BaseURL
    }
    return base_url
}
