## My Information

```
Name:    Mohamed Bamoh  
Email:   mbamoh@icloud.com 
```

## Description

This repository is an alternative implementation of
the [Ardan Labs Service Repository](https://github.com/ardanlabs/service), inspired by Bill Kennedy's **Ultimate
Software Design with Kubernetes** course. It combines the architectural principles and design guidelines from the
original project with:

1. **Contract-First Development**: Using OpenAPI Specification and the [Ogen](https://github.com/ogen-go/ogen) tool for
   server stub generation.
2. **Efficient SQL Queries**: Leveraging [sqlc](https://sqlc.dev/) to generate type-safe, performant Go code for the
   storage layer.

## Motivation

In many modern organizations, contract-first development plays a pivotal role in software design. By defining an OpenAPI
Specification (OAS) as the contract between various stakeholders—such as frontend and backend teams, external
integrators, or microservices—it becomes possible to achieve:

- **Improved collaboration**: Clear API definitions that act as a source of truth.
- **Faster development**: Automated code generation for client and server stubs.
- **Stronger alignment**: Ensuring teams agree on API expectations upfront.

This project also introduces **sqlc** to the mix, which simplifies database interactions by:

- Generating **type-safe queries** directly from SQL statements.
- Eliminating the need for ORMs while maintaining maintainability and high performance.
- Aligning with the **clean architecture principles** taught in the course.

By combining **Ogen** and **sqlc**, this repository demonstrates how to incorporate modern tools into a well-structured
and maintainable service design.

## Why Use This Repo?

This implementation is ideal for developers who:

- Want to learn or apply **contract-first development** using OpenAPI and Go.
- Appreciate the structured approach to service design taught by Ardan Labs.
- Seek a modern alternative to the original service repo by leveraging tools like **Ogen** and **sqlc**.
- Want a simpler yet powerful database interaction layer without relying on heavy ORMs.

## Acknowledgements

Special thanks to [Ardan Labs](https://github.com/ardanlabs/service)
and [Bill Kennedy](https://github.com/ardan-bkennedy) for the foundational principles and course material that inspired
this project.
