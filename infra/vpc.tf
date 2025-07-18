resource "aws_vpc" "dev_vpc" {
  cidr_block           = "10.0.0.0/16"
  enable_dns_hostnames = true

  tags = {
    Name = "dev-vpc"
  }
}

resource "aws_subnet" "public_subnet_1" {
  vpc_id            = aws_vpc.dev_vpc.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "ca-central-1a"
  map_public_ip_on_launch = true

  tags = {
    Name = "dev-public-subnet-1"
  }
}

resource "aws_subnet" "public_subnet_2" {
  vpc_id            = aws_vpc.dev_vpc.id
  cidr_block        = "10.0.2.0/24"
  availability_zone = "ca-central-1b"
  map_public_ip_on_launch = true

  tags = {
    Name = "dev-public-subnet-2"
  }
}

resource "aws_internet_gateway" "dev_igw" {
  vpc_id = aws_vpc.dev_vpc.id
}

resource "aws_route_table" "public_rt" {
  vpc_id = aws_vpc.dev_vpc.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.dev_igw.id
  }
}

resource "aws_route_table_association" "public_subnet_1_assoc" {
  subnet_id      = aws_subnet.public_subnet_1.id
  route_table_id = aws_route_table.public_rt.id
}

resource "aws_route_table_association" "public_subnet_2_assoc" {
  subnet_id      = aws_subnet.public_subnet_2.id
  route_table_id = aws_route_table.public_rt.id
}

