import {NgModule} from '@angular/core';
import {ApolloModule, APOLLO_OPTIONS} from 'apollo-angular';
import {ApolloClientOptions, InMemoryCache, split} from '@apollo/client/core';
import {WebSocketLink} from '@apollo/client/link/ws'
import {HttpLink} from 'apollo-angular/http';
import {getMainDefinition} from '@apollo/client/utilities'

// const uri = "http://localhost:8080/query"; // <-- add the URL of the GraphQL server here
// export function createApollo(httpLink: HttpLink): ApolloClientOptions<any> {
//   const ws = new WebSocketLink({
//     uri: `ws://localhost:8080/query`,
//     options: {
//       reconnect: true,
//     },
//   });
//   return {
//     link: httpLink.create({uri}),
//     cache: new InMemoryCache(),
//   };
// }
export function createApollo(httpLink: HttpLink): ApolloClientOptions<any> {
 
  const http = httpLink.create({
    uri:"http://localhost:8080/query"
  });
 
  const ws = new WebSocketLink({
    uri:`wss://localhost:8080/query`,
    options:{
      reconnect:true
    }
  });
 
  const link = split(
    ({query}) => {
      const data = getMainDefinition(query);
      return (
        data.kind === 'OperationDefinition' && data.operation === 'subscription'
      );
    },
    ws,
    http
  )
 
  return {
    link: link,
    cache: new InMemoryCache(),
  };
}

@NgModule({
  exports: [ApolloModule],
  providers: [
    {
      provide: APOLLO_OPTIONS,
      useFactory: createApollo,
      deps: [HttpLink],
    },
  ],
})
export class GraphQLModule {}
