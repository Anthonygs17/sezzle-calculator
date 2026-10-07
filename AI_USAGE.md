# AI Usage

AI assistance was used during the development of this project.

## Tool

- ChatGPT

## How AI Was Used

AI was used as a development assistant for:

- Discussing backend and frontend architecture
- Reviewing Go project structure and idiomatic practices
- Designing unit tests
- Debugging Vitest and coverage configuration
- Reviewing frontend behavior and validation
- Improving documentation

All generated suggestions were reviewed and adapted before being included in the project.

## Example Prompts

Some of the prompts used during development included:

> What folder structure would be appropriate for the Go backend following common production practices?

> Is `internal/http` a good package name, or would it conflict conceptually with Go's `net/http` package?

> Should Go tests for this project use the table-driven testing pattern?

> Is it better to use pointers for the request operands so the backend can distinguish a missing value from zero?

> How should `service_test.go` look after moving missing operand validation to the HTTP handler?

> How should I test a React calculator component using Vitest and React Testing Library?

> Why is Vitest reporting multiple elements with the same text in my calculator tests?

> How can I test the frontend API service while mocking `fetch`?

> Is this level of frontend and backend test coverage reasonable for a small coding assessment?

> Should the generated coverage folder be committed to the repository?

> What should be included in the README for this take-home assessment?

## Notes

AI was used for guidance, review, and troubleshooting. The final implementation decisions, code integration, testing, and validation were performed manually.
