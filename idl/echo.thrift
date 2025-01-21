namespace go api

struct Request {
    1: string message
}

struct Response {
    1: string message
}

service Echo {
    Response echo(1: Request req)
}


// ➜  demo_thrift git:(main) ✗ 

// cwgo server --type RPC 
// --idl ../../idl/echo.thrift 
// --service demo_thrift 
// --module github.com/xilepeng/gomall/demo/demo_thrift

// ➜  demo_thrift git:(main) ✗ cwgo server --type RPC --module github.com/xilepeng/gomall/demo/demo_thrift --service demo_thrift --idl ../../idl/echo.thrift