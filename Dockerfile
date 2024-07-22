# # Use the Go base image for building the Go application
FROM golang:1.20-alpine AS builder

# Set up the working directory
WORKDIR /app

# Copy Go module files
# COPY go.mod go.sum ./
# RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o go-basex-validator .

# Use a base image with Java (BaseX requires Java)
FROM openjdk:11-jre-slim

# Install required packages and tools
RUN apt-get update && \
    apt-get install -y wget unzip && \
    # Clean up the apt cache to reduce image size
    rm -rf /var/lib/apt/lists/*

# Define BaseX version
ENV BASEX_VERSION=11.1

# Create the BaseX directory
RUN mkdir -p /basex/${BASEX_VERSION}

# Download and install BaseX
RUN wget https://files.basex.org/releases/${BASEX_VERSION}/BaseX111.zip -O BaseX-${BASEX_VERSION}.zip && \
    unzip BaseX-${BASEX_VERSION}.zip -d /basex && \
    mv /basex/basex/* /basex/${BASEX_VERSION} && \
    rm -rf /basex/basex/ && \
    rm BaseX-${BASEX_VERSION}.zip

# Set environment variables
ENV BASEX_HOME=/basex/${BASEX_VERSION}
ENV PATH="$PATH:$BASEX_HOME/bin"

# Download the validation artifacts from the specified release
ENV VALIDATION_ZIP_URL=https://github.com/ConnectingEurope/eInvoicing-EN16931/releases/download/validation-1.3.12/en16931-ubl-1.3.12.zip
RUN wget ${VALIDATION_ZIP_URL} -O en16931-ubl-1.3.12.zip && \
    unzip en16931-ubl-1.3.12.zip -d /basex/${BASEX_VERSION}/validation-artifacts && \
    rm en16931-ubl-1.3.12.zip

# Debug: List the files in the validation-artifacts directory
# RUN echo "Listing contents of /basex/${BASEX_VERSION}/validation-artifacts/en16931-ubl-1.3.12:" && \
#     ls -la /basex/${BASEX_VERSION}/validation-artifacts/en16931-ubl-1.3.12/schematron/EN16931-UBL-validation.sch && \
#     echo "Listing contents of /basex/${BASEX_VERSION}/validation-artifacts/xml:" && \
#     ls -la /basex/${BASEX_VERSION}/validation-artifacts/xml

# Copy the Go binary from the builder stage
COPY --from=builder /app/go-basex-validator /usr/local/bin/go-basex-validator

# Set working directory
WORKDIR /basex/${BASEX_VERSION}

