# Agentic Guidelines

These guidelines are meant for autonomous AI agents contributing to this codebase.

## 1. Idiomatic Golang
*   Always write idiomatic Golang code.
*   Keep files logically separated. Use subpackages when appropriate.
*   Enforce single responsibility for structs and interfaces.

## 2. Dependency Injection & Interfaces (First Defaults)
*   **Interfaces over Structs:** Whenever creating API clients, services, or internal dependencies, define an interface first.
*   **Dependency Injection:** Resources, Data Sources, and other business logic must accept dependencies via interfaces, not concrete types.
*   Example: `network_res` depends on `client.NetworkClient` interface, not `*client.Client`.

## 3. Unit Testing, Mocking & Testdata
*   **Unit Tests are Mandatory:** Every new resource or data source must include comprehensive unit tests.
*   **Testdata Pattern:** Do not inline large raw text blocks or JSON payloads in your test files. Instead, use the `testdata/` directory pattern to store mock inputs (e.g., API responses) and expected outputs (e.g., Terraform states or plan objects). Read these files during test execution.
*   **Mocking:** Use the dependency injected interfaces to mock external API interactions where possible, or use Terraform's testing framework (`resource.UnitTest`) to validate configurations and plans using the aforementioned `testdata/` files.
*   Ensure that edge cases (like invalid types or missing required fields) are captured in unit tests.

## 4. Terraform Plugin Framework (v6)
*   Strictly use the `terraform-plugin-framework` instead of the old SDKv2.
*   Use separate domain models for Plans and States.
*   Implement explicit type conversions (`mapToEndpointDto`, etc.) between the Terraform models and the API DTOs.
*   Configure plan modifiers (`UseStateForUnknown`, `RequiresReplace`) accurately according to the API behavior.

## 5. Security & Credentials
*   Never hardcode API keys or sensitive endpoints.
*   Use `terraform-plugin-framework`'s provider configuration options to accept credentials, or fallback to environment variables.
