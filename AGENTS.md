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
*   **Unit Tests are Mandatory:** Every new resource, data source, CLI command, or utility function must include comprehensive unit tests. Please be overzealous in unit testing everything. Maximise unit test coverage.
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

## 6. Versioning & Commit Standards
*   **Semantic Versioning:** Strictly follow Semantic Versioning (SemVer). The `VERSION` variable in `Taskfile.yml` must adhere to these rules.
*   **Conventional Commits:** Always use Conventional Commits (e.g., `feat:`, `fix:`, `chore:`).
*   **Commit Message Structure:** All commit messages must explicitly include the following sections in the body:
    *   **Context:** (Explain why this change is being made)
    *   **Changes:** (Describe what modifications are included)
    *   **Tests:** (Detail how this change was validated)

## 7. Linting & Formatting
*   **Single Source of Truth:** Use `golangci-lint` as the primary tool for linting, formatting, vetting, and security checks.
*   **No Redundant Steps:** Do not introduce separate CI steps or standalone commands for `go fmt`, `go vet`, or `gosec`. Rely entirely on the `.golangci.yml` configuration.
*   **Pre-Commit Requirement:** ALWAYS run `go test ./...` and `golangci-lint run` and ensure they both pass before committing any code changes.

## 8. OpenAPI Specifications
When developing, do not run web searches for the API. Always use the following specs:
- Network: [https://developer.ui.com/network/v10.4.57/openapi.json](https://developer.ui.com/network/v10.4.57/openapi.json)
- Protect: [https://developer.ui.com/protect/v7.1.87/openapi.json](https://developer.ui.com/protect/v7.1.87/openapi.json)
