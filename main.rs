use reqwest::Client;
use tiny_http::{Server, Request, Response};
use std::sync::Arc;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    // Create a lightweight HTTP server.
    let server = Server::http("127.0.0.1:8080")?;
    let client = Arc::new(Client::new());
    println!("Proxy server running on http://127.0.0.1:8080");

    for request in server.incoming_requests() {
        let client = Arc::clone(&client);

        tokio::spawn(async move {
            // Extract the request details.
            let method = request.method().as_str();
            let url = format!("https://your-artifactory-url.com{}", request.url());

            // Create the proxy request.
            let builder = match method {
                "GET" => client.get(&url),
                "POST" => client.post(&url),
                "PUT" => client.put(&url),
                "DELETE" => client.delete(&url),
                _ => {
                    let response = Response::from_string("Method Not Allowed").with_status_code(405);
                    request.respond(response).ok();
                    return;
                }
            };

            // Add the Bearer token and send the request.
            let response = builder.bearer_auth("your-bearer-token").send().await;

            match response {
                Ok(resp) => {
                    let mut proxy_response = Response::from_string(resp.text().await.unwrap_or_default())
                        .with_status_code(resp.status().as_u16());

                    // Forward headers.
                    for (key, value) in resp.headers() {
                        if let Ok(value) = value.to_str() {
                            proxy_response.add_header(format!("{}: {}", key, value));
                        }
                    }

                    request.respond(proxy_response).ok();
                }
                Err(_) => {
                    let response = Response::from_string("Bad Gateway").with_status_code(502);
                    request.respond(response).ok();
                }
            }
        });
    }

    Ok(())
}
