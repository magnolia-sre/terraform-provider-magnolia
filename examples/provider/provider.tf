terraform {
  required_providers {
    magnolia = {
      version = "0.0.1"
      source = "magnolia/magnolia"
    }
  }
}

provider "magnolia" {
  token = "absadkadsjka" // or MAGNOLIA_TOKEN
}
