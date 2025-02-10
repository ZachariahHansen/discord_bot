# Read config.json and set environment variables
$config = Get-Content -Raw -Path ".\api\config.json" | ConvertFrom-Json

# Set environment variables
$env:POSTGRES_USER = $config.POSTGRES_USER
$env:POSTGRES_PASSWORD = $config.POSTGRES_PASSWORD
$env:POSTGRES_DB = $config.POSTGRES_DB

# Run docker-compose
docker-compose up --build -d
