# app.rb
require 'sinatra'
require 'rack/cors'

# Configure CORS
use Rack::Cors do
  allow do
    origins '*'  # Allows all origins, replace '*' with a specific origin in production
    resource '*', 
      headers: :any, 
      methods: [:get, :post, :put, :delete, :options]
  end
end

# Health check endpoint
get '/' do
  'UP'
end
