# Krow

Krow is a file based REST API client. If you are familiar with tools like Postman, Insomnia, or curl, then you will feel right at home with Krow. This is file based, opinionated, command line tool. Please follow Installation and Usage guide to get started.

## Installation

### MacOS x86_64

```bash
curl -s https://raw.githubusercontent.com/st-osi/krow/main/install/darwin-amd64.sh | bash
```

### MacOS arm64

```bash
curl -s https://raw.githubusercontent.com/st-osi/krow/main/install/darwin-arm64.sh | bash
```

### Linux x86_64

```bash
curl -s https://raw.githubusercontent.com/st-osi/krow/main/install/linux-amd64.sh | bash
```

## Getting Started

1. Create request yaml file
2. Create .env or .env.local file if you want to use environment variables in your request file
3. Run `krow run <request-file>`

## Example Of Request File

```yaml
# File name: sample-post.yaml
Name: Create Post
URL: "${API_URL}/posts" # ${API_URL} is set on the environment file (.env or .env.local)
Method: POST # Request method (GET, POST, PUT, DELETE, PATCH)

# Add Request Headers as needed
Headers:
  Accept: text/html, application/json
  Accept-Encoding: utf-8
  X-Proxy-Agent: https://something.com:8080
  User-Agent: krow-client-0.0.1
  Content-Type: application/json

Body:
  name: John
  age: 30
  address:
    street: 123 Main St
    city: Anytown
  hobbies: # Array of hobbies
    - reading
    - gaming
  contact:
    email: john.doe@example.com
    phone:
      home: "123-456-7890"
      work: "098-765-4321"
  employment:
    company: Example Corp
    position: Software Engineer
    details:
      start_date: "2020-01-15"
      end_date: null
      projects: # Array of projects
        - name: Project Alpha
          role: Lead Developer
        - name: Project Beta
          role: Contributor

# After hook is for setting environment variables from the response
# Body and Header are the only supported sources for now
# User [] notation to access elements from map or array
# Check your environment file (.env or .env.local) for new/updated variables
After:
  Env:
    ADDRESS_STREET: Body[address][street] # Accessing street from address map
    HOBBIES: Body[hobbies][0] # Accessing first element of hobbies array
    RESPONSE_DATE: Header[Date]
    EMPLOYMENT_START_DATE: Body[employment][details][start_date]
```

To run the request file:

```bash
krow run sample-post.yaml
```

## Response

Response will be written to the markdown file (as most of the editor supports formatting for markdown) in separate folder named as `.res.<request-name>` and response file will be in the format of `.<request-name>.<http-method>.<timestamp>.md`. If you don't like the way it is managed, feel free to create issue or PR.

### Response file structure

[Sample File](./sample/sample-response)

You can ignore the response folder for git by adding `**/.res.**/` to `.gitignore` file as it will be different for each request file and will be generated on runtime.

## Contribution

### Build

```bash
go build -o krow bin/main.go

```
