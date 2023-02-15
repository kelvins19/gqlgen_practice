# README

## Tech Stacks
- Go V1.19
- gqlgen for GraphQL framework
- PostgreSQL for database
- bun for ORM

**To Generate a GraphQL Resolver**
Run `go run github.com/99designs/gqlgen generate`

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

#### Update Product

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