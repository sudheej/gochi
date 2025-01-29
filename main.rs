use reqwest::Client;
use tiny_http::{Server, Response, Header};
use std::sync::Arc;
use std::error::Error;

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error + Send + Sync>> {
    // Start the proxy server
    let server = Server::http("127.0.0.1:8080")?;
    let client = Arc::new(Client::new());

    println!("Proxy server running on http://127.0.0.1:8080");

    for request in server.incoming_requests() {
        let client = Arc::clone(&client);

        tokio::spawn(async move {
            let method = request.method().as_str();
            let url = format!("https://your-artifactory-url.com{}", request.url());

            let builder = match method {
                "GET" => client.get(&url),
                "POST" => client.post(&url),
                "PUT" => client.put(&url),
                "DELETE" => client.delete(&url),
                _ => {
                    let response = Response::from_string("405 Method Not Allowed").with_status_code(405);
                    request.respond(response).ok();
                    return;
                }
            };

            // Send the request with a Bearer token
            match builder.bearer_auth("your-bearer-token").send().await {
                Ok(resp) => {
                    let status_code = resp.status().as_u16(); // Extract status code **before consuming response**
                    let headers = resp.headers().clone(); // Clone headers **before consuming response**
                    match resp.text().await {
                        Ok(body) => {
                            let mut proxy_response = Response::from_string(body).with_status_code(status_code);

                            // Convert response headers correctly
                            for (key, value) in headers.iter() {
                                if let Ok(value_str) = value.to_str() {
                                    if let Ok(header) = Header::from_bytes(key.as_str(), value_str) {
                                        proxy_response.add_header(header);
                                    }
                                }
                            }

                            request.respond(proxy_response).ok();
                        }
                        Err(_) => {
                            let response = Response::from_string("502 Bad Gateway").with_status_code(502);
                            request.respond(response).ok();
                        }
                    }
                }
                Err(_) => {
                    let response = Response::from_string("502 Bad Gateway").with_status_code(502);
                    request.respond(response).ok();
                }
            }
        });
    }

    Ok(())
}
