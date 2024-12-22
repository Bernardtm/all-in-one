use actix_cors::Cors;
use actix_web::{web, App, HttpResponse, HttpServer, Responder};

// Get a user by ID
async fn get_status() -> impl Responder {
    HttpResponse::Ok().body("UP")
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(move || {
        let cors = Cors::default()
            .allow_any_origin()
            .allowed_methods(vec!["GET", "POST"]) // Specify allowed methods
            .allowed_headers(vec![
                actix_web::http::header::AUTHORIZATION,
                actix_web::http::header::ACCEPT,
            ])
            .allowed_header(actix_web::http::header::CONTENT_TYPE)
            .supports_credentials();

        App::new()
            .wrap(cors) // Apply the CORS middleware to the app
            .route("/", web::get().to(get_status))
    })
    .bind("127.0.0.1:7000")?
    .run()
    .await
}
