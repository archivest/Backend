# Backend

The Backend API for BibleBot.

## Internal Organization

While the backend repository itself is a monolith, multiple packages will exist in `internal/` that all hook into the base API. These packages are internal, so cannot be imported standalone.