# STEP 1
FROM 1.23.0-alpine3.20

# STEP 2
WORKDIR /app

# STEP 3
COPY go.mod go.sum ./

# STEP 4
RUN go mod download

# STEP 5
COPY . .

# STEP 6
RUN make run