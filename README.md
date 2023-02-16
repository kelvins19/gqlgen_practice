# README

## Tech Stacks
- Go V1.19
- gqlgen for GraphQL framework
- PostgreSQL for database
- bun for ORM
- sql-migrate for database migrations

**To Generate a GraphQL Resolver**
Run `go run github.com/99designs/gqlgen generate`

## How to run migrations
1. Run `go install github.com/rubenv/sql-migrate/...@latest`
2. Run `sql-migrate up -config=dbconfig.yml`

**To Create new migrations file**
Run `sql-migrate new -config=dbconfig.yml -env="ENV" <MIGRATION_NAME>`

## How to setup the application
1. Go to the project directory
2. Run `go build`
3. Run `go mod tidy`
4. Run `go run server.go`
5. Open browser and access http://localhost:8080

## Query
### Categories
#### Create Category
```
mutation createCategory {
 createCategory(
   input: {
     name: "test"
     description: "test"
   }
 ){
   id
 }
}
```

#### Update Category
```
mutation updateCategory {
  updateCategory (
    id: 6,
    input: {
      name: "Test 6",
      description: "Test6"
    }
  ) {
    id
    name
    description
  }
}
```

#### Delete Category
```
mutation deleteCategory {
  deleteCategory (id: 7) {
    id
    name
    description
  }
}
```

#### Get Categories
```
query getCategories {
  getCategories {
    id
    name
    description
    products {
      name
      description
      price
    }
  }
}
```

#### Get Single Category By ID
```
query getSingleCategories {
  getSingleCategories (Id: 2) {
    id
    name
    description
    products {
      name
      description
      price
    }
  }
}
```

### Products
#### Create Product
```
mutation createProduct {
 createProduct(
   input: {
     name: "test"
     description: "test"
     price: 1000
     categories: [1,2]
   }
 ){
   id
 }
}
```

#### Update Product
```
mutation updateProduct {
  updateProduct (
    id: 2,
    input: {
      name: "Adidas UltraBoost",
      description: "Ultraboost",
      price: 2299000,
      categories: [2,3]
    }
  ) {
    id
    name
    description
    price
    categories {
      id
      name
      description
    }
  }
}
```

#### Delete Product
```
mutation deleteProduct {
  deleteProduct (id: 3) {
    id
    name
    description
    price
  }
}
```

#### Get Products
```
query getProduct {
 products {
    id
    name
    description
    price
    categories {
      name
      description
    }
  }
}
```

#### Get Single Product By ID
```
query getSingleProducts {
  getSingleProducts (Id: 2) {
    id
    name
    description
    price
    categories {
      name
      description
    }
  }
}
```