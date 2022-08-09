import { BlogSubGQL, BlogGQL } from './../../generated/graphql';
import { Component, OnInit } from '@angular/core';
import { connect, StringCodec } from 'nats.ws';

@Component({
  selector: 'app-websocket',
  templateUrl: './websocket.component.html',
  styleUrls: ['./websocket.component.css'],
})
export class WebsocketComponent implements OnInit {
  constructor(
    private blogSub: BlogSubGQL,
    private blogNormal: BlogGQL
  ) {}

  async conn() {
    const nc = await connect({ servers: 'ws://0.0.0.0:8081' });
    const sc = StringCodec();
    const sub = nc.subscribe("foo");
    (async () => {
      for await (const m of sub) {
        console.log("==============================");
        console.log("From nats");
        console.log(`[${sub.getProcessed()}]: ${sc.decode(m.data)}`)
        const blog = JSON.parse(sc.decode(m.data))
        console.log(blog)
        console.log("===============================")
      }
      console.log('subscription closed');
    })();
  }

  blog(){
    this.blogNormal.fetch({blogId: 1}).subscribe({next: ({data}) => {
      console.log("normal fetch graphql", data);
    }})
  }

  ngOnInit(): void {
    this.conn()
    this.blog()
    console.log("============");
    
    this.blogSub.subscribe()
  }
}
