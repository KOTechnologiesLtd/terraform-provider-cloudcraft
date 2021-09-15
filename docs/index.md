---
page_title: "cloudcraft Provider"
subcategory: ""
description: |-

---

# cloudcraft Provider





## Schema

### Required

- **apikey** (String, Sensitive) apikey for cloudcraft, can be set using environment variable `CLOUDCRAFT_APITOKEN`

### Optional

- **baseurl** (String) Host URL for cloudcraft, can be set using environment variable, `CLOUDCRAFT_HOST`
- **max_retries** (Number) Max Retries, can be set using environment variable `CLOUDCRAFT_MAX_RETRIES` (defaults to 1).
