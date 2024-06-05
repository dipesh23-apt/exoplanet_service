
# 🌌 Exoplanet Microservice

## Justification for Meeting Acceptance Criteria

### 1. ✅ Correctness of the Implementation

**Requirements Met:**
- **Add an Exoplanet:** The service allows users to add a new exoplanet by providing the required properties, with validation ensuring correct data.
- **List Exoplanets:** The service lists all available exoplanets stored in memory.
- **Get Exoplanet by ID:** Users can retrieve information about a specific exoplanet by its unique ID.
- **Update Exoplanet:** Users can update details of an existing exoplanet.
- **Delete Exoplanet:** Users can remove an exoplanet from the catalog.
- **Fuel Estimation:** The service calculates the fuel cost for a trip to a particular exoplanet based on the provided formula.
- **Constraints Adhered:** Validation ensures that the constraints for distance, mass, and radius are met.

### 2. 🧩 Code Quality, Organization, and Readability

**Modular Design:**
- **Handlers:** Separated into distinct functions for each operation.
- **Models:** Clearly defined structs and constants for exoplanet types.
- **Factories:** Abstracted creation logic using the Factory Pattern.
- **Validators:** Ensuring data integrity before processing.
- **Store:** A simple in-memory store that encapsulates CRUD operations.

**Readability:**
- **Consistent Naming:** Clear and descriptive variable, function, and type names.
- **Documentation:** Comments explaining complex logic.
- **Error Handling:** Clear and consistent error messages.

### 3. 💻 Proper Use of Go Language Features and Best Practices

- **Packages and Imports:** Proper use of standard and third-party packages (e.g., `gorilla/mux`, `go-playground/validator`).
- **Error Handling:** Consistent error checking and handling throughout the code.
- **Interfaces:** Used for defining behavior (e.g., `ExoplanetFactory`).
- **Concurrency:** While not directly needed in the current scope, the design allows easy integration of concurrency if required in the future.

### 4. 🚨 Error Handling and Validation of User Inputs

**Validation:**
- **Using `go-playground/validator`:** Ensures that all input data meets the specified constraints.
- **Custom Validation:** Specific checks for the type of exoplanet and required fields.

**Error Responses:**
- **Clear HTTP Status Codes:** Proper status codes for different error scenarios.
- **Error Messages:** Detailed error messages for invalid requests.

### 5. 🔧 Ease of Extensibility

**Factory Pattern:**
- New exoplanet types can be added by creating a new factory without modifying existing code.
- Ensures that the system is open for extension but closed for modification (Open/Closed Principle).

**SOLID Principles:**
- **Single Responsibility Principle:** Each component handles a specific part of the logic.
- **Open/Closed Principle:** Easily extendable to add new exoplanet types or additional functionality.
- **Dependency Inversion:** Components depend on abstractions rather than concrete implementations.

### 6. 🐳 Dockerization of the Service

**Dockerfile:**
- A Dockerfile is provided to build and run the microservice in a containerized environment.
- Ensures that the service can run consistently across different environments.

---

