# CEL and CloudEvents

Uses CEL to evaluate expressions against a CloudEvent envelope.

## Default expression

```
go run main.go
```

```
expr: ce.type == "com.github.pull_request.create"
ParsedExpr: expr:<id:4 call_expr:<function:"_==_" args:<id:2 select_expr:<operand:<id:1 ident_expr:<name:"ce" > > field:"type" > > args:<id:3 const_expr:<string_value:"com.github.pull_request.create" > > > > source_info:<location:"<input>" line_offsets:44 positions:<key:1 value:0 > positions:<key:2 value:2 > positions:<key:3 value:11 > positions:<key:4 value:8 > > 
CheckedExpr: reference_map:<key:1 value:<name:"ce" > > reference_map:<key:4 value:<overload_id:"equals" > > type_map:<key:1 value:<message_type:"dev.knative.CloudEvent" > > type_map:<key:2 value:<primitive:STRING > > type_map:<key:3 value:<primitive:STRING > > type_map:<key:4 value:<primitive:BOOL > > source_info:<location:"<input>" line_offsets:44 positions:<key:1 value:0 > positions:<key:2 value:2 > positions:<key:3 value:11 > positions:<key:4 value:8 > > expr:<id:4 call_expr:<function:"_==_" args:<id:2 select_expr:<operand:<id:1 ident_expr:<name:"ce" > > field:"type" > > args:<id:3 const_expr:<string_value:"com.github.pull_request.create" > > > > 
cloudEvent:  type:"com.github.pull_request.create" source:"github.com/grantr/cel-playground/pulls/21" 
out: true
```

## Custom expression
```
go run main.go 'ce.source == "foo"'
```

```
expr: ce.source == "foo"
ParsedExpr: expr:<id:4 call_expr:<function:"_==_" args:<id:2 select_expr:<operand:<id:1 ident_expr:<name:"ce" > > field:"source" > > args:<id:3 const_expr:<string_value:"foo" > > > > source_info:<location:"<input>" line_offsets:19 positions:<key:1 value:0 > positions:<key:2 value:2 > positions:<key:3 value:13 > positions:<key:4 value:10 > > 
CheckedExpr: reference_map:<key:1 value:<name:"ce" > > reference_map:<key:4 value:<overload_id:"equals" > > type_map:<key:1 value:<message_type:"dev.knative.CloudEvent" > > type_map:<key:2 value:<primitive:STRING > > type_map:<key:3 value:<primitive:STRING > > type_map:<key:4 value:<primitive:BOOL > > source_info:<location:"<input>" line_offsets:19 positions:<key:1 value:0 > positions:<key:2 value:2 > positions:<key:3 value:13 > positions:<key:4 value:10 > > expr:<id:4 call_expr:<function:"_==_" args:<id:2 select_expr:<operand:<id:1 ident_expr:<name:"ce" > > field:"source" > > args:<id:3 const_expr:<string_value:"foo" > > > > 
cloudEvent:  type:"com.github.pull_request.create" source:"github.com/grantr/cel-playground/pulls/21" 
out: false
```
