FROM golang:1.18.4
RUN apt-get update && apt-get install -y make unzip
RUN curl -sL https://releases.hashicorp.com/terraform/1.4.6/terraform_1.4.6_linux_amd64.zip -o terraform.zip && unzip terraform.zip && mv terraform /bin/terraform && chmod +x /bin/terraform
WORKDIR /build
COPY . .
RUN make